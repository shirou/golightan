package main

import (
	"fmt"
	"strings"
	"text/scanner"
)

func main() {
	const src = `
"""Compatibility helpers for the different Python versions."""

import sys

PY34 = sys.version_info >= (3, 4)
PY35 = sys.version_info >= (3, 5)
PY352 = sys.version_info >= (3, 5, 2)


def flatten_list_bytes(list_of_data):
    """Concatenate a sequence of bytes-like objects."""
    if not PY34:
        # On Python 3.3 and older, bytes.join() doesn't handle
        # memoryview.
        list_of_data = (
            bytes(data) if isinstance(data, memoryview) else data
            for data in list_of_data)
    return b''.join(list_of_data)
`
	var s scanner.Scanner
	s.Filename = "example"
	s.Init(strings.NewReader(src))
	var tok rune
	for tok != scanner.EOF {
		tok = s.Scan()
		fmt.Println("At position", s.Pos(), ":", s.TokenText())
	}
}
