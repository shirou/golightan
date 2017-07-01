package lexer

import (
	"bufio"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/shirou/golightan"
	"github.com/shirou/golightan/formatter"
)

type TestCase struct {
	src, exp string
}

func loadPygments(filename string) golightan.Tokens {
	fp, err := os.Open(filepath.Join("testcase", filename))
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	tokens := make(golightan.Tokens, 0)

	reader := bufio.NewReaderSize(fp, 4096)
	for {
		line, _, err := reader.ReadLine()

		fields := strings.SplitN(string(line), "\t", -1)
		if len(fields) != 2 {
			break
		}
		text := strings.Trim(fields[1], `u'`)
		text = strings.Replace(text, `\n`, "\n", 1)

		if strings.TrimSpace(text) == "" {
			continue
		}

		tokens = append(tokens, golightan.Token{
			TokenType: convert(fields[0]),
			Text:      text,
		})
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}
	return tokens
}

// removeWHNode remove WhiteSpace or CR node to just a testing
func removeWHNode(tokens golightan.Tokens) golightan.Tokens {
	ret := make(golightan.Tokens, 0)
	for _, token := range tokens {
		text := strings.Replace(token.Text, "\t", "", -1)
		if strings.TrimSpace(text) != "" {
			ret = append(ret, token)
		}
	}
	return ret

}

func rawDiff(t *testing.T, test TestCase, target string) {
	lexer, err := Factory(target)
	if err != nil {
		t.Fatal(err)
	}
	input, err := antlr.NewFileStream(filepath.Join("testcase", test.src))
	if err != nil {
		t.Fatal(err)
	}

	tokens, err := lexer.Tokenize(input)
	if err != nil {
		t.Fatal(err)
	}
	if len(tokens) == 0 {
		t.Fatal("tokenize failed")
	}

	f, _ := formatter.Factory("terminal", "")
	f.FormatTokens(os.Stdout, tokens)
	/*

		exps := loadPygments(test.exp)

		for i, token := range removeWHNode(tokens) {
				if len(exps) <= i {
					t.Errorf("length missmatch: %d", i)
					break
				}
				exp := exps[i]
					if token.Text != exp.Text {
						t.Errorf("text: %d:%s expected: %s -> actual: %s",
							i, token.Text, exp.Text, token.Text)
					}
					if token.TokenType != exp.TokenType {
						t.Errorf("type: %d:%s expected: %d(%s) -> actual: %d",
							i, token.Text, exp.TokenType, golightan.CSSMap[exp.TokenType], token.TokenType)
					}
		}
	*/

}

func runTests(t *testing.T, tests []TestCase, target string) {
	for _, test := range tests {
		t.Run(test.src, func(t *testing.T) {
			rawDiff(t, test, target)
		})
	}
}

func TestSQLite(t *testing.T) {
	tests := []TestCase{
		TestCase{"sqlite/input-1.sql", "sqlite/input-1.raw"},
		TestCase{"sqlite/input-2.sql", "sqlite/input-2.raw"},
	}
	runTests(t, tests, "sqlite")
}
func TestJSON(t *testing.T) {
	tests := []TestCase{
		TestCase{"json/example1.json", "json/example1.raw"},
	}
	runTests(t, tests, "json")
}

func TestGolang(t *testing.T) {
	tests := []TestCase{
		TestCase{"golang/example.go", "golang/example.raw"},
		TestCase{"golang/example2.go", "golang/example2.raw"},
	}
	runTests(t, tests, "golang")
}

func TestGraphQL(t *testing.T) {
	tests := []TestCase{
		TestCase{"graphql/example.go", "graphql/example.raw"},
	}
	runTests(t, tests, "graphql")
}
func TestXML(t *testing.T) {
	tests := []TestCase{
		TestCase{"xml/books.xml", "xml/books.raw"},
		TestCase{"xml/web.xml", "xml/web.raw"},
	}
	runTests(t, tests, "xml")
}

func TestPython3(t *testing.T) {
	tests := []TestCase{
		TestCase{"python3/__init__.py", "python3/__init__.raw"},
		TestCase{"python3/tasks.py", "python3/tasks.raw"},
	}
	runTests(t, tests, "python3")
}

func TestC(t *testing.T) {
	tests := []TestCase{
		TestCase{"c/BinaryDigit.c", "c/BinaryDigit.raw"},
		TestCase{"c/add.c", "c/add.raw"},
		TestCase{"c/bt.c", "c/bt.raw"},
	}
	runTests(t, tests, "c")
}

