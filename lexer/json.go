package lexer

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"

	json "github.com/shirou/antlr-grammars-v4-go/json"

	"github.com/shirou/highlighter"
)

type JSONLexer struct {
	lexer       antlr.Lexer
	ruleMap     TypeMap
	literalMap  TypeMap
	symbolicMap TypeMap
}

func (l JSONLexer) Tokenize(input antlr.CharStream) (highlighter.Tokens, error) {
	le := json.NewJSONLexer(input)
	return CommonTokenize(le, l.symbolicMap)
}

func NewJSONLexer() Lexer {
	symbolicMap := TypeMap{
		json.JSONLexerSTRING: highlighter.TokenTypeNameClass,
		json.JSONLexerNUMBER: highlighter.TokenTypeNameClass,
		json.JSONLexerWS:     highlighter.TokenTypeNameClass,
	}
	return Python3Lexer{
		symbolicMap: symbolicMap,
	}
}
