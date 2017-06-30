package lexer

import (
	graphql "github.com/shirou/antlr-grammars-v4-go/graphql"

	"github.com/shirou/highlighter"
)

func NewGraphQLTokenMap() TokenMap {
	return TokenMap{
		symbolicMap: newGraphQLSymbolicMap(),
		ruleMap:     newGraphQLRuleMap(),
		keywordMap:  newGraphQLKeywordMap(),
	}
}
func newGraphQLKeywordMap() map[string]highlighter.TokenType {
	return map[string]highlighter.TokenType{
		"errors": highlighter.TokenTypeError,
	}
}

func newGraphQLRuleMap() RuleMap {
	return RuleMap{
		Rule{graphql.GraphQLParserRULE_operationType,
			graphql.GraphQLParserT__3}: highlighter.TokenTypeOperatorWord,
		Rule{graphql.GraphQLParserRULE_argument,
			graphql.GraphQLParserNAME}: highlighter.TokenTypeNameClass,
		Rule{graphql.GraphQLParserRULE_value,
			graphql.GraphQLParserSTRING}: highlighter.TokenTypeStringDouble,
		Rule{graphql.GraphQLParserRULE_fieldName,
			graphql.GraphQLParserNAME}: highlighter.TokenTypeKeyword,
	}
}

func newGraphQLSymbolicMap() TypeMap {
	return TypeMap{
		graphql.GraphQLParserT__1:  highlighter.TokenTypePunctuation,
		graphql.GraphQLLexerSTRING: highlighter.TokenTypeNameClass,
		graphql.GraphQLLexerNUMBER: highlighter.TokenTypeNameClass,
	}
}
