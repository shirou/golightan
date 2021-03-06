package lexer

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"

	golang "github.com/shirou/antlr-grammars-v4-go/golang"

	"github.com/shirou/golightan"
)

type GolangLexer struct {
	lexer    antlr.Lexer
	tokenMap TokenMap
}

func (l GolangLexer) Tokenize(input antlr.CharStream) (golightan.Tokens, error) {
	le := golang.NewGolangLexer(input)
	stream := NewAllTokenStream(le)
	p := golang.NewGolangParser(stream)

	// TODO: error handling
	//p.SetErrorHandler(golightan.NewNullErrorStrategy())
	//p.SetErrorHandler(antlr.NewBailErrorStrategy())
	//p.RemoveErrorListeners()

	listener := NewCommonParseTreeListener(l.tokenMap)
	tree := p.SourceFile()

	// listener.SetDebug(le, p)

	antlr.ParseTreeWalkerDefault.Walk(listener, tree)
	return listener.GetTokens(), nil
}

func NewGolangLexer() Lexer {
	return GolangLexer{
		tokenMap: NewGolangTokenMap(),
	}
}
