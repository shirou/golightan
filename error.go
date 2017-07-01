package golightan

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type NullErrorStrategy struct {
	*antlr.DefaultErrorStrategy
	errorRecoveryMode bool
}

func NewNullErrorStrategy() *NullErrorStrategy {
	b := new(NullErrorStrategy)
	b.DefaultErrorStrategy = antlr.NewDefaultErrorStrategy()
	return b
}

func (e NullErrorStrategy) RecoverInline(parser antlr.Parser) antlr.Token {
	return parser.GetCurrentToken()
}