func convert(s string) golightan.TokenType {
	m := map[string]golightan.TokenType{
		"Token.Text":                   golightan.TokenTypeText,
		"Token.Whitespace":             golightan.TokenTypeWhitespace,
		"Token.Error":                  golightan.TokenTypeError,
		"Token.Other":                  golightan.TokenTypeOther,
		"Token.Keyword":                golightan.TokenTypeKeyword,
		"Token.Keyword.Constant":       golightan.TokenTypeKeywordConstant,
		"Token.Keyword.Declaration":    golightan.TokenTypeKeywordDeclaration,
		"Token.Keyword.Namespace":      golightan.TokenTypeKeywordNamespace,
		"Token.Keyword.Pseudo":         golightan.TokenTypeKeywordPseudo,
		"Token.Keyword.Reserved":       golightan.TokenTypeKeywordReserved,
		"Token.Keyword.Type":           golightan.TokenTypeKeywordType,
		"Token.Name":                   golightan.TokenTypeName,
		"Token.Name.Attribute":         golightan.TokenTypeNameAttribute,
		"Token.Name.Builtin":           golightan.TokenTypeNameBuiltin,
		"Token.Name.BuiltinPseudo":     golightan.TokenTypeNameBuiltinPseudo,
		"Token.Name.Class":             golightan.TokenTypeNameClass,
		"Token.Name.Constant":          golightan.TokenTypeNameConstant,
		"Token.Name.Decorator":         golightan.TokenTypeNameDecorator,
		"Token.Name.Entity":            golightan.TokenTypeNameEntity,
		"Token.Name.Exception":         golightan.TokenTypeNameException,
		"Token.Name.Function":          golightan.TokenTypeNameFunction,
		"Token.Name.Property":          golightan.TokenTypeNameProperty,
		"Token.Name.Label":             golightan.TokenTypeNameLabel,
		"Token.Name.Namespace":         golightan.TokenTypeNameNamespace,
		"Token.Name.Other":             golightan.TokenTypeNameOther,
		"Token.Name.Tag":               golightan.TokenTypeNameTag,
		"Token.Name.Variable":          golightan.TokenTypeNameVariable,
		"Token.Name.Variable.Class":    golightan.TokenTypeNameVariableClass,
		"Token.Name.Variable.Global":   golightan.TokenTypeNameVariableGlobal,
		"Token.Name.Variable.Instance": golightan.TokenTypeNameVariableInstance,
		"Token.Literal":                golightan.TokenTypeLiteral,
		"Token.Literal.Date":           golightan.TokenTypeLiteralDate,
		"Token.String":                 golightan.TokenTypeString,
		"Token.String.Backtick":        golightan.TokenTypeStringBacktick,
		"Token.String.Char":            golightan.TokenTypeStringChar,
		"Token.String.Doc":             golightan.TokenTypeStringDoc,
		"Token.String.Double":          golightan.TokenTypeStringDouble,
		"Token.String.Escape":          golightan.TokenTypeStringEscape,
		"Token.String.Heredoc":         golightan.TokenTypeStringHeredoc,
		"Token.String.Interpol":        golightan.TokenTypeStringInterpol,
		"Token.String.Other":           golightan.TokenTypeStringOther,
		"Token.String.Regex":           golightan.TokenTypeStringRegex,
		"Token.String.Single":          golightan.TokenTypeStringSingle,
		"Token.String.Symbol":          golightan.TokenTypeStringSymbol,
		"Token.Number":                 golightan.TokenTypeNumber,
		"Token.Number.Float":           golightan.TokenTypeNumberFloat,
		"Token.Number.Hex":             golightan.TokenTypeNumberHex,
		"Token.Number.Integer":         golightan.TokenTypeNumberInteger,
		"Token.Number.IntegerLong":     golightan.TokenTypeNumberIntegerLong,
		"Token.Number.Oct":             golightan.TokenTypeNumberOct,
		"Token.Operator":               golightan.TokenTypeOperator,
		"Token.Operator.Word":          golightan.TokenTypeOperatorWord,
		"Token.Punctuation":            golightan.TokenTypePunctuation,
		"Token.Comment":                golightan.TokenTypeComment,
		"Token.Comment.Multiline":      golightan.TokenTypeCommentMultiline,
		"Token.Comment.Preproc":        golightan.TokenTypeCommentPreproc,
		"Token.Comment.Single":         golightan.TokenTypeCommentSingle,
		"Token.Comment.Special":        golightan.TokenTypeCommentSpecial,
		"Token.Generic":                golightan.TokenTypeGeneric,
		"Token.Generic.Deleted":        golightan.TokenTypeGenericDeleted,
		"Token.Generic.Emph":           golightan.TokenTypeGenericEmph,
		"Token.Generic.Error":          golightan.TokenTypeGenericError,
		"Token.Generic.Heading":        golightan.TokenTypeGenericHeading,
		"Token.Generic.Inserted":       golightan.TokenTypeGenericInserted,
		"Token.Generic.Output":         golightan.TokenTypeGenericOutput,
		"Token.Generic.Prompt":         golightan.TokenTypeGenericPrompt,
		"Token.Generic.Strong":         golightan.TokenTypeGenericStrong,
		"Token.Generic.Subheading":     golightan.TokenTypeGenericSubheading,
		"Token.Generic.Traceback":      golightan.TokenTypeGenericTraceback,
	}

	return m[s]
}
