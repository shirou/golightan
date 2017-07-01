package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/shirou/golightan/lexer"
)

type filterSlice []string

var version string

func main() {
	var lexer string
	var filter string
	var formatter string
	var style string
	var out string
	var list bool

	flag.StringVar(&lexer, "l", "", "lexer")
	flag.StringVar(&formatter, "f", "html", "formatter")
	flag.StringVar(&style, "s", "default", "html style")
	flag.StringVar(&filter, "F", "", "filters")
	flag.BoolVar(&list, "L", false, "list available lexers and formatters")
	flag.StringVar(&out, "o", "-", "output")

	flag.Parse()

	if list {
		showList()
		return
	}

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

func showList() {
	fmt.Println("# Available Lexers")
	tmpl := `  * %s (%s)
      %s
`
	for _, p := range lexer.AvailableLexers() {
		fmt.Printf(tmpl,
			strings.Join(p.Targets, ", "),
			strings.Join(p.Exts, ", "),
			p.Description)
	}
	fmt.Println("")
	fmt.Println("# Available Formatters")
	f := []string{"html (default)", "terminal", "raw"}
	for _, p := range f {
		fmt.Printf("  * %s\n", p)
	}

}
