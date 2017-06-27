package formatter

import (
	"fmt"
	"io"

	"github.com/shirou/highlighter"
)

var tokenTypeNameMap = map[highlighter.TokenType]string{
	highlighter.TokenTypeText:                 "Token.Text",
	highlighter.TokenTypeWhitespace:           "Token.Whitespace",
	highlighter.TokenTypeError:                "Token.Error",
	highlighter.TokenTypeOther:                "Token.Other",
	highlighter.TokenTypeKeyword:              "Token.Keyword",
	highlighter.TokenTypeKeywordConstant:      "Token.Keyword.Constant",
	highlighter.TokenTypeKeywordDeclaration:   "Token.Keyword.Declaration",
	highlighter.TokenTypeKeywordNamespace:     "Token.Keyword.Namespace",
	highlighter.TokenTypeKeywordPseudo:        "Token.Keyword.Pseudo",
	highlighter.TokenTypeKeywordReserved:      "Token.Keyword.Reserved",
	highlighter.TokenTypeKeywordType:          "Token.Keyword.Type",
	highlighter.TokenTypeName:                 "Token.Name",
	highlighter.TokenTypeNameAttribute:        "Token.Name.Attribute",
	highlighter.TokenTypeNameBuiltin:          "Token.Name.Builtin",
	highlighter.TokenTypeNameBuiltinPseudo:    "Token.Name.Builtin.Pseudo",
	highlighter.TokenTypeNameClass:            "Token.Name.Class",
	highlighter.TokenTypeNameConstant:         "Token.Name.Constant",
	highlighter.TokenTypeNameDecorator:        "Token.Name.Decorator",
	highlighter.TokenTypeNameEntity:           "Token.Name.Entity",
	highlighter.TokenTypeNameException:        "Token.Name.Exception",
	highlighter.TokenTypeNameFunction:         "Token.Name.Function",
	highlighter.TokenTypeNameProperty:         "Token.Name.Property",
	highlighter.TokenTypeNameLabel:            "Token.Name.Label",
	highlighter.TokenTypeNameNamespace:        "Token.Name.Namespace",
	highlighter.TokenTypeNameOther:            "Token.Name.Other",
	highlighter.TokenTypeNameTag:              "Token.Name.Tag",
	highlighter.TokenTypeNameVariable:         "Token.Name.Variable",
	highlighter.TokenTypeNameVariableClass:    "Token.Name.Variable.Class",
	highlighter.TokenTypeNameVariableGlobal:   "Token.Name.Variable.Global",
	highlighter.TokenTypeNameVariableInstance: "Token.Name.Variable.Instance",
	highlighter.TokenTypeLiteral:              "Token.Literal",
	highlighter.TokenTypeLiteralDate:          "Token.Literal.Date",
	highlighter.TokenTypeString:               "Token.String",
	highlighter.TokenTypeStringBacktick:       "Token.String.Backtick",
	highlighter.TokenTypeStringChar:           "Token.String.Char",
	highlighter.TokenTypeStringDoc:            "Token.String.Doc",
	highlighter.TokenTypeStringDouble:         "Token.String.Double",
	highlighter.TokenTypeStringEscape:         "Token.String.Escape",
	highlighter.TokenTypeStringHeredoc:        "Token.String.Heredoc",
	highlighter.TokenTypeStringInterpol:       "Token.String.Interpol",
	highlighter.TokenTypeStringOther:          "Token.String.Other",
	highlighter.TokenTypeStringRegex:          "Token.String.Regex",
	highlighter.TokenTypeStringSingle:         "Token.String.Single",
	highlighter.TokenTypeStringSymbol:         "Token.String.Symbol",
	highlighter.TokenTypeNumber:               "Token.Number",
	highlighter.TokenTypeNumberFloat:          "Token.Number.Float",
	highlighter.TokenTypeNumberHex:            "Token.Number.Hex",
	highlighter.TokenTypeNumberInteger:        "Token.Number.Integer",
	highlighter.TokenTypeNumberIntegerLong:    "Token.Number.Integer.Long",
	highlighter.TokenTypeNumberOct:            "Token.Number.Oct",
	highlighter.TokenTypeOperator:             "Token.Operator",
	highlighter.TokenTypeOperatorWord:         "Token.Operator.Word",
	highlighter.TokenTypePunctuation:          "Token.Punctuation",
	highlighter.TokenTypeComment:              "Token.Comment",
	highlighter.TokenTypeCommentMultiline:     "Token.Comment.Multiline",
	highlighter.TokenTypeCommentPreproc:       "Token.Comment.Preproc",
	highlighter.TokenTypeCommentSingle:        "Token.Comment.Single",
	highlighter.TokenTypeCommentSpecial:       "Token.Comment.Special",
	highlighter.TokenTypeGeneric:              "Token.Generic",
	highlighter.TokenTypeGenericDeleted:       "Token.Generic.Deleted",
	highlighter.TokenTypeGenericEmph:          "Token.Generic.Emph",
	highlighter.TokenTypeGenericError:         "Token.Generic.Error",
	highlighter.TokenTypeGenericHeading:       "Token.Generic.Heading",
	highlighter.TokenTypeGenericInserted:      "Token.Generic.Inserted",
	highlighter.TokenTypeGenericOutput:        "Token.Generic.Output",
	highlighter.TokenTypeGenericPrompt:        "Token.Generic.Prompt",
	highlighter.TokenTypeGenericStrong:        "Token.Generic.Strong",
	highlighter.TokenTypeGenericSubheading:    "Token.Generic.Subheading",
	highlighter.TokenTypeGenericTraceback:     "Token.Generic.Traceback",
}

type RawTokenFormat struct {
}

func NewRawTokenFormat() RawTokenFormat {
	return RawTokenFormat{}
}

func (f RawTokenFormat) FormatTokens(w io.Writer, tokens highlighter.Tokens) {
	for _, token := range tokens {
		f.Format(w, token)
		w.Write([]byte("\n"))
	}
}

func (f RawTokenFormat) Format(w io.Writer, token highlighter.Token) {
	fmt.Fprintf(w, "%s: %#v", tokenTypeNameMap[token.TokenType], token.Text)
}
