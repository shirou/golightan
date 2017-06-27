package lexer

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/shirou/highlighter/formatter"
)

func TestPython3(t *testing.T) {
	tests := []TestCase{
		TestCase{"python3/__init__.py", "python3/raw-__init__.py"},
	}

	for _, test := range tests {
		input, err := antlr.NewFileStream(filepath.Join("testcase", test.src))
		if err != nil {
			t.Fatal(err)
		}
		l, tm, err := LexerFactory("python3", input)
		if err != nil {
			t.Fatalf("LexerFactory failed :%s", err)
		}
		tokens, err := Tokenize(l, tm)
		if err != nil {
			t.Fatalf("tokens :%s", err)
		}

		ff, err := formatter.FormatterFactory("raw", "")
		if err != nil {
			t.Fatalf("tokens :%s", err)
		}
		ff.FormatTokens(os.Stdout, tokens)
	}
}
