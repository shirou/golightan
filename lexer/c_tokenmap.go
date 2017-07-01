package lexer

import (
	c "github.com/shirou/antlr-grammars-v4-go/c"

	"github.com/shirou/golightan"
)

func NewCTokenMap() TokenMap {
	return TokenMap{
		symbolicMap: newCSymbolicMap(),
		ruleMap:     newCRuleMap(),
		keywordMap:  newCKeywordMap(),
	}
}

func newCKeywordMap() map[string]golightan.TokenType {
	return map[string]golightan.TokenType{
		"NULL": golightan.TokenTypeKeyword,
	}
}

func newCRuleMap() RuleMap {
	return RuleMap{}
}

func newCSymbolicMap() TypeMap {
	return TypeMap{
		c.CParserInt:          golightan.TokenTypeKeywordType,
		c.CParserLong:         golightan.TokenTypeKeywordType,
		c.CParserFloat:        golightan.TokenTypeKeywordType,
		c.CParserShort:        golightan.TokenTypeKeywordType,
		c.CParserVoid:         golightan.TokenTypeKeywordType,
		c.CParserBool:         golightan.TokenTypeKeywordType,
		c.CParserLess:         golightan.TokenTypeOperator,
		c.CParserLessEqual:    golightan.TokenTypeOperator,
		c.CParserGreater:      golightan.TokenTypeOperator,
		c.CParserGreaterEqual: golightan.TokenTypeOperator,
		c.CParserLeftShift:    golightan.TokenTypeOperator,
		c.CParserRightShift:   golightan.TokenTypeOperator,
		c.CParserPlus:         golightan.TokenTypeOperator,
		c.CParserPlusPlus:     golightan.TokenTypeOperator,
		c.CParserMinus:        golightan.TokenTypeOperator,
		c.CParserMinusMinus:   golightan.TokenTypeOperator,
		c.CParserStar:         golightan.TokenTypeOperator,
		c.CParserDiv:          golightan.TokenTypeOperator,
		c.CParserMod:          golightan.TokenTypeOperator,
		c.CParserAnd:          golightan.TokenTypeOperator,
		c.CParserOr:           golightan.TokenTypeOperator,
		c.CParserAndAnd:       golightan.TokenTypeOperator,
		c.CParserOrOr:         golightan.TokenTypeOperator,
		c.CParserEqual:        golightan.TokenTypeOperator,
		c.CParserAssign:       golightan.TokenTypeOperator,
		c.CParserBlockComment: golightan.TokenTypeComment,
		c.CParserLineComment:  golightan.TokenTypeCommentSingle,
		//		c.CParserIdentifier:    golightan.TokenTypeNameFunction,
		c.CParserFor:           golightan.TokenTypeKeyword,
		c.CParserReturn:        golightan.TokenTypeKeyword,
		c.CParserBreak:         golightan.TokenTypeKeyword,
		c.CParserConst:         golightan.TokenTypeKeyword,
		c.CParserContinue:      golightan.TokenTypeKeyword,
		c.CParserDefault:       golightan.TokenTypeKeyword,
		c.CParserSizeof:        golightan.TokenTypeKeyword,
		c.CParserDo:            golightan.TokenTypeKeyword,
		c.CParserIf:            golightan.TokenTypeKeyword,
		c.CParserStruct:        golightan.TokenTypeKeyword,
		c.CParserTypedef:       golightan.TokenTypeKeyword,
		c.CParserUnion:         golightan.TokenTypeKeyword,
		c.CParserElse:          golightan.TokenTypeKeyword,
		c.CParserStringLiteral: golightan.TokenTypeString,
	}
}
