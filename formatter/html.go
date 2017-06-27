package formatter

import (
	"fmt"
	"io"

	"github.com/shirou/highlighter"
)

type HtmlFormat struct {
	Style string
}

func NewHtmlFormat(style string) HtmlFormat {
	return HtmlFormat{
		Style: style,
	}
}

func (f HtmlFormat) FormatTokens(w io.Writer, tokens highlighter.Tokens) {
	w.Write([]byte(`<div class="highlight"><pre>`))

	for _, token := range tokens {
		f.Format(w, token)
	}
	w.Write([]byte(`</pre></div>`))
}

func (f HtmlFormat) Format(w io.Writer, token highlighter.Token) {
	/*
		if token.TokenType == 0 {
			fmt.Fprint(w, "<span></span>")
			return
		}
	*/

	cls, ok := highlighter.CSSMap[token.TokenType]
	if ok {
		fmt.Fprintf(w, `<span class="%s">%s</span>`, cls, token.Text)
	} else {
		fmt.Fprintf(w, "<span>%s</span>", token.Text)
	}

}
