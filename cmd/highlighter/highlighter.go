package main

import (
	"io"
	"io/ioutil"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/shirou/highlighter/formatter"
	"github.com/shirou/highlighter/lexer"
)

const bufferSize = 655000

func HighLight(lex, ft, style string, filters []string, r io.Reader, w io.Writer) error {
	buf, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	input := antlr.NewInputStream(string(buf))

	pf, tm, err := lexer.LexerFactory(lex, input)
	if err != nil {
		return err
	}

	tokens, err := lexer.Tokenize(pf, tm)
	if err != nil {
		return err
	}

	ff, err := formatter.FormatterFactory(ft, style)
	if err != nil {
		return err
	}

	ff.FormatTokens(w, tokens)

	return nil
}
