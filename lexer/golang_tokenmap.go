package lexer

import (
	golang "github.com/shirou/antlr-grammars-v4-go/golang"

	"github.com/shirou/highlighter"
)

func NewGolangTokenMap() TokenMap {
	return TokenMap{
		symbolicMap: newGolangSymbolicMap(),
		ruleMap:     newGolangRuleMap(),
	}
}

func newGolangRuleMap() RuleMap {
	return RuleMap{}
}
func newGolangSymbolicMap() TypeMap {
	return TypeMap{
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
