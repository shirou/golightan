package formatter

import (
	"fmt"
	"io"

	"github.com/shirou/golightan"
)

var tokenTypeNameMap = map[golightan.TokenType]string{
	golightan.TokenTypeText:                 "Token.Text",
	golightan.TokenTypeWhitespace:           "Token.Whitespace",
	golightan.TokenTypeError:                "Token.Error",
	golightan.TokenTypeOther:                "Token.Other",
	golightan.TokenTypeKeyword:              "Token.Keyword",
	golightan.TokenTypeKeywordConstant:      "Token.Keyword.Constant",
	golightan.TokenTypeKeywordDeclaration:   "Token.Keyword.Declaration",
	golightan.TokenTypeKeywordNamespace:     "Token.Keyword.Namespace",
	golightan.TokenTypeKeywordPseudo:        "Token.Keyword.Pseudo",
	golightan.TokenTypeKeywordReserved:      "Token.Keyword.Reserved",
	golightan.TokenTypeKeywordType:          "Token.Keyword.Type",
	golightan.TokenTypeName:                 "Token.Name",
	golightan.TokenTypeNameAttribute:        "Token.Name.Attribute",
	golightan.TokenTypeNameBuiltin:          "Token.Name.Builtin",
	golightan.TokenTypeNameBuiltinPseudo:    "Token.Name.Builtin.Pseudo",
	golightan.TokenTypeNameClass:            "Token.Name.Class",
	golightan.TokenTypeNameConstant:         "Token.Name.Constant",
	golightan.TokenTypeNameDecorator:        "Token.Name.Decorator",
	golightan.TokenTypeNameEntity:           "Token.Name.Entity",
	golightan.TokenTypeNameException:        "Token.Name.Exception",
	golightan.TokenTypeNameFunction:         "Token.Name.Function",
	golightan.TokenTypeNameProperty:         "Token.Name.Property",
	golightan.TokenTypeNameLabel:            "Token.Name.Label",
	golightan.TokenTypeNameNamespace:        "Token.Name.Namespace",
	golightan.TokenTypeNameOther:            "Token.Name.Other",
	golightan.TokenTypeNameTag:              "Token.Name.Tag",
	golightan.TokenTypeNameVariable:         "Token.Name.Variable",
	golightan.TokenTypeNameVariableClass:    "Token.Name.Variable.Class",
	golightan.TokenTypeNameVariableGlobal:   "Token.Name.Variable.Global",
	golightan.TokenTypeNameVariableInstance: "Token.Name.Variable.Instance",
	golightan.TokenTypeLiteral:              "Token.Literal",
	golightan.TokenTypeLiteralDate:          "Token.Literal.Date",
	golightan.TokenTypeString:               "Token.String",
	golightan.TokenTypeStringBacktick:       "Token.String.Backtick",
	golightan.TokenTypeStringChar:           "Token.String.Char",
	golightan.TokenTypeStringDoc:            "Token.String.Doc",
	golightan.TokenTypeStringDouble:         "Token.String.Double",
	golightan.TokenTypeStringEscape:         "Token.String.Escape",
	golightan.TokenTypeStringHeredoc:        "Token.String.Heredoc",
	golightan.TokenTypeStringInterpol:       "Token.String.Interpol",
	golightan.TokenTypeStringOther:          "Token.String.Other",
	golightan.TokenTypeStringRegex:          "Token.String.Regex",
	golightan.TokenTypeStringSingle:         "Token.String.Single",
	golightan.TokenTypeStringSymbol:         "Token.String.Symbol",
	golightan.TokenTypeNumber:               "Token.Number",
	golightan.TokenTypeNumberFloat:          "Token.Number.Float",
	golightan.TokenTypeNumberHex:            "Token.Number.Hex",
	golightan.TokenTypeNumberInteger:        "Token.Number.Integer",
	golightan.TokenTypeNumberIntegerLong:    "Token.Number.Integer.Long",
	golightan.TokenTypeNumberOct:            "Token.Number.Oct",
	golightan.TokenTypeOperator:             "Token.Operator",
	golightan.TokenTypeOperatorWord:         "Token.Operator.Word",
	golightan.TokenTypePunctuation:          "Token.Punctuation",
	golightan.TokenTypeComment:              "Token.Comment",
	golightan.TokenTypeCommentMultiline:     "Token.Comment.Multiline",
	golightan.TokenTypeCommentPreproc:       "Token.Comment.Preproc",
	golightan.TokenTypeCommentSingle:        "Token.Comment.Single",
	golightan.TokenTypeCommentSpecial:       "Token.Comment.Special",
	golightan.TokenTypeGeneric:              "Token.Generic",
	golightan.TokenTypeGenericDeleted:       "Token.Generic.Deleted",
	golightan.TokenTypeGenericEmph:          "Token.Generic.Emph",
	golightan.TokenTypeGenericError:         "Token.Generic.Error",
	golightan.TokenTypeGenericHeading:       "Token.Generic.Heading",
	golightan.TokenTypeGenericInserted:      "Token.Generic.Inserted",
	golightan.TokenTypeGenericOutput:        "Token.Generic.Output",
	golightan.TokenTypeGenericPrompt:        "Token.Generic.Prompt",
	golightan.TokenTypeGenericStrong:        "Token.Generic.Strong",
	golightan.TokenTypeGenericSubheading:    "Token.Generic.Subheading",
	golightan.TokenTypeGenericTraceback:     "Token.Generic.Traceback",
}

type RawTokenFormat struct {
}

func NewRawTokenFormat() RawTokenFormat {
	return RawTokenFormat{}
}

func (f RawTokenFormat) FormatTokens(w io.Writer, tokens golightan.Tokens) {
	for _, token := range tokens {
		f.Format(w, token)
		w.Write([]byte("\n"))
	}
}

func (f RawTokenFormat) Format(w io.Writer, token golightan.Token) {
	fmt.Fprintf(w, "%s: %#v", tokenTypeNameMap[token.TokenType], token.Text)
}
