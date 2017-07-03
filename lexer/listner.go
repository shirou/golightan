package lexer

import (
	"fmt"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/shirou/golightan"
)

const (
	InitialStackCapacity = 10
	InitialTokenCapacity = 100
)

type CommonParseTreeListener struct {
	tokens   golightan.Tokens
	tokenMap TokenMap
	rule     int
	stack    *Stack
	lexer    antlr.Lexer
}

func (b *CommonParseTreeListener) Token(node antlr.TerminalNode) {
	token := node.GetSymbol()
	t := token.GetTokenType()
	if t < 0 {
		return
	}

	text := node.GetText()
	new := golightan.Token{
		OriginalToken: token,
		TokenType:     b.tokenMap.Convert(b.stack.Last(), t, text),
		Text:          text,
	}

	if b.lexer != nil {
		b.debug(node)
	}

	b.tokens = append(b.tokens, new)
}

func (b *CommonParseTreeListener) GetTokens() golightan.Tokens { return b.tokens }

func (b *CommonParseTreeListener) VisitTerminal(node antlr.TerminalNode) {
	b.Token(node)
}
func (b *CommonParseTreeListener) VisitErrorNode(node antlr.ErrorNode) {
	b.Token(node)
}
func (b *CommonParseTreeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	b.stack.Push(ctx.GetRuleIndex())
}
func (b *CommonParseTreeListener) ExitEveryRule(ctx antlr.ParserRuleContext) {
	b.stack.Pop()
}

func (b *CommonParseTreeListener) SetDebug(lexer antlr.Lexer) {
	b.lexer = lexer
}

func (b *CommonParseTreeListener) debug(node antlr.TerminalNode) {
	token := node.GetSymbol()
	t := token.GetTokenType()
	s := make([]string, b.stack.Len())
	names := b.lexer.GetRuleNames()
	symbol := b.lexer.GetSymbolicNames()
	for i, r := range b.stack.stack {
		s[i] = names[r]
	}
	fmt.Printf("(%s, %s)-> %v\n", strings.Join(s, ","), symbol[t], node.GetText())
}

func NewCommonParseTreeListener(tm TokenMap) *CommonParseTreeListener {
	return &CommonParseTreeListener{
		tokens:   make(golightan.Tokens, 0, InitialTokenCapacity),
		tokenMap: tm,
		stack:    NewStack(InitialStackCapacity),
	}
}
