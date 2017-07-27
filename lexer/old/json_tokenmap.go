package languages

import (
	json "github.com/shirou/antlr-grammars-v4-go/json"

	"github.com/shirou/golightan"
)

func NewJSONTokenMap() TokenMap {
	return TokenMap{
		symbolicMap: newJSONSymbolicMap(),
		ruleMap:     newJSONRuleMap(),
	}
}

func newJSONRuleMap() RuleMap {
	return RuleMap{
		Rule{json.JSONParserRULE_pair, json.JSONParserSTRING}:  golightan.TokenTypeNameTag,
		Rule{json.JSONParserRULE_pair, json.JSONParserT__3}:    golightan.TokenTypePunctuation,
		Rule{json.JSONParserRULE_obj, json.JSONParserT__0}:     golightan.TokenTypePunctuation,
		Rule{json.JSONParserRULE_value, json.JSONParserSTRING}: golightan.TokenTypeStringDouble,
	}
}

func newJSONSymbolicMap() TypeMap {
	return TypeMap{
		json.JSONLexerSTRING:     golightan.TokenTypeStringDouble,
		json.JSONLexerKEY_STRING: golightan.TokenTypeNameTag,
		json.JSONLexerNUMBER:     golightan.TokenTypeNumber,
	}
}
