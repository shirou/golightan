package lexer

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"

	graphql "github.com/shirou/antlr-grammars-v4-go/graphql"

	"github.com/shirou/highlighter"
)

type GraphQLLexer struct {
	lexer    antlr.Lexer
	tokenMap TokenMap
}

func (l GraphQLLexer) Tokenize(input antlr.CharStream) (highlighter.Tokens, error) {
	le := graphql.NewGraphQLLexer(input)
	stream := antlr.NewCommonTokenStream(le, antlr.TokenDefaultChannel)
	p := graphql.NewGraphQLParser(stream)

	// TODO: error handling
	p.SetErrorHandler(highlighter.NewNullErrorStrategy())
	p.RemoveErrorListeners()

	listener := NewCommonParseTreeListener(l.tokenMap)
	tree := p.Document()

	antlr.ParseTreeWalkerDefault.Walk(listener, tree)
	return listener.GetTokens(), nil
}

func NewGraphQLLexer() Lexer {
	return GraphQLLexer{
		tokenMap: NewGraphQLTokenMap(),
	}
}
