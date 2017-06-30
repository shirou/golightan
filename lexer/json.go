package lexer

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"

	json "github.com/shirou/antlr-grammars-v4-go/json"

	"github.com/shirou/highlighter"
)

type JSONLexer struct {
	lexer    antlr.Lexer
	tokenMap TokenMap
}

func (l JSONLexer) Tokenize(input antlr.CharStream) (highlighter.Tokens, error) {
	le := json.NewJSONLexer(input)
	stream := antlr.NewCommonTokenStream(le, antlr.TokenDefaultChannel)
	p := json.NewJSONParser(stream)

	// TODO: error handling
	p.SetErrorHandler(highlighter.NewNullErrorStrategy())
	p.RemoveErrorListeners()

	listener := NewCommonParseTreeListener(l.tokenMap)
	tree := p.Json()

	antlr.ParseTreeWalkerDefault.Walk(listener, tree)
	return listener.GetTokens(), nil
}

func NewJSONLexer() Lexer {
	return JSONLexer{
		tokenMap: NewJSONTokenMap(),
	}
}
