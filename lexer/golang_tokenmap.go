package lexer

import (
	golang "github.com/shirou/antlr-grammars-v4-go/golang"

	"github.com/shirou/golightan"
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
		golang.GolangParserIDENTIFIER: golightan.TokenTypeKeyword,
		golang.GolangParserKEYWORD:    golightan.TokenTypeKeyword,
		//		golang.GolangParserBINARY_OP:              golightan.TokenTypeKeyword,
		//		golang.GolangParserREL_OP:                 golightan.TokenTypeKeyword,
		//		golang.GolangParserADD_OP:                 golightan.TokenTypeKeyword,
		//		golang.GolangParserMUL_OP:                 golightan.TokenTypeKeyword,
		//		golang.GolangParserUNARY_OP:               golightan.TokenTypeKeyword,
		//		golang.GolangParserINT_LIT:                golightan.TokenTypeKeyword,
		//		golang.GolangParserDECIMAL_LIT:            golightan.TokenTypeKeyword,
		//		golang.GolangParserOCTAL_LIT:              golightan.TokenTypeKeyword,
		//		golang.GolangParserHEX_LIT:                golightan.TokenTypeKeyword,
		//		golang.GolangParserFLOAT_LIT:              golightan.TokenTypeKeyword,
		//		golang.GolangParserDECIMALS:               golightan.TokenTypeKeyword,
		//		golang.GolangParserEXPONENT:               golightan.TokenTypeKeyword,
		//		golang.GolangParserIMAGINARY_LIT:          golightan.TokenTypeKeyword,
		//		golang.GolangParserRUNE_LIT:               golightan.TokenTypeKeyword,
		//		golang.GolangParserUNICODE_VALUE:          golightan.TokenTypeKeyword,
		//		golang.GolangParserBYTE_VALUE:             golightan.TokenTypeKeyword,
		//		golang.GolangParserOCTAL_BYTE_VALUE:       golightan.TokenTypeKeyword,
		//		golang.GolangParserHEX_BYTE_VALUE:         golightan.TokenTypeKeyword,
		//		golang.GolangParserLITTLE_U_VALUE:         golightan.TokenTypeKeyword,
		//		golang.GolangParserBIG_U_VALUE:            golightan.TokenTypeKeyword,
		//		golang.GolangParserESCAPED_CHAR:           golightan.TokenTypeKeyword,
		//		golang.GolangParserSTRING_LIT:             golightan.TokenTypeKeyword,
		//		golang.GolangParserRAW_STRING_LIT:         golightan.TokenTypeKeyword,
		//		golang.GolangParserINTERPRETED_STRING_LIT: golightan.TokenTypeKeyword,
		//		golang.GolangParserLETTER:                 golightan.TokenTypeKeyword,
		//		golang.GolangParserDECIMAL_DIGIT:          golightan.TokenTypeKeyword,
		//		golang.GolangParserOCTAL_DIGIT:            golightan.TokenTypeKeyword,
		//		golang.GolangParserHEX_DIGIT:              golightan.TokenTypeKeyword,
		//		golang.GolangParserNEWLINE:                golightan.TokenTypeKeyword,
		//		golang.GolangParserUNICODE_CHAR:           golightan.TokenTypeKeyword,
		//		golang.GolangParserUNICODE_DIGIT:          golightan.TokenTypeKeyword,
		//		golang.GolangParserUNICODE_LETTER:         golightan.TokenTypeKeyword,
		golang.GolangParserWS:          golightan.TokenTypeKeyword,
		golang.GolangParserCOMMENT:     golightan.TokenTypeKeyword,
		golang.GolangParserTERMINATOR:  golightan.TokenTypeKeyword,
		golang.GolangLexerLINE_COMMENT: golightan.TokenTypeCommentSingle,
	}
}
