package lexer

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"

	json "github.com/shirou/antlr-grammars-v4-go/json"

	"github.com/shirou/highlighter"
)

type JsonTypeMap TypeMap

var jsonTypeMap JsonTypeMap

func init() {
	jsonTypeMap = NewJsonTypeMap()
}

func NewJsonTypeMap() JsonTypeMap {
	return JsonTypeMap{
		json.JSONLexerSTRING: highlighter.TokenTypeNameClass,
		json.JSONLexerNUMBER: highlighter.TokenTypeNameClass,
		json.JSONLexerWS:     highlighter.TokenTypeNameClass,
	}
}

func (tm JsonTypeMap) Get(type_ int) highlighter.TokenType {
	s, ok := tm[type_]
	if !ok {
		return 0
	}
	return s
}

func NewJSONLexer(input antlr.CharStream) antlr.Lexer {
	return json.NewJSONLexer(input)
}
