/**
 * Notice: This visitor file is an testing sample. THIS DOES NOT WORK!!!
 */

package lexer

import (
	"fmt"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/shirou/golightan"
)

type CommonParseTreeVisitor struct {
	tokens   golightan.Tokens
	tokenMap TokenMap
	rule     int
	stack    *Stack
	lexer    antlr.Lexer
}

func (b *CommonParseTreeVisitor) Token(node antlr.TerminalNode) {
	token := node.GetSymbol()

	t := token.GetTokenType()
	if t < 0 {
		return
	}

	text := node.GetText()
	new := golightan.Token{
		OriginalToken: token,
		TokenType:     b.tokenMap.Convert(b.stack.Last(), t, text),
		Text:          text,
	}

	if b.lexer != nil {
		b.debug(node)
	}

	b.tokens = append(b.tokens, new)
}

func (b *CommonParseTreeVisitor) GetTokens() golightan.Tokens { return b.tokens }

func (b *CommonParseTreeVisitor) VisitTerminal(node antlr.TerminalNode) interface{} {
	fmt.Println(node)
	b.Token(node)
	return node.Accept(b)
}
func (b *CommonParseTreeVisitor) VisitErrorNode(node antlr.ErrorNode) interface{} {
	fmt.Println("ERR", node)
	b.Token(node)
	return node.Accept(b)
}

func (b *CommonParseTreeVisitor) VisitChildren(node antlr.RuleNode) interface{} {
	//b.stack.Push(ctx.GetRuleIndex())
	fmt.Println("Children", node.GetText())

	return b.VisitChildren(node)
}
func (b *CommonParseTreeVisitor) Visit(tree antlr.ParseTree) interface{} {
	//	b.stack.Pop()
	fmt.Println("Visit", tree)
	return tree.Accept(b)
}

// SetDebug enable debug print which shows rule stack and symbol.
func (b *CommonParseTreeVisitor) SetDebug(lexer antlr.Lexer) {
	b.lexer = lexer
}

func (b *CommonParseTreeVisitor) debug(node antlr.TerminalNode) {
	token := node.GetSymbol()
	t := token.GetTokenType()
	s := make([]string, b.stack.Len())
	names := b.lexer.GetRuleNames()
	symbol := b.lexer.GetSymbolicNames()
	for i, r := range b.stack.stack {
		s[i] = names[r]
	}
	fmt.Printf("(%s, %s)-> %v\n", strings.Join(s, ","), symbol[t], node.GetText())
}

func NewCommonParseTreeVisitor(tm TokenMap) *CommonParseTreeVisitor {
	return &CommonParseTreeVisitor{
		tokens:   make(golightan.Tokens, 0, InitialTokenCapacity),
		tokenMap: tm,
		stack:    NewStack(InitialStackCapacity),
	}
}
