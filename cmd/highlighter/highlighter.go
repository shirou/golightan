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
	pf, err := lexer.LexerFactory(lex)
	if err != nil {
		return err
	}

	buf, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	input := antlr.NewInputStream(string(buf))

	tokens, err := pf.Tokenize(input)
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
