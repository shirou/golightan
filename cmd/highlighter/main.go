package main

import (
	"flag"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/shirou/highlighter/lexer"
)

type filterSlice []string

var version string

func main() {
	var lexer string
	var filter string
	var formatter string
	var style string
	var out string

	flag.StringVar(&lexer, "l", "", "lexer")
	flag.StringVar(&formatter, "f", "html", "formatter")
	flag.StringVar(&style, "s", "default", "html style")
	flag.StringVar(&filter, "F", "", "filters")
	flag.StringVar(&out, "o", "-", "output")

	flag.Parse()

	var output io.Writer
	var err error
	if out == "-" {
		output = os.Stdout
	}

	var filters []string // TODO

	for _, in := range flag.Args() {
		var input io.Reader

		if in == "-" {
			if lexer == "" {
				log.Fatalf("specify lexer if read from stdin")
			}
			input = os.Stdin
		} else {
			input, err = os.Open(in)
			if err != nil {
				log.Fatalf("open failed: %s", err)
			}
			if lexer == "" {
				lexer = guessByExt(in)
			}
		}
		if lexer == "" {
			log.Fatalf("lexer is not specified or could not guess")
		}
		if err := HighLight(lexer, formatter, style, filters, input, output); err != nil {
			log.Fatalf("highlight failed: %s", err)
		}
	}

}

func guessByExt(filename string) string {
	ext := filepath.Ext(filename)

	for _, p := range lexer.AvailableLexers() {
		for _, e := range p.Exts {
			if e == ext {
				return p.Lexer
			}
		}
	}
	return ""
}
