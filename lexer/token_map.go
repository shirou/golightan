package lexer

import (
	"github.com/shirou/highlighter"
)

type TypeMap map[int]highlighter.TokenType

const (
	InvalidToken = -1
)

func (tm TypeMap) Get(tokenType int) highlighter.TokenType {
	s, ok := tm[tokenType]
	if !ok {
		return InvalidToken
	}
	return s
}

type Rule [2]int // 0: rule, 1: antlr node token type
type RuleMap map[Rule]highlighter.TokenType

func (t RuleMap) Get(rule, tokenType int) highlighter.TokenType {
	s, ok := t[Rule{rule, tokenType}]
	if !ok {
		return InvalidToken
	}
	return s
}

type TokenMap struct {
	ruleMap     RuleMap
	symbolicMap TypeMap
	keywordMap  map[string]highlighter.TokenType
}

// Convert converts from rule and antlr TokenType to highlighter.TokenType.
// 1. search Keyword Map
// 2. search RuleMap using rule and tokentype
// 3. If not in RuleMap, search symbolicMap using tokenType
// 4. If not , return TokenTypeText as normal text
func (t TokenMap) Convert(rule, tokenType int, text string) highlighter.TokenType {
	if tmp, ok := t.keywordMap[text]; ok {
		return tmp
	}

	if tmp := t.ruleMap.Get(rule, tokenType); tmp != InvalidToken {
		return tmp
	}
	if tmp := t.symbolicMap.Get(tokenType); tmp != InvalidToken {
		return tmp
	}
	return highlighter.TokenTypeText

}