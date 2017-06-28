package lexer

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	golang "github.com/shirou/antlr-grammars-v4-go/golang"

	"github.com/shirou/highlighter"
)

type GolangLexer struct {
	lexer       antlr.Lexer
	ruleMap     TypeMap
	literalMap  TypeMap
	symbolicMap TypeMap
}

func (l GolangLexer) Tokenize(input antlr.CharStream) (highlighter.Tokens, error) {
	le := golang.NewGolangLexer(input)
	return CommonTokenize(le, l.symbolicMap)
}

func NewGolangLexer() Lexer {
	symbolicMap := TypeMap{
		golang.GolangParserIDENTIFIER: highlighter.TokenTypeKeyword,
		golang.GolangParserKEYWORD:    highlighter.TokenTypeKeyword,
		//		golang.GolangParserBINARY_OP:              highlighter.TokenTypeKeyword,
		//		golang.GolangParserREL_OP:                 highlighter.TokenTypeKeyword,
		//		golang.GolangParserADD_OP:                 highlighter.TokenTypeKeyword,
		//		golang.GolangParserMUL_OP:                 highlighter.TokenTypeKeyword,
		//		golang.GolangParserUNARY_OP:               highlighter.TokenTypeKeyword,
		//		golang.GolangParserINT_LIT:                highlighter.TokenTypeKeyword,
		//		golang.GolangParserDECIMAL_LIT:            highlighter.TokenTypeKeyword,
		//		golang.GolangParserOCTAL_LIT:              highlighter.TokenTypeKeyword,
		//		golang.GolangParserHEX_LIT:                highlighter.TokenTypeKeyword,
		//		golang.GolangParserFLOAT_LIT:              highlighter.TokenTypeKeyword,
		//		golang.GolangParserDECIMALS:               highlighter.TokenTypeKeyword,
		//		golang.GolangParserEXPONENT:               highlighter.TokenTypeKeyword,
		//		golang.GolangParserIMAGINARY_LIT:          highlighter.TokenTypeKeyword,
		//		golang.GolangParserRUNE_LIT:               highlighter.TokenTypeKeyword,
		//		golang.GolangParserUNICODE_VALUE:          highlighter.TokenTypeKeyword,
		//		golang.GolangParserBYTE_VALUE:             highlighter.TokenTypeKeyword,
		//		golang.GolangParserOCTAL_BYTE_VALUE:       highlighter.TokenTypeKeyword,
		//		golang.GolangParserHEX_BYTE_VALUE:         highlighter.TokenTypeKeyword,
		//		golang.GolangParserLITTLE_U_VALUE:         highlighter.TokenTypeKeyword,
		//		golang.GolangParserBIG_U_VALUE:            highlighter.TokenTypeKeyword,
		//		golang.GolangParserESCAPED_CHAR:           highlighter.TokenTypeKeyword,
		//		golang.GolangParserSTRING_LIT:             highlighter.TokenTypeKeyword,
		//		golang.GolangParserRAW_STRING_LIT:         highlighter.TokenTypeKeyword,
		//		golang.GolangParserINTERPRETED_STRING_LIT: highlighter.TokenTypeKeyword,
		//		golang.GolangParserLETTER:                 highlighter.TokenTypeKeyword,
		//		golang.GolangParserDECIMAL_DIGIT:          highlighter.TokenTypeKeyword,
		//		golang.GolangParserOCTAL_DIGIT:            highlighter.TokenTypeKeyword,
		//		golang.GolangParserHEX_DIGIT:              highlighter.TokenTypeKeyword,
		//		golang.GolangParserNEWLINE:                highlighter.TokenTypeKeyword,
		//		golang.GolangParserUNICODE_CHAR:           highlighter.TokenTypeKeyword,
		//		golang.GolangParserUNICODE_DIGIT:          highlighter.TokenTypeKeyword,
		//		golang.GolangParserUNICODE_LETTER:         highlighter.TokenTypeKeyword,
		golang.GolangParserWS:          highlighter.TokenTypeKeyword,
		golang.GolangParserCOMMENT:     highlighter.TokenTypeKeyword,
		golang.GolangParserTERMINATOR:  highlighter.TokenTypeKeyword,
		golang.GolangLexerLINE_COMMENT: highlighter.TokenTypeCommentSingle,
	}

	return GolangLexer{
		symbolicMap: symbolicMap,
	}
}

type GolangParseTreeListener struct {
	tokens        highlighter.Tokens
	ruleNames     []string
	literalNames  []string
	symbolicNames []string
	inComment     bool
}

func (b GolangParseTreeListener) VisitTerminal(node antlr.TerminalNode) {
	token := node.GetSymbol()
	t := token.GetTokenType()
	if t < 0 {
		return
	}
	b.tokens = append(b.tokens, highlighter.Token{
		OriginalToken: token,
		//			TokenType:     tm.Get(t),
		Text: token.GetText(),
	})
}
func (b GolangParseTreeListener) VisitErrorNode(node antlr.ErrorNode) {
}
func (b GolangParseTreeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println(b.ruleNames[ctx.GetRuleIndex()])

	//	fmt.Println(ctx.GetStart())
}
func (b GolangParseTreeListener) ExitEveryRule(ctx antlr.ParserRuleContext) {
	//	fmt.Println(ctx.GetText())
	//	fmt.Println(ctx.GetText())
	//fmt.Println(ctx.GetStart())
}

func NewGolangParseTreeListener(p *golang.GolangParser) GolangParseTreeListener {
	return GolangParseTreeListener{
		ruleNames:     p.RuleNames,
		literalNames:  p.LiteralNames,
		symbolicNames: p.SymbolicNames,
		tokens:        make(highlighter.Tokens, 0),
	}
}
