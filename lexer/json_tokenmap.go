package lexer

import (
	json "github.com/shirou/antlr-grammars-v4-go/json"

	"github.com/shirou/highlighter"
)

func NewJSONTokenMap() TokenMap {
	return TokenMap{
		symbolicMap: newJSONSymbolicMap(),
		ruleMap:     newJSONRuleMap(),
	}
}

func newJSONRuleMap() RuleMap {
	return RuleMap{
		Rule{json.JSONParserRULE_pair, json.JSONParserSTRING}:  highlighter.TokenTypeNameTag,
		Rule{json.JSONParserRULE_pair, json.JSONParserT__3}:    highlighter.TokenTypePunctuation,
		Rule{json.JSONParserRULE_obj, json.JSONParserT__0}:     highlighter.TokenTypePunctuation,
		Rule{json.JSONParserRULE_value, json.JSONParserSTRING}: highlighter.TokenTypeStringDouble,
	}
}

func newJSONSymbolicMap() TypeMap {
	return TypeMap{
		json.JSONParserT__1:  highlighter.TokenTypePunctuation,
		json.JSONLexerSTRING: highlighter.TokenTypeNameClass,
		json.JSONLexerNUMBER: highlighter.TokenTypeNameClass,
		json.JSONLexerWS:     highlighter.TokenTypeNameClass,
	}
}
