package lexer

import (
	"path/filepath"
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type TestCase struct {
	src, exp string
}

func Run(t *testing.T, tests []TestCase, target string) {
	for _, test := range tests {
		input, err := antlr.NewFileStream(filepath.Join("testcase", test.src))
		if err != nil {
			t.Fatal(err)
		}

		lexer, tm, err := LexerFactory(target, input)
		if err != nil {
			t.Fatal(err)
		}
		ret, err := Tokenize(lexer, tm)
		if err != nil {
			t.Fatal(err)
		}
		if len(ret) == 0 {
			t.Fatal("tokenize failed")
		}
	}
}

func TestSQLite(t *testing.T) {
	files := []TestCase{
		TestCase{"sqlite/input-1.sql", "sqlite/output-1.sql"},
		TestCase{"sqlite/input-2.sql", "sqlite/output-2.sql"},
	}
	Run(t, files, "sql")
}

func TestGolang(t *testing.T) {
	files := []TestCase{
		TestCase{"golang/example.go", "golang/output-1.sql"},
		TestCase{"golang/ill_but_correct.go", "golang/output-2.sql"},
	}
	Run(t, files, "golang")
}
