package lexer

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/shirou/golightan"
)

type Lexer interface {
	Tokenize(input antlr.CharStream) (golightan.Tokens, error)
}

type AvailableLexer struct {
	Targets       []string
	Lexer         string
	Description   string
	Exts          []string
	FactoryMethod func() Lexer
}

func AvailableLexers() []AvailableLexer {
	return []AvailableLexer{
		AvailableLexer{
			Targets:       []string{"sqlite3", "sqlite"},
			Lexer:         "sqlite3",
			Description:   "sqlite3con",
			Exts:          []string{".sql"},
			FactoryMethod: NewSQLiteLexer,
		},
		AvailableLexer{
			Targets:       []string{"go", "golang"},
			Lexer:         "golang",
			Description:   "golang",
			Exts:          []string{".go"},
			FactoryMethod: NewGolangLexer,
		},
		AvailableLexer{
			Targets:       []string{"json"},
			Lexer:         "json",
			Description:   "json",
			Exts:          []string{".json"},
			FactoryMethod: NewJSONLexer,
		},
		AvailableLexer{
			Targets:       []string{"python3"},
			Lexer:         "python3",
			Description:   "python3",
			Exts:          []string{".py"},
			FactoryMethod: NewPython3Lexer,
		},
		AvailableLexer{
			Targets:       []string{"graphql"},
			Lexer:         "graphql",
			Description:   "graphql",
			Exts:          []string{".graphql"},
			FactoryMethod: NewGraphQLLexer,
		},
		AvailableLexer{
			Targets:       []string{"xml"},
			Lexer:         "xml",
			Description:   "xml",
			Exts:          []string{".xml"},
			FactoryMethod: NewXMLLexer,
		},
		AvailableLexer{
			Targets:       []string{"c"},
			Lexer:         "c",
			Description:   "c",
			Exts:          []string{".c"},
			FactoryMethod: NewCLexer,
		},
	}
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func Factory(target string) (Lexer, error) {
	for _, al := range AvailableLexers() {
		if contains(al.Targets, target) {
			return al.FactoryMethod(), nil
		}
	}
	return nil, fmt.Errorf("target %s not found", target)
}

// CommonTokenize use
func CommonTokenize(lexer antlr.Lexer, tm TypeMap) (golightan.Tokens, error) {
	stream := NewAllTokenStream(lexer)

	// Get All tokens
	num := 0
	for ; stream.Sync(num); num++ {
	}

	tokens := make(golightan.Tokens, num)

	for i, token := range stream.GetAllTokens() {
		t := token.GetTokenType()
		if t < 0 {
			break
		}
		tokens[i] = golightan.Token{
			OriginalToken: token,
			TokenType:     tm.Get(t),
			Text:          token.GetText(),
		}
	}
	return tokens, nil
}
