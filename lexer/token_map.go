package lexer

import (
	"github.com/shirou/golightan"
)

type TypeMap map[int]golightan.TokenType

const (
	InvalidToken = -1
)

func (tm TypeMap) Get(tokenType int) golightan.TokenType {
	s, ok := tm[tokenType]
	if !ok {
		return InvalidToken
	}
	return s
}

type Rule [2]int // 0: rule, 1: antlr node token type
type RuleMap map[Rule]golightan.TokenType

func (t RuleMap) Get(rule, tokenType int) golightan.TokenType {
	s, ok := t[Rule{rule, tokenType}]
	if !ok {
		return InvalidToken
	}
	return s
}

type TokenMap struct {
	ruleMap     RuleMap
	symbolicMap TypeMap
	keywordMap  map[string]golightan.TokenType
}

// Convert converts from rule and antlr TokenType to golightan.TokenType.
// 1. search Keyword Map
// 2. search RuleMap using rule and tokentype
// 3. If not in RuleMap, search symbolicMap using tokenType
// 4. If not , return TokenTypeText as normal text
func (t TokenMap) Convert(rule, tokenType int, text string) golightan.TokenType {
	if tmp, ok := t.keywordMap[text]; ok {
		return tmp
	}

	if tmp := t.ruleMap.Get(rule, tokenType); tmp != InvalidToken {
		return tmp
	}
	if tmp := t.symbolicMap.Get(tokenType); tmp != InvalidToken {
		return tmp
	}
	return golightan.TokenTypeText

}
