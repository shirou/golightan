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

		l := golang.NewGolangLexer(input)
		stream := antlr.NewCommonTokenStream(l, 0)
		p := golang.NewGolangParser(stream)
		gl := GolangLexer{}
		listener := NewGolangParseTreeListener(p, gl)

		tree := p.SourceFile()

		antlr.ParseTreeWalkerDefault.Walk(listener, tree)
		fmt.Println(listener.GetTokens())
	}
}
