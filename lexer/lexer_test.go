package lexer

import (
	"bufio"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/shirou/highlighter"
)

type TestCase struct {
	src, exp string
}

func loadPygments(filename string) highlighter.Tokens {
	fp, err := os.Open(filepath.Join("testcase", filename))
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	tokens := make(highlighter.Tokens, 0)

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

		tokens = append(tokens, highlighter.Token{
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
func removeWHNode(tokens highlighter.Tokens) highlighter.Tokens {
	ret := make(highlighter.Tokens, 0)
	for _, token := range tokens {
		text := strings.Replace(token.Text, "\t", "", -1)
		if strings.TrimSpace(text) != "" {
			ret = append(ret, token)
		}
	}
	return ret

}

func rawDiff(t *testing.T, test TestCase, target string) {
	lexer, err := LexerFactory(target)
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
				i, token.Text, exp.TokenType, highlighter.CSSMap[exp.TokenType], token.TokenType)
		}
	}
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
		TestCase{"golang/ill_but_correct.go", "golang/ill_but_correct.raw"},
	}
	for _, test := range tests {
		t.Run(test.src, func(t *testing.T) {
			rawDiff(t, test, "golang")
		})
	}
}

func TestPython3(t *testing.T) {
	tests := []TestCase{
		TestCase{"python3/__init__.py", "python3/__init__.raw"},
		TestCase{"python3/tasks.py", "python3/tasks.raw"},
	}
	for _, test := range tests {
		t.Run(test.src, func(t *testing.T) {
			rawDiff(t, test, "python3")
		})
	}

}

func convert(s string) highlighter.TokenType {
	m := map[string]highlighter.TokenType{
		"Token.Text":                   highlighter.TokenTypeText,
		"Token.Whitespace":             highlighter.TokenTypeWhitespace,
		"Token.Error":                  highlighter.TokenTypeError,
		"Token.Other":                  highlighter.TokenTypeOther,
		"Token.Keyword":                highlighter.TokenTypeKeyword,
		"Token.Keyword.Constant":       highlighter.TokenTypeKeywordConstant,
		"Token.Keyword.Declaration":    highlighter.TokenTypeKeywordDeclaration,
		"Token.Keyword.Namespace":      highlighter.TokenTypeKeywordNamespace,
		"Token.Keyword.Pseudo":         highlighter.TokenTypeKeywordPseudo,
		"Token.Keyword.Reserved":       highlighter.TokenTypeKeywordReserved,
		"Token.Keyword.Type":           highlighter.TokenTypeKeywordType,
		"Token.Name":                   highlighter.TokenTypeName,
		"Token.Name.Attribute":         highlighter.TokenTypeNameAttribute,
		"Token.Name.Builtin":           highlighter.TokenTypeNameBuiltin,
		"Token.Name.BuiltinPseudo":     highlighter.TokenTypeNameBuiltinPseudo,
		"Token.Name.Class":             highlighter.TokenTypeNameClass,
		"Token.Name.Constant":          highlighter.TokenTypeNameConstant,
		"Token.Name.Decorator":         highlighter.TokenTypeNameDecorator,
		"Token.Name.Entity":            highlighter.TokenTypeNameEntity,
		"Token.Name.Exception":         highlighter.TokenTypeNameException,
		"Token.Name.Function":          highlighter.TokenTypeNameFunction,
		"Token.Name.Property":          highlighter.TokenTypeNameProperty,
		"Token.Name.Label":             highlighter.TokenTypeNameLabel,
		"Token.Name.Namespace":         highlighter.TokenTypeNameNamespace,
		"Token.Name.Other":             highlighter.TokenTypeNameOther,
		"Token.Name.Tag":               highlighter.TokenTypeNameTag,
		"Token.Name.Variable":          highlighter.TokenTypeNameVariable,
		"Token.Name.Variable.Class":    highlighter.TokenTypeNameVariableClass,
		"Token.Name.Variable.Global":   highlighter.TokenTypeNameVariableGlobal,
		"Token.Name.Variable.Instance": highlighter.TokenTypeNameVariableInstance,
		"Token.Literal":                highlighter.TokenTypeLiteral,
		"Token.Literal.Date":           highlighter.TokenTypeLiteralDate,
		"Token.String":                 highlighter.TokenTypeString,
		"Token.String.Backtick":        highlighter.TokenTypeStringBacktick,
		"Token.String.Char":            highlighter.TokenTypeStringChar,
		"Token.String.Doc":             highlighter.TokenTypeStringDoc,
		"Token.String.Double":          highlighter.TokenTypeStringDouble,
		"Token.String.Escape":          highlighter.TokenTypeStringEscape,
		"Token.String.Heredoc":         highlighter.TokenTypeStringHeredoc,
		"Token.String.Interpol":        highlighter.TokenTypeStringInterpol,
		"Token.String.Other":           highlighter.TokenTypeStringOther,
		"Token.String.Regex":           highlighter.TokenTypeStringRegex,
		"Token.String.Single":          highlighter.TokenTypeStringSingle,
		"Token.String.Symbol":          highlighter.TokenTypeStringSymbol,
		"Token.Number":                 highlighter.TokenTypeNumber,
		"Token.Number.Float":           highlighter.TokenTypeNumberFloat,
		"Token.Number.Hex":             highlighter.TokenTypeNumberHex,
		"Token.Number.Integer":         highlighter.TokenTypeNumberInteger,
		"Token.Number.IntegerLong":     highlighter.TokenTypeNumberIntegerLong,
		"Token.Number.Oct":             highlighter.TokenTypeNumberOct,
		"Token.Operator":               highlighter.TokenTypeOperator,
		"Token.Operator.Word":          highlighter.TokenTypeOperatorWord,
		"Token.Punctuation":            highlighter.TokenTypePunctuation,
		"Token.Comment":                highlighter.TokenTypeComment,
		"Token.Comment.Multiline":      highlighter.TokenTypeCommentMultiline,
		"Token.Comment.Preproc":        highlighter.TokenTypeCommentPreproc,
		"Token.Comment.Single":         highlighter.TokenTypeCommentSingle,
		"Token.Comment.Special":        highlighter.TokenTypeCommentSpecial,
		"Token.Generic":                highlighter.TokenTypeGeneric,
		"Token.Generic.Deleted":        highlighter.TokenTypeGenericDeleted,
		"Token.Generic.Emph":           highlighter.TokenTypeGenericEmph,
		"Token.Generic.Error":          highlighter.TokenTypeGenericError,
		"Token.Generic.Heading":        highlighter.TokenTypeGenericHeading,
		"Token.Generic.Inserted":       highlighter.TokenTypeGenericInserted,
		"Token.Generic.Output":         highlighter.TokenTypeGenericOutput,
		"Token.Generic.Prompt":         highlighter.TokenTypeGenericPrompt,
		"Token.Generic.Strong":         highlighter.TokenTypeGenericStrong,
		"Token.Generic.Subheading":     highlighter.TokenTypeGenericSubheading,
		"Token.Generic.Traceback":      highlighter.TokenTypeGenericTraceback,
	}

	return m[s]
}
