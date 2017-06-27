package formatter

import (
	"fmt"
	"io"

	"github.com/shirou/highlighter"
)

type Formatter interface {
	FormatTokens(w io.Writer, tokens highlighter.Tokens)
	Format(w io.Writer, token highlighter.Token)
}

func FormatterFactory(f, style string) (Formatter, error) {
	switch f {
	case "raw":
		return NewRawTokenFormat(), nil
	case "html":
		return NewHtmlFormat(style), nil
	case "terminal":
		return NewTerminalFormat(), nil
	default:
		return nil, fmt.Errorf("unknown formatter: %s", f)
	}
}
