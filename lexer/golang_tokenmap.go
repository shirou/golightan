package lexer

import (
	golang "github.com/shirou/antlr-grammars-v4-go/golang"

	"github.com/shirou/golightan"
)

func NewGolangTokenMap() TokenMap {
	return TokenMap{
		symbolicMap: newGolangSymbolicMap(),
		ruleMap:     newGolangRuleMap(),
		keywordMap:  newGolangKeywordMap(),
	}
}

func newGolangKeywordMap() map[string]golightan.TokenType {
	return map[string]golightan.TokenType{
		"error":       golightan.TokenTypeError,
		"break":       golightan.TokenTypeKeyword,
		"default":     golightan.TokenTypeKeyword,
		"func":        golightan.TokenTypeKeyword,
		"interface":   golightan.TokenTypeKeyword,
		"select":      golightan.TokenTypeKeyword,
		"case":        golightan.TokenTypeKeyword,
		"defer":       golightan.TokenTypeKeyword,
		"go":          golightan.TokenTypeKeyword,
		"map":         golightan.TokenTypeKeyword,
		"struct":      golightan.TokenTypeKeyword,
		"chan":        golightan.TokenTypeKeyword,
		"else":        golightan.TokenTypeKeyword,
		"goto":        golightan.TokenTypeKeyword,
		"package":     golightan.TokenTypeKeyword,
		"switch":      golightan.TokenTypeKeyword,
		"const":       golightan.TokenTypeKeyword,
		"fallthrough": golightan.TokenTypeKeyword,
		"if":          golightan.TokenTypeKeyword,
		"range":       golightan.TokenTypeKeyword,
		"type":        golightan.TokenTypeKeyword,
		"continue":    golightan.TokenTypeKeyword,
		"for":         golightan.TokenTypeKeyword,
		"import":      golightan.TokenTypeKeyword,
		"return":      golightan.TokenTypeKeyword,
		"var":         golightan.TokenTypeKeyword,
		"nil":         golightan.TokenTypeKeyword,
		":=":          golightan.TokenTypeOperator,
		"||":          golightan.TokenTypeOperator,
		"&&":          golightan.TokenTypeOperator,
		"==":          golightan.TokenTypeOperator,
		"!=":          golightan.TokenTypeOperator,
		"<":           golightan.TokenTypeOperator,
		"<=":          golightan.TokenTypeOperator,
		">":           golightan.TokenTypeOperator,
		">=":          golightan.TokenTypeOperator,
		"|":           golightan.TokenTypeOperator,
		"^":           golightan.TokenTypeOperator,
		"*":           golightan.TokenTypeOperator,
		"/":           golightan.TokenTypeOperator,
		"%":           golightan.TokenTypeOperator,
		"<<":          golightan.TokenTypeOperator,
		">>":          golightan.TokenTypeOperator,
		"&":           golightan.TokenTypeOperator,
		"&^":          golightan.TokenTypeOperator,
		"+":           golightan.TokenTypeOperator,
		"-":           golightan.TokenTypeOperator,
		"!":           golightan.TokenTypeOperator,
		"<-":          golightan.TokenTypeOperator,
	}
}

func newGolangRuleMap() RuleMap {
	return RuleMap{}
}

func newGolangSymbolicMap() TypeMap {
	return TypeMap{
		golang.GolangParserKEYWORD:     golightan.TokenTypeKeyword,
		golang.GolangParserBINARY_OP:   golightan.TokenTypeOperator,
		golang.GolangParserCOMMENT:     golightan.TokenTypeComment,
		golang.GolangLexerLINE_COMMENT: golightan.TokenTypeComment,
		golang.GolangLexerSTRING_LIT:   golightan.TokenTypeStringDouble,
	}
}
