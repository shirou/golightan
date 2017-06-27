package lexer

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	golang "github.com/shirou/antlr-grammars-v4-go/golang"
)

func TestGolang2(t *testing.T) {
	tests := []TestCase{
		TestCase{"golang/example.go", "golang/raw-example.go"},
	}

	for _, test := range tests {
		input, err := antlr.NewFileStream(filepath.Join("testcase", test.src))
		if err != nil {
			t.Fatal(err)
		}

		lexer := golang.NewGolangLexer(input)
		stream := antlr.NewCommonTokenStream(lexer, 0)
		p := golang.NewGolangLexer(stream)
		listener := NewParseTreeListener()
		p.AddParseListener(listener)

		tree := p.SourceFile()

		fmt.Println(tree.ToStringTree([]string{}, p))
	}
}
