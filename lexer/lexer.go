package lexer

import (
	"fmt"
	"io"

	"github.com/shirou/golightan"
	"github.com/shirou/golightan/lexer/languages"
)

type Lexer interface {
	Tokenize(r io.Reader) (golightan.Tokens, error)
}

type AvailableLexer struct {
	Targets       []string
	Lexer         string
	Description   string
	Exts          []string
	FactoryMethod func() languages.Languager
}

func AvailableLexers() []AvailableLexer {
	return []AvailableLexer{
		AvailableLexer{
			Targets:       []string{"python3"},
			Lexer:         "python3",
			Description:   "python3",
			Exts:          []string{".py"},
			FactoryMethod: languages.NewPython3Language,
		},
		/*
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
			AvailableLexer{
				Targets:       []string{"ruby"},
				Lexer:         "ruby",
				Description:   "ruby",
				Exts:          []string{".rb"},
				FactoryMethod: NewCLexer,
			},
		*/
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

func Factory(target string) (languages.Languager, error) {
	for _, al := range AvailableLexers() {
		if contains(al.Targets, target) {
			return al.FactoryMethod(), nil
		}
	}
	return nil, fmt.Errorf("target %s not found", target)
}
