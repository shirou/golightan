package lexer

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/shirou/highlighter"
)

type TypeMap map[int]highlighter.TokenType
type TokenTypeMapper interface {
	Get(int) highlighter.TokenType
}

type AvailableLexer struct {
	Target      []string
	Lexer       string
	Description string
	Exts        []string
}

func AvailableLexers() []AvailableLexer {
	return []AvailableLexer{
		AvailableLexer{
			Target:      []string{"sqlite3"},
			Lexer:       "sqlite3",
			Description: "sqlite3con",
			Exts:        []string{".sql"},
		},
		AvailableLexer{
			Target:      []string{"go", "golang"},
			Lexer:       "golang",
			Description: "golang",
			Exts:        []string{".go"},
		},
		AvailableLexer{
			Target:      []string{"json"},
			Lexer:       "json",
			Description: "json",
			Exts:        []string{".json"},
		},
		AvailableLexer{
			Target:      []string{"python3"},
			Lexer:       "python3",
			Description: "python3",
			Exts:        []string{".py"},
		},
	}
}

func LexerFactory(target string, input antlr.CharStream) (antlr.Lexer, TokenTypeMapper, error) {
	switch target {
	case "sqlite3":
		return NewSQLiteLexer(input), sqliteTypeMap, nil
	case "go", "golang":
		return NewGolangLexer(input), golangTypeMap, nil
	case "json":
		return NewJSONLexer(input), jsonTypeMap, nil
	case "python3":
		return NewPython3Lexer(input), python3TypeMap, nil
	default:
		return nil, nil, fmt.Errorf("target %s not found", target)
	}
}

func Tokenize(lexer antlr.Lexer, tm TokenTypeMapper) (highlighter.Tokens, error) {
	stream := antlr.NewCommonTokenStream(lexer, 0)

	// Get All tokens
	num := 0
	for ; stream.Sync(num); num++ {
	}

	tokens := make(highlighter.Tokens, num)

	for i, token := range stream.GetAllTokens() {
		t := token.GetTokenType()
		if t < 0 {
			break
		}
		tokens[i] = highlighter.Token{
			OriginalToken: token,
			TokenType:     tm.Get(t),
			Text:          token.GetText(),
		}
	}

	return tokens, nil
}
