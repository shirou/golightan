package formatter

import (
	"fmt"
	"html"
	"io"

	"github.com/shirou/golightan"
)

type HtmlFormat struct {
	Style string
}

func NewHtmlFormat(style string) HtmlFormat {
	return HtmlFormat{
		Style: style,
	}
}

func (f HtmlFormat) FormatTokens(w io.Writer, tokens golightan.Tokens) {
	w.Write([]byte(`<div class="highlight"><pre>`))

	for _, token := range tokens {
		f.Format(w, token)
	}
	w.Write([]byte(`</pre></div>`))
}

func (f HtmlFormat) Format(w io.Writer, token golightan.Token) {
	/*
		if token.TokenType == 0 {
			fmt.Fprint(w, "<span></span>")
			return
		}
	*/

	text := html.EscapeString(token.Text)

	cls, ok := golightan.CSSMap[token.TokenType]
	if ok {
		fmt.Fprintf(w, `<span class="%s">%s</span>`, cls, text)
	} else {
		fmt.Fprintf(w, "<span>%s</span>", text)
	}

}
