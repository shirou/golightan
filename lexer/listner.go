package lexer

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/shirou/golightan"
)

type CommonParseTreeListener struct {
	tokens   golightan.Tokens
	tokenMap TokenMap
	rule     int
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
		TokenType:     b.tokenMap.Convert(b.rule, t, text),
		Text:          text,
	}

	// If debugging, comment in this line to show current node
	//	fmt.Println(b.rule, t, node.GetText())

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
	b.rule = ctx.GetRuleIndex()
}
func (b *CommonParseTreeListener) ExitEveryRule(ctx antlr.ParserRuleContext) {
	b.rule = -1
}

func NewCommonParseTreeListener(tm TokenMap) *CommonParseTreeListener {
	return &CommonParseTreeListener{
		tokens:   make(golightan.Tokens, 0, 100),
		tokenMap: tm,
	}
}
