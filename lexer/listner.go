package lexer

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/shirou/highlighter"
)

type CommonParseTreeListener struct {
	tokens   highlighter.Tokens
	tokenMap TokenMap
	rule     int
}

func (b *CommonParseTreeListener) GetTokens() highlighter.Tokens { return b.tokens }

func (b *CommonParseTreeListener) VisitTerminal(node antlr.TerminalNode) {
	token := node.GetSymbol()
	t := token.GetTokenType()
	if t < 0 {
		return
	}

	text := node.GetText()
	new := highlighter.Token{
		OriginalToken: token,
		TokenType:     b.tokenMap.Convert(b.rule, t, text),
		Text:          text,
	}

	//	fmt.Println("terminal", b.rule, t, node.GetText())

	b.tokens = append(b.tokens, new)
}
func (b *CommonParseTreeListener) VisitErrorNode(node antlr.ErrorNode) {
	token := node.GetSymbol()
	t := token.GetTokenType()
	if t < 0 {
		return
	}
	text := node.GetText()
	new := highlighter.Token{
		OriginalToken: token,
		TokenType:     b.tokenMap.Convert(b.rule, t, text),
		Text:          text,
	}

	//	fmt.Println("error", b.rule, t, node.GetText())

	b.tokens = append(b.tokens, new)
}
func (b *CommonParseTreeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	b.rule = ctx.GetRuleIndex()
}
func (b *CommonParseTreeListener) ExitEveryRule(ctx antlr.ParserRuleContext) {
	b.rule = -1
}

func NewCommonParseTreeListener(tm TokenMap) *CommonParseTreeListener {
	return &CommonParseTreeListener{
		tokens:   make(highlighter.Tokens, 0, 100),
		tokenMap: tm,
	}
}
