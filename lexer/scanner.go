package lexer

import (
	"bufio"
	"bytes"
	"fmt"
	"io"

	"github.com/shirou/golightan"
	"github.com/shirou/golightan/lexer/languages"
)

const (
	initialStackSize = 20
	initialTokenSize = 100
)

type Scanner struct {
	r               *bufio.Reader
	buf             bytes.Buffer
	stack           *Stack
	modeMap         golightan.ModeMap
	caseInsensitive bool
	position        *golightan.Position
}

// NewScanner returns a new instance of Scanner.
func NewScanner(r io.Reader, lang languages.Languager) *Scanner {
	return &Scanner{
		r:               bufio.NewReader(r),
		modeMap:         lang.GetModeMap(),
		caseInsensitive: lang.GetCaseInsentive(),
		position:        golightan.NewPosition(""),
		stack:           NewStack(initialStackSize),
	}
}

func (s *Scanner) Tokenize(r io.Reader) (golightan.Tokens, error) {
	tokens := make(golightan.Tokens, 0, initialTokenSize)

	mode := s.modeMap["root"]
	var token golightan.Token
	var err error
	for {
		token, mode, err = s.Scan(mode)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		tokens = append(tokens, token)
	}

	return tokens, nil
}

// read reads the next rune from the bufferred reader.
// Returns the io.EOF if an error occurs (or io.EOF is returned).
func (s *Scanner) read() (rune, error) {
	ch, _, err := s.r.ReadRune()
	if err != nil {
		return 0, io.EOF
	}
	return ch, nil
}

func (s *Scanner) unread() error {
	return s.r.UnreadRune()
}

// Scan returns the next literal value.
func (s *Scanner) Scan(mode golightan.Mode) (golightan.Token, golightan.Mode, error) {
	buf := new(bytes.Buffer)
	retMode := mode

	for {
		ch, err := s.read()
		if err != nil {
			return golightan.Token{}, golightan.Mode{}, err
		}
		buf.WriteRune(ch)
		fmt.Println(mode)
		match := mode.Match(ch, buf)
		if match != nil {
			if match.A != "" {
				mode = s.modeMap[match.A]
			}
			return golightan.Token{
				TokenType: match.T,
				Text:      buf.String(),
				Position:  *s.position,
			}, mode, nil
		}

		if isWhiteSpace(ch) {
			if ch == '\n' {
				s.position.AddLine(1)
			}
			// TODO 空白続きかの判定

			return golightan.Token{
				TokenType: golightan.TokenTypeWhitespace,
				Text:      buf.String(),
				Position:  *s.position,
			}, mode, nil
		}

		//		fmt.Println(buf.String())

		/*
			for ch, err := s.read(); isWhiteSpace(ch); ch, err = s.read() {
				if err != nil {
					return golightan.Token{}, err
				}

				buf.WriteRune(ch)
				ws = true
				if ch == '\n' {
					s.position.AddLine(1)
				}
			}
			if ws {
				s.unread() // restore last non-whitespace rune
				ws := golightan.Token{
					TokenType: golightan.TokenTypeWhitespace,
					Text:      buf.String(),
					Position: golightan.Position{
						Line: s.position.Line,
					},
				}
				return ws, nil
			}

			switch ch {
			case ' ', '\t', '\n', '\r':
			case eof:
				return buf.String(), io.EOF
			default:
				buf.WriteRune(ch)
			}
		*/
	}

	return golightan.Token{}, retMode, nil
}

func isWhiteSpace(ch rune) bool {
	switch ch {
	case ' ', '\t', '\n', '\r':
		return true
	default:
		return false
	}
}
