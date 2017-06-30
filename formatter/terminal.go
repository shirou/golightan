package formatter

import (
	"fmt"
	"io"

	"github.com/mgutz/ansi"

	"github.com/shirou/highlighter"
)

type TerminalFormat struct {
	terms map[highlighter.TokenType]func(string) string
}

func NewTerminalFormat() TerminalFormat {
	m := map[highlighter.TokenType]func(string) string{
		highlighter.TokenTypeText:                 nil,
		highlighter.TokenTypeWhitespace:           nil,
		highlighter.TokenTypeError:                ansi.ColorFunc("red+h:black"),
		highlighter.TokenTypeOther:                nil,
		highlighter.TokenTypeKeyword:              ansi.ColorFunc("green:black"),
		highlighter.TokenTypeKeywordConstant:      ansi.ColorFunc("green:black"),
		highlighter.TokenTypeKeywordDeclaration:   ansi.ColorFunc("green:black"),
		highlighter.TokenTypeKeywordNamespace:     ansi.ColorFunc("green:black"),
		highlighter.TokenTypeKeywordPseudo:        nil,
		highlighter.TokenTypeKeywordReserved:      nil,
		highlighter.TokenTypeKeywordType:          nil,
		highlighter.TokenTypeName:                 nil,
		highlighter.TokenTypeNameAttribute:        nil,
		highlighter.TokenTypeNameBuiltin:          nil,
		highlighter.TokenTypeNameBuiltinPseudo:    nil,
		highlighter.TokenTypeNameClass:            ansi.ColorFunc("blue+h:black"),
		highlighter.TokenTypeNameConstant:         nil,
		highlighter.TokenTypeNameDecorator:        nil,
		highlighter.TokenTypeNameEntity:           nil,
		highlighter.TokenTypeNameException:        nil,
		highlighter.TokenTypeNameFunction:         nil,
		highlighter.TokenTypeNameProperty:         nil,
		highlighter.TokenTypeNameLabel:            nil,
		highlighter.TokenTypeNameNamespace:        nil,
		highlighter.TokenTypeNameOther:            nil,
		highlighter.TokenTypeNameTag:              ansi.ColorFunc("green:black"),
		highlighter.TokenTypeNameVariable:         nil,
		highlighter.TokenTypeNameVariableClass:    nil,
		highlighter.TokenTypeNameVariableGlobal:   nil,
		highlighter.TokenTypeNameVariableInstance: nil,
		highlighter.TokenTypeLiteral:              nil,
		highlighter.TokenTypeLiteralDate:          nil,
		highlighter.TokenTypeString:               nil,
		highlighter.TokenTypeStringBacktick:       nil,
		highlighter.TokenTypeStringChar:           nil,
		highlighter.TokenTypeStringDoc:            nil,
		highlighter.TokenTypeStringDouble:         ansi.ColorFunc("red:black"),
		highlighter.TokenTypeStringEscape:         nil,
		highlighter.TokenTypeStringHeredoc:        nil,
		highlighter.TokenTypeStringInterpol:       nil,
		highlighter.TokenTypeStringOther:          nil,
		highlighter.TokenTypeStringRegex:          nil,
		highlighter.TokenTypeStringSingle:         nil,
		highlighter.TokenTypeStringSymbol:         nil,
		highlighter.TokenTypeNumber:               nil,
		highlighter.TokenTypeNumberFloat:          nil,
		highlighter.TokenTypeNumberHex:            nil,
		highlighter.TokenTypeNumberInteger:        ansi.ColorFunc("black+h:black"),
		highlighter.TokenTypeNumberIntegerLong:    nil,
		highlighter.TokenTypeNumberOct:            nil,
		highlighter.TokenTypeOperator:             ansi.ColorFunc("black+h:black"),
		highlighter.TokenTypeOperatorWord:         ansi.ColorFunc("blue+h:black"),
		highlighter.TokenTypePunctuation:          nil,
		highlighter.TokenTypeComment:              ansi.ColorFunc("cyan:black"),
		highlighter.TokenTypeCommentMultiline:     nil,
		highlighter.TokenTypeCommentPreproc:       nil,
		highlighter.TokenTypeCommentSingle:        nil,
		highlighter.TokenTypeCommentSpecial:       nil,
		highlighter.TokenTypeGeneric:              nil,
		highlighter.TokenTypeGenericDeleted:       nil,
		highlighter.TokenTypeGenericEmph:          nil,
		highlighter.TokenTypeGenericError:         nil,
		highlighter.TokenTypeGenericHeading:       nil,
		highlighter.TokenTypeGenericInserted:      nil,
		highlighter.TokenTypeGenericOutput:        nil,
		highlighter.TokenTypeGenericPrompt:        nil,
		highlighter.TokenTypeGenericStrong:        nil,
		highlighter.TokenTypeGenericSubheading:    nil,
		highlighter.TokenTypeGenericTraceback:     nil,
	}

	return TerminalFormat{
		terms: m,
	}

}

func (f TerminalFormat) FormatTokens(w io.Writer, tokens highlighter.Tokens) {
	for _, token := range tokens {
		f.Format(w, token)
	}
}

func (f TerminalFormat) Format(w io.Writer, token highlighter.Token) {
	cf, ok := f.terms[token.TokenType]
	if ok && cf != nil {
		fmt.Fprintf(w, cf(token.Text))
	} else {
		fmt.Fprintf(w, token.Text)
	}

}
