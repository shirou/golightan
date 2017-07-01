package lexer

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"

	sqlite "github.com/shirou/antlr-grammars-v4-go/sqlite"

	"github.com/shirou/golightan"
)

type SQLiteLexer struct {
	lexer    antlr.Lexer
	tokenMap TokenMap
}

func (l SQLiteLexer) Tokenize(input antlr.CharStream) (golightan.Tokens, error) {
	le := sqlite.NewSQLiteLexer(input)
	stream := antlr.NewCommonTokenStream(le, antlr.TokenDefaultChannel)
	p := sqlite.NewSQLiteParser(stream)

	// TODO: error handling
	p.SetErrorHandler(golightan.NewNullErrorStrategy())
	p.RemoveErrorListeners()

	listener := NewCommonParseTreeListener(l.tokenMap)
	tree := p.Parse()

	antlr.ParseTreeWalkerDefault.Walk(listener, tree)
	return listener.GetTokens(), nil
}

func NewSQLiteLexer() Lexer {
	return SQLiteLexer{
		tokenMap: NewSQLiteTokenMap(),
	}
}
