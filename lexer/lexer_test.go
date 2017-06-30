package lexer

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/shirou/highlighter/formatter"
)

type TestCase struct {
	src, exp string
}

func Run(t *testing.T, tests []TestCase, target string) {
	ff, _ := formatter.FormatterFactory("terminal", "")

	for _, test := range tests {
		lexer, err := LexerFactory(target)
		if err != nil {
			t.Fatal(err)
		}
		input, err := antlr.NewFileStream(filepath.Join("testcase", test.src))
		if err != nil {
			t.Fatal(err)
		}

		tokens, err := lexer.Tokenize(input)
		if err != nil {
			t.Fatal(err)
		}
		if len(tokens) == 0 {
			t.Fatal("tokenize failed")
		}

		ff.FormatTokens(os.Stdout, tokens)
	}
}

func TestSQLite(t *testing.T) {
	files := []TestCase{
		TestCase{"sqlite/input-1.sql", "sqlite/output-1.sql"},
		TestCase{"sqlite/input-2.sql", "sqlite/output-2.sql"},
	}
	Run(t, files, "sqlite")
}

func TestGolang(t *testing.T) {
	files := []TestCase{
		TestCase{"golang/example.go", "golang/output-1.sql"},
		TestCase{"golang/ill_but_correct.go", "golang/output-2.sql"},
	}
	Run(t, files, "golang")
}

func TestPython3(t *testing.T) {
	files := []TestCase{
		//		TestCase{"python3/__init__.py", "python3/raw-__init__.py"},
		TestCase{"python3/tasks.py", "python3/raw-tasks.py"},
	}
	Run(t, files, "python3")

}
