package lexer

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	golang "github.com/shirou/antlr-grammars-v4-go/golang"

	"github.com/shirou/highlighter"
)

type GolangTypeMap TypeMap

var golangTypeMap GolangTypeMap

func init() {
	golangTypeMap = NewGolangTypeMap()
}

func NewGolangTypeMap() GolangTypeMap {
	return GolangTypeMap{
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
}

func (tm GolangTypeMap) Get(type_ int) highlighter.TokenType {
	s, ok := tm[type_]
	if !ok {
		return 0
	}
	return s
}

type BaseParseTreeListner struct {
	tokens highlighter.Tokens
}

func (b BaseParseTreeListner) VisitTerminal(node antlr.TerminalNode) {
	fmt.Println(node.GetText())
}
func (b BaseParseTreeListner) VisitErrorNode(node antlr.ErrorNode) {
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

	//	fmt.Println(node.GetText())
}
func (b BaseParseTreeListner) EnterEveryRule(ctx antlr.ParserRuleContext) {
	//	fmt.Println(ctx.GetStart())
}
func (b BaseParseTreeListner) ExitEveryRule(ctx antlr.ParserRuleContext) {
	//	fmt.Println(ctx.GetText())
	//	fmt.Println(ctx.GetText())
	//fmt.Println(ctx.GetStart())

}

func NewParseTreeListener() BaseParseTreeListner {
	return BaseParseTreeListner{
		tokens: make(highlighter.Tokens, 0),
	}
}

func NewGolangLexer(input antlr.CharStream) antlr.Lexer {
	return golang.NewGolangLexer(input)
}
