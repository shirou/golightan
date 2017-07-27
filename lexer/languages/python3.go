package languages

import (
	g "github.com/shirou/golightan"
)

type Python3Language struct {
	modeMap         g.ModeMap
	caseInsensitive bool
	detect          g.Detect
}

func NewPython3Language() Languager {
	return Python3Language{
		modeMap: NewPython3ModeMap(),
	}
}

func (l Python3Language) GetModeMap() g.ModeMap  { return l.modeMap }
func (l Python3Language) GetCaseInsentive() bool { return l.caseInsensitive }
func (l Python3Language) GetDetect() g.Detect    { return l.detect }

func NewPython3ModeMap() g.ModeMap {
	keywords := g.MakeMatches([]string{
		"assert", "break", "continue", "del", "elif", "else", "except",
		"exec", "finally", "for", "global", "if", "lambda", "pass",
		"print", "raise", "return", "try", "while", "yield",
		"yield from", "as", "with"}, g.TokenTypeKeyword)
	builtins := g.MakeMatches([]string{
		"__import__", "abs", "all", "any", "apply", "basestring", "bin",
		"bool", "buffer", "bytearray", "bytes", "callable", "chr", "classmethod",
		"cmp", "coerce", "compile", "complex", "delattr", "dict", "dir", "divmod",
		"enumerate", "eval", "execfile", "exit", "file", "filter", "float",
		"frozenset", "getattr", "globals", "hasattr", "hash", "hex", "id",
		"input", "int", "intern", "isinstance", "issubclass", "iter", "len",
		"list", "locals", "long", "map", "max", "min", "next", "object",
		"oct", "open", "ord", "pow", "property", "range", "raw_input", "reduce",
		"reload", "repr", "reversed", "round", "set", "setattr", "slice",
		"sorted", "staticmethod", "str", "sum", "super", "tuple", "type",
		"unichr", "unicode", "vars", "xrange", "zip"}, g.TokenTypeNameBuiltin)

	m := g.ModeMap{
		"root": g.Mode{
			Matches: g.NewMatch(
				keywords,
				builtins,
				g.Matches{
					g.Match{S: []string{`//`}, T: g.TokenTypeComment},
					g.Match{S: []string{`"`}, A: "string"},
				},
			),
		},
		"string": g.Mode{
			Matches: g.NewMatch(
				g.Matches{
					g.Match{S: []string{`"`}, A: "pop", T: g.TokenTypeComment},
				},
			),
		},
	}

	return m
}
