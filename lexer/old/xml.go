package languages

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"

	xml "github.com/shirou/antlr-grammars-v4-go/xml"

	"github.com/shirou/golightan"
	"github.com/shirou/golightan/lexer"
)

type XMLLexer struct {
	lexer    antlr.Lexer
	tokenMap lexer.TokenMap
}

func (l XMLLexer) Tokenize(input antlr.CharStream) (golightan.Tokens, error) {
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
