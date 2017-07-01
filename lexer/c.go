package lexer

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"

	c "github.com/shirou/antlr-grammars-v4-go/c"

	"github.com/shirou/golightan"
)

type CLexer struct {
	lexer    antlr.Lexer
	tokenMap TokenMap
}

func (l CLexer) Tokenize(input antlr.CharStream) (golightan.Tokens, error) {
	le := c.NewCLexer(input)
	stream := antlr.NewCommonTokenStream(le, antlr.TokenDefaultChannel)
	p := c.NewCParser(stream)

	// TODO: error handling
	//	p.SetErrorHandler(golightan.NewNullErrorStrategy())
	//p.RemoveErrorListeners()

	listener := NewCommonParseTreeListener(l.tokenMap)
	tree := p.PrimaryExpression()

	antlr.ParseTreeWalkerDefault.Walk(listener, tree)
	return listener.GetTokens(), nil
}

func NewCLexer() Lexer {
	return CLexer{
		tokenMap: NewCTokenMap(),
	}
}
