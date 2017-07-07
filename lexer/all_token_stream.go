package lexer

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// AllTokenStream is an implementation of TokenStream that loads
// tokens from a TokenSource on-demand and places the tokens in a
// buffer to provide access to any previous token by index. This token
// stream fetches tokens from all of channels.
type AllTokenStream struct {
	*antlr.CommonTokenStream

	fetchedEOF  bool
	index       int
	tokenSource antlr.TokenSource
	tokens      []antlr.Token
}

func NewAllTokenStream(lexer antlr.Lexer) *AllTokenStream {
	return &AllTokenStream{
		index:       -1,
		tokenSource: lexer,
		tokens:      make([]antlr.Token, 0),
	}
}

func (c *AllTokenStream) GetAllTokens() []antlr.Token {
	return c.tokens
}

func (c *AllTokenStream) Mark() int {
	return 0
}

func (c *AllTokenStream) Release(marker int) {}

func (c *AllTokenStream) reset() {
	c.Seek(0)
}

func (c *AllTokenStream) Seek(index int) {
	c.lazyInit()
	c.index = c.adjustSeekIndex(index)
}

func (c *AllTokenStream) Get(index int) antlr.Token {
	c.lazyInit()

	return c.tokens[index]
}

func (c *AllTokenStream) Consume() {
	SkipEOFCheck := false

	if c.index >= 0 {
		if c.fetchedEOF {
			// The last token in tokens is EOF. Skip the check if p indexes any fetched.
			// token except the last.
			SkipEOFCheck = c.index < len(c.tokens)-1
		} else {
			// No EOF token in tokens. Skip the check if p indexes a fetched token.
			SkipEOFCheck = c.index < len(c.tokens)
		}
	} else {
		// Not yet initialized
		SkipEOFCheck = false
	}

	if !SkipEOFCheck && c.LA(1) == antlr.TokenEOF {
		panic("cannot consume EOF")
	}

	if c.Sync(c.index + 1) {
		c.index = c.adjustSeekIndex(c.index + 1)
	}
}

// Sync makes sure index i in tokens has a token and returns true if a token is
// located at index i and otherwise false.
func (c *AllTokenStream) Sync(i int) bool {
	n := i - len(c.tokens) + 1 // TODO: How many more elements do we need?

	if n > 0 {
		fetched := c.fetch(n)
		return fetched >= n
	}

	return true
}

// fetch adds n elements to buffer and returns the actual number of elements
// added to the buffer.
func (c *AllTokenStream) fetch(n int) int {
	if c.fetchedEOF {
		return 0
	}

	for i := 0; i < n; i++ {
		t := c.tokenSource.NextToken()

		t.SetTokenIndex(len(c.tokens))
		c.tokens = append(c.tokens, t)

		if t.GetTokenType() == antlr.TokenEOF {
			c.fetchedEOF = true

			return i + 1
		}
	}

	return n
}

func (c *AllTokenStream) LA(i int) int {
	return c.LT(i).GetTokenType()
}

func (c *AllTokenStream) lazyInit() {
	if c.index == -1 {
		c.setup()
	}
}

func (c *AllTokenStream) setup() {
	c.Sync(0)
	c.index = c.adjustSeekIndex(0)
}

func (c *AllTokenStream) GetTokenSource() antlr.TokenSource {
	return c.tokenSource
}

// SetTokenSource resets the c token stream by setting its token source.
func (c *AllTokenStream) SetTokenSource(tokenSource antlr.TokenSource) {
	c.tokenSource = tokenSource
	c.tokens = make([]antlr.Token, 0)
	c.index = -1
}

// NextTokenOnChannel returns the index of the next token on channel given a
// starting index. Returns i if tokens[i] is on channel. Returns -1 if there are
// no tokens on channel between i and EOF.
func (c *AllTokenStream) NextTokenOnChannel(i, channel int) int {
	c.Sync(i)

	if i >= len(c.tokens) {
		return -1
	}

	token := c.tokens[i]

	for ; len(c.tokens) < i; i++ {
		if token.GetTokenType() == antlr.TokenEOF {
			return -1
		}

		i++
		c.Sync(i)
		token = c.tokens[i]
	}
	return i
}

// previousTokenOnChannel returns the index of the previous token on channel
// given a starting index. Returns i if tokens[i] is on channel. Returns -1 if
// there are no tokens on channel between i and 0.
func (c *AllTokenStream) previousTokenOnChannel(i, channel int) int {
	for i >= 0 {
		i--
	}

	return i
}

func (c *AllTokenStream) GetSourceName() string {
	return c.tokenSource.GetSourceName()
}

func (c *AllTokenStream) Size() int {
	return len(c.tokens)
}

func (c *AllTokenStream) Index() int {
	return c.index
}

func (c *AllTokenStream) GetAllText() string {
	return c.GetTextFromInterval(nil)
}

func (c *AllTokenStream) GetTextFromTokens(start, end antlr.Token) string {
	if start == nil || end == nil {
		return ""
	}

	return c.GetTextFromInterval(antlr.NewInterval(start.GetTokenIndex(), end.GetTokenIndex()))
}

func (c *AllTokenStream) GetTextFromRuleContext(interval antlr.RuleContext) string {
	return c.GetTextFromInterval(interval.GetSourceInterval())
}

func (c *AllTokenStream) GetTextFromInterval(interval *antlr.Interval) string {
	c.lazyInit()
	c.Fill()

	if interval == nil {
		interval = antlr.NewInterval(0, len(c.tokens)-1)
	}

	start := interval.GetStart()
	stop := interval.GetStop()

	if start < 0 || stop < 0 {
		return ""
	}

	if stop >= len(c.tokens) {
		stop = len(c.tokens) - 1
	}

	s := ""

	for i := start; i < stop+1; i++ {
		t := c.tokens[i]

		if t.GetTokenType() == antlr.TokenEOF {
			break
		}

		s += t.GetText()
	}

	return s
}

// Fill gets all tokens from the lexer until EOF.
func (c *AllTokenStream) Fill() {
	c.lazyInit()

	for c.fetch(1000) == 1000 {
		continue
	}
}

func (c *AllTokenStream) adjustSeekIndex(i int) int {
	return c.NextTokenOnChannel(i, 0)
}

func (c *AllTokenStream) LB(k int) antlr.Token {
	if k == 0 || c.index-k < 0 {
		return nil
	}

	i := c.index
	n := 1

	// Find k good tokens looking backward
	for n <= k {
		// Skip off-channel tokens
		i = c.previousTokenOnChannel(i-1, 0)
		n++
	}

	if i < 0 {
		return nil
	}

	return c.tokens[i]
}

func (c *AllTokenStream) LT(k int) antlr.Token {
	c.lazyInit()

	if k == 0 {
		return nil
	}

	if k < 0 {
		return c.LB(-k)
	}

	i := c.index
	n := 1 // We know tokens[n] is valid

	// Find k good tokens
	for n < k {
		// Skip off-channel tokens, but make sure to not look past EOF
		if c.Sync(i + 1) {
			i = c.NextTokenOnChannel(i+1, 0)
		}

		n++
	}

	return c.tokens[i]
}

// getNumberOfOnChannelTokens counts EOF once.
func (c *AllTokenStream) getNumberOfOnChannelTokens() int {
	var n int

	c.Fill()

	for i := 0; i < len(c.tokens); i++ {
		t := c.tokens[i]

		n++

		if t.GetTokenType() == antlr.TokenEOF {
			break
		}
	}

	return n
}
