package lexer

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	json "github.com/shirou/antlr-grammars-v4-go/json"

	"github.com/shirou/highlighter"
)

type JSONLexer struct {
	lexer antlr.Lexer
}

func (l JSONLexer) Tokenize(input antlr.CharStream) (highlighter.Tokens, error) {
	le := json.NewJSONLexer(input)
	//	return CommonTokenize(le, l.symbolicMap)
	fmt.Println(le)
	tokens := make(highlighter.Tokens, 0)
	return tokens, nil
}

func NewJSONLexer() Lexer {
	/*
		symbolicMap := TypeMap{
			json.JSONLexerSTRING: highlighter.TokenTypeNameClass,
			json.JSONLexerNUMBER: highlighter.TokenTypeNameClass,
			json.JSONLexerWS:     highlighter.TokenTypeNameClass,
		}
	*/
	return Python3Lexer{}
}
