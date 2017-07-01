package formatter

import (
	"fmt"
	"io"

	"github.com/shirou/golightan"
)

type Formatter interface {
	FormatTokens(w io.Writer, tokens golightan.Tokens)
	Format(w io.Writer, token golightan.Token)
}

func Factory(f, style string) (Formatter, error) {
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
