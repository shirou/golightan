package lexer

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"

	xml "github.com/shirou/antlr-grammars-v4-go/xml"

	"github.com/shirou/highlighter"
)

type XMLLexer struct {
	lexer    antlr.Lexer
	tokenMap TokenMap
}

func (l XMLLexer) Tokenize(input antlr.CharStream) (highlighter.Tokens, error) {
	le := xml.NewXMLLexer(input)
	stream := antlr.NewCommonTokenStream(le, antlr.TokenDefaultChannel)
	p := xml.NewXMLParser(stream)

	listener := NewCommonParseTreeListener(l.tokenMap)
	tree := p.Document()

	antlr.ParseTreeWalkerDefault.Walk(listener, tree)
	return listener.GetTokens(), nil
}

func NewXMLLexer() Lexer {
	return XMLLexer{
		tokenMap: NewXMLTokenMap(),
	}
}
