package formatter

import (
	"fmt"
	"io"

	"github.com/mgutz/ansi"

	"github.com/shirou/golightan"
)

type TerminalFormat struct {
	terms map[golightan.TokenType]func(string) string
}

func NewTerminalFormat() TerminalFormat {
	m := map[golightan.TokenType]func(string) string{
		golightan.TokenTypeText:                 nil,
		golightan.TokenTypeWhitespace:           nil,
		golightan.TokenTypeError:                ansi.ColorFunc("red+h:black"),
		golightan.TokenTypeOther:                nil,
		golightan.TokenTypeKeyword:              ansi.ColorFunc("green:black"),
		golightan.TokenTypeKeywordConstant:      ansi.ColorFunc("green:black"),
		golightan.TokenTypeKeywordDeclaration:   ansi.ColorFunc("green:black"),
		golightan.TokenTypeKeywordNamespace:     ansi.ColorFunc("green:black"),
		golightan.TokenTypeKeywordPseudo:        nil,
		golightan.TokenTypeKeywordReserved:      nil,
		golightan.TokenTypeKeywordType:          nil,
		golightan.TokenTypeName:                 nil,
		golightan.TokenTypeNameAttribute:        ansi.ColorFunc("yellow:black"),
		golightan.TokenTypeNameBuiltin:          nil,
		golightan.TokenTypeNameBuiltinPseudo:    nil,
		golightan.TokenTypeNameClass:            ansi.ColorFunc("blue+h:black"),
		golightan.TokenTypeNameConstant:         nil,
		golightan.TokenTypeNameDecorator:        nil,
		golightan.TokenTypeNameEntity:           nil,
		golightan.TokenTypeNameException:        nil,
		golightan.TokenTypeNameFunction:         nil,
		golightan.TokenTypeNameProperty:         nil,
		golightan.TokenTypeNameLabel:            nil,
		golightan.TokenTypeNameNamespace:        nil,
		golightan.TokenTypeNameOther:            nil,
		golightan.TokenTypeNameTag:              ansi.ColorFunc("green:black"),
		golightan.TokenTypeNameVariable:         nil,
		golightan.TokenTypeNameVariableClass:    nil,
		golightan.TokenTypeNameVariableGlobal:   nil,
		golightan.TokenTypeNameVariableInstance: nil,
		golightan.TokenTypeLiteral:              nil,
		golightan.TokenTypeLiteralDate:          nil,
		golightan.TokenTypeString:               ansi.ColorFunc("red:black"),
		golightan.TokenTypeStringBacktick:       nil,
		golightan.TokenTypeStringChar:           ansi.ColorFunc("red:black"),
		golightan.TokenTypeStringDoc:            ansi.ColorFunc("red:black"),
		golightan.TokenTypeStringDouble:         ansi.ColorFunc("red:black"),
		golightan.TokenTypeStringEscape:         nil,
		golightan.TokenTypeStringHeredoc:        nil,
		golightan.TokenTypeStringInterpol:       nil,
		golightan.TokenTypeStringOther:          nil,
		golightan.TokenTypeStringRegex:          nil,
		golightan.TokenTypeStringSingle:         nil,
		golightan.TokenTypeStringSymbol:         nil,
		golightan.TokenTypeNumber:               nil,
		golightan.TokenTypeNumberFloat:          nil,
		golightan.TokenTypeNumberHex:            nil,
		golightan.TokenTypeNumberInteger:        ansi.ColorFunc("black+h:black"),
		golightan.TokenTypeNumberIntegerLong:    nil,
		golightan.TokenTypeNumberOct:            nil,
		golightan.TokenTypeOperator:             ansi.ColorFunc("black+h:black"),
		golightan.TokenTypeOperatorWord:         ansi.ColorFunc("blue+h:black"),
		golightan.TokenTypePunctuation:          nil,
		golightan.TokenTypeComment:              ansi.ColorFunc("cyan:black"),
		golightan.TokenTypeCommentMultiline:     nil,
		golightan.TokenTypeCommentPreproc:       nil,
		golightan.TokenTypeCommentSingle:        nil,
		golightan.TokenTypeCommentSpecial:       nil,
		golightan.TokenTypeGeneric:              nil,
		golightan.TokenTypeGenericDeleted:       nil,
		golightan.TokenTypeGenericEmph:          nil,
		golightan.TokenTypeGenericError:         nil,
		golightan.TokenTypeGenericHeading:       nil,
		golightan.TokenTypeGenericInserted:      nil,
		golightan.TokenTypeGenericOutput:        nil,
		golightan.TokenTypeGenericPrompt:        nil,
		golightan.TokenTypeGenericStrong:        nil,
		golightan.TokenTypeGenericSubheading:    nil,
		golightan.TokenTypeGenericTraceback:     nil,
	}

	return TerminalFormat{
		terms: m,
	}

}

func (f TerminalFormat) FormatTokens(w io.Writer, tokens golightan.Tokens) {
	for _, token := range tokens {
		f.Format(w, token)
	}
}

func (f TerminalFormat) Format(w io.Writer, token golightan.Token) {
	cf, ok := f.terms[token.TokenType]
	if ok && cf != nil {
		fmt.Fprintf(w, cf(token.Text))
	} else {
		fmt.Fprintf(w, token.Text)
	}

}
