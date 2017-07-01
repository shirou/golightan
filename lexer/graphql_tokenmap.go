package lexer

import (
	graphql "github.com/shirou/antlr-grammars-v4-go/graphql"

	"github.com/shirou/golightan"
)

func NewGraphQLTokenMap() TokenMap {
	return TokenMap{
		symbolicMap: newGraphQLSymbolicMap(),
		ruleMap:     newGraphQLRuleMap(),
		keywordMap:  newGraphQLKeywordMap(),
	}
}
func newGraphQLKeywordMap() map[string]golightan.TokenType {
	return map[string]golightan.TokenType{
		"errors": golightan.TokenTypeError,
	}
}

func newGraphQLRuleMap() RuleMap {
	return RuleMap{
		Rule{graphql.GraphQLParserRULE_operationType,
			graphql.GraphQLParserT__3}: golightan.TokenTypeOperatorWord,
		Rule{graphql.GraphQLParserRULE_argument,
			graphql.GraphQLParserNAME}: golightan.TokenTypeNameClass,
		Rule{graphql.GraphQLParserRULE_value,
			graphql.GraphQLParserSTRING}: golightan.TokenTypeStringDouble,
		Rule{graphql.GraphQLParserRULE_fieldName,
			graphql.GraphQLParserNAME}: golightan.TokenTypeKeyword,
	}
}

func newGraphQLSymbolicMap() TypeMap {
	return TypeMap{
		graphql.GraphQLParserT__1:  golightan.TokenTypePunctuation,
		graphql.GraphQLLexerSTRING: golightan.TokenTypeNameClass,
		graphql.GraphQLLexerNUMBER: golightan.TokenTypeNameClass,
	}
}
