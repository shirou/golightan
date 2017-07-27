package golightan

import (
	"bytes"
)

type ActionType string

const (
	ActionTypePush ActionType = "#push"
	ActionTypePop             = "#pop"
	ActionTypePop2            = "#pop:2"
	ActionTypePop3            = "#pop:3"
	ActionTypePop4            = "#pop:4"
	ActionTypePop5            = "#pop:5"
	ActionTypeMore            = "#more"
)

type Mode struct {
	Matches Matches
}

func (m Mode) Match(ch rune, buf *bytes.Buffer) *Match {
	chstr := string(ch)

	for _, match := range m.Matches {
		if !match.Match(chstr) {
			continue
		}
		if !match.Match(buf.String()) {
			continue
		}
		return &match
	}
	return nil
}

type Match struct {
	S []string  // String
	T TokenType // TokenType
	A string    // Action
}
type Matches []Match

func (m Match) Match(ch string) bool {
	for _, t := range m.S {
		if t == ch {
			return true
		}
	}
	return false
}

func MakeMatches(s []string, tt TokenType) Matches {
	ret := make(Matches, len(s))
	for i, m := range s {
		ret[i] = Match{
			S: []string{m},
			T: tt,
		}
	}
	return ret
}

func NewMatch(s ...Matches) Matches {
	ret := make(Matches, 0, len(s))
	for _, r := range s {
		ret = append(ret, r...)
	}

	return ret
}

type ModeMap map[string]Mode
