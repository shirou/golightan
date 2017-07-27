package main

import (
	"io"

	"github.com/shirou/golightan/formatter"
	"github.com/shirou/golightan/lexer"
)

const bufferSize = 655000

func HighLight(lex, ft, style string, filters []string, r io.Reader, w io.Writer) error {
	pf, err := lexer.Factory(lex)
	if err != nil {
		return err
	}
	scanner := lexer.NewScanner(r, pf)
	tokens, err := scanner.Tokenize(r)
	if err != nil {
		return err
	}

	ff, err := formatter.Factory(ft, style)
	if err != nil {
		return err
	}

	ff.FormatTokens(w, tokens)

	return nil
}
