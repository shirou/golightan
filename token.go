package golightan

import (
	"fmt"
)

type TokenType int

type Token struct {
	TokenType TokenType
	Text      string
	Position  Position
}

type Tokens []Token

func (t Token) String() string {
	return fmt.Sprintf("%d: %s", t.TokenType, t.Text)
}

const (
	TokenTypeText TokenType = iota
	TokenTypeWhitespace
	TokenTypeError
	TokenTypeOther
	TokenTypeKeyword
	TokenTypeKeywordConstant
	TokenTypeKeywordDeclaration
	TokenTypeKeywordNamespace
	TokenTypeKeywordPseudo
	TokenTypeKeywordReserved
	TokenTypeKeywordType
	TokenTypeName
	TokenTypeNameAttribute
	TokenTypeNameBuiltin
	TokenTypeNameBuiltinPseudo
	TokenTypeNameClass
	TokenTypeNameConstant
	TokenTypeNameDecorator
	TokenTypeNameEntity
	TokenTypeNameException
	TokenTypeNameFunction
	TokenTypeNameProperty
	TokenTypeNameLabel
	TokenTypeNameNamespace
	TokenTypeNameOther
	TokenTypeNameTag
	TokenTypeNameVariable
	TokenTypeNameVariableClass
	TokenTypeNameVariableGlobal
	TokenTypeNameVariableInstance
	TokenTypeLiteral
	TokenTypeLiteralDate
	TokenTypeString
	TokenTypeStringBacktick
	TokenTypeStringChar
	TokenTypeStringDoc
	TokenTypeStringDouble
	TokenTypeStringEscape
	TokenTypeStringHeredoc
	TokenTypeStringInterpol
	TokenTypeStringOther
	TokenTypeStringRegex
	TokenTypeStringSingle
	TokenTypeStringSymbol
	TokenTypeNumber
	TokenTypeNumberFloat
	TokenTypeNumberHex
	TokenTypeNumberInteger
	TokenTypeNumberIntegerLong
	TokenTypeNumberOct
	TokenTypeOperator
	TokenTypeOperatorWord
	TokenTypePunctuation
	TokenTypeComment
	TokenTypeCommentMultiline
	TokenTypeCommentPreproc
	TokenTypeCommentSingle
	TokenTypeCommentSpecial
	TokenTypeGeneric
	TokenTypeGenericDeleted
	TokenTypeGenericEmph
	TokenTypeGenericError
	TokenTypeGenericHeading
	TokenTypeGenericInserted
	TokenTypeGenericOutput
	TokenTypeGenericPrompt
	TokenTypeGenericStrong
	TokenTypeGenericSubheading
	TokenTypeGenericTraceback
)

var CSSMap = map[TokenType]string{
	TokenTypeText:                 "",
	TokenTypeWhitespace:           "w",
	TokenTypeError:                "err",
	TokenTypeOther:                "x",
	TokenTypeKeyword:              "k",
	TokenTypeKeywordConstant:      "kc",
	TokenTypeKeywordDeclaration:   "kd",
	TokenTypeKeywordNamespace:     "kn",
	TokenTypeKeywordPseudo:        "kp",
	TokenTypeKeywordReserved:      "kr",
	TokenTypeKeywordType:          "kt",
	TokenTypeName:                 "n",
	TokenTypeNameAttribute:        "na",
	TokenTypeNameBuiltin:          "nb",
	TokenTypeNameBuiltinPseudo:    "bp",
	TokenTypeNameClass:            "nc",
	TokenTypeNameConstant:         "no",
	TokenTypeNameDecorator:        "nd",
	TokenTypeNameEntity:           "ni",
	TokenTypeNameException:        "ne",
	TokenTypeNameFunction:         "nf",
	TokenTypeNameProperty:         "py",
	TokenTypeNameLabel:            "nl",
	TokenTypeNameNamespace:        "nn",
	TokenTypeNameOther:            "nx",
	TokenTypeNameTag:              "nt",
	TokenTypeNameVariable:         "nv",
	TokenTypeNameVariableClass:    "vc",
	TokenTypeNameVariableGlobal:   "vg",
	TokenTypeNameVariableInstance: "vi",
	TokenTypeLiteral:              "l",
	TokenTypeLiteralDate:          "ld",
	TokenTypeString:               "s",
	TokenTypeStringBacktick:       "sb",
	TokenTypeStringChar:           "sc",
	TokenTypeStringDoc:            "sd",
	TokenTypeStringDouble:         "s2",
	TokenTypeStringEscape:         "se",
	TokenTypeStringHeredoc:        "sh",
	TokenTypeStringInterpol:       "si",
	TokenTypeStringOther:          "sx",
	TokenTypeStringRegex:          "sr",
	TokenTypeStringSingle:         "s1",
	TokenTypeStringSymbol:         "ss",
	TokenTypeNumber:               "m",
	TokenTypeNumberFloat:          "mf",
	TokenTypeNumberHex:            "mh",
	TokenTypeNumberInteger:        "mi",
	TokenTypeNumberIntegerLong:    "il",
	TokenTypeNumberOct:            "mo",
	TokenTypeOperator:             "o",
	TokenTypeOperatorWord:         "ow",
	TokenTypePunctuation:          "p",
	TokenTypeComment:              "c",
	TokenTypeCommentMultiline:     "cm",
	TokenTypeCommentPreproc:       "cp",
	TokenTypeCommentSingle:        "c1",
	TokenTypeCommentSpecial:       "cs",
	TokenTypeGeneric:              "g",
	TokenTypeGenericDeleted:       "gd",
	TokenTypeGenericEmph:          "ge",
	TokenTypeGenericError:         "gr",
	TokenTypeGenericHeading:       "gh",
	TokenTypeGenericInserted:      "gi",
	TokenTypeGenericOutput:        "go",
	TokenTypeGenericPrompt:        "gp",
	TokenTypeGenericStrong:        "gs",
	TokenTypeGenericSubheading:    "gu",
	TokenTypeGenericTraceback:     "gt",
}
