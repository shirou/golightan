package lexer

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"

	python3 "github.com/shirou/antlr-grammars-v4-go/python3"

	"github.com/shirou/highlighter"
)

type Python3Lexer struct {
	lexer       antlr.Lexer
	ruleMap     TypeMap
	literalMap  TypeMap
	symbolicMap TypeMap
}

func (l Python3Lexer) Tokenize(input antlr.CharStream) (highlighter.Tokens, error) {
	le := python3.NewPython3Lexer(input)
	stream := antlr.NewCommonTokenStream(le, 0)

	// Get All tokens
	num := 0
	for ; stream.Sync(num); num++ {
	}

	tokens := make(highlighter.Tokens, num)
	for i, token := range stream.GetAllTokens() {
		t := token.GetTokenType()
		if t < 0 {
			break
		}
		tokens[i] = highlighter.Token{
			OriginalToken: token,
			TokenType:     l.symbolicMap.Get(t),
			Text:          token.GetText(),
		}
	}

	return tokens, nil
}

type Python3ParseTreeListner struct {
	tokens        highlighter.Tokens
	ruleNames     []string
	literalNames  []string
	symbolicNames []string
	inComment     bool
}

func NewPython3Lexer() Lexer {
	symbolicMap := TypeMap{
		python3.Python3LexerSTRING:             highlighter.TokenTypeText,
		python3.Python3LexerNUMBER:             highlighter.TokenTypeKeyword,
		python3.Python3LexerINTEGER:            highlighter.TokenTypeNumberInteger,
		python3.Python3LexerDEF:                highlighter.TokenTypeKeyword,
		python3.Python3LexerRETURN:             highlighter.TokenTypeKeyword,
		python3.Python3LexerRAISE:              highlighter.TokenTypeKeyword,
		python3.Python3LexerFROM:               highlighter.TokenTypeKeyword,
		python3.Python3LexerIMPORT:             highlighter.TokenTypeKeywordNamespace,
		python3.Python3LexerAS:                 highlighter.TokenTypeKeyword,
		python3.Python3LexerGLOBAL:             highlighter.TokenTypeKeyword,
		python3.Python3LexerNONLOCAL:           highlighter.TokenTypeKeyword,
		python3.Python3LexerASSERT:             highlighter.TokenTypeKeyword,
		python3.Python3LexerIF:                 highlighter.TokenTypeKeyword,
		python3.Python3LexerELIF:               highlighter.TokenTypeKeyword,
		python3.Python3LexerELSE:               highlighter.TokenTypeKeyword,
		python3.Python3LexerWHILE:              highlighter.TokenTypeKeyword,
		python3.Python3LexerFOR:                highlighter.TokenTypeKeyword,
		python3.Python3LexerIN:                 highlighter.TokenTypeKeyword,
		python3.Python3LexerTRY:                highlighter.TokenTypeKeyword,
		python3.Python3LexerFINALLY:            highlighter.TokenTypeKeyword,
		python3.Python3LexerWITH:               highlighter.TokenTypeKeyword,
		python3.Python3LexerEXCEPT:             highlighter.TokenTypeKeyword,
		python3.Python3LexerLAMBDA:             highlighter.TokenTypeKeyword,
		python3.Python3LexerOR:                 highlighter.TokenTypeKeyword,
		python3.Python3LexerAND:                highlighter.TokenTypeKeyword,
		python3.Python3LexerNOT:                highlighter.TokenTypeKeyword,
		python3.Python3LexerIS:                 highlighter.TokenTypeKeyword,
		python3.Python3LexerNONE:               highlighter.TokenTypeKeyword,
		python3.Python3LexerTRUE:               highlighter.TokenTypeKeyword,
		python3.Python3LexerFALSE:              highlighter.TokenTypeKeyword,
		python3.Python3LexerCLASS:              highlighter.TokenTypeKeyword,
		python3.Python3LexerYIELD:              highlighter.TokenTypeKeyword,
		python3.Python3LexerDEL:                highlighter.TokenTypeKeyword,
		python3.Python3LexerPASS:               highlighter.TokenTypeKeyword,
		python3.Python3LexerCONTINUE:           highlighter.TokenTypeKeyword,
		python3.Python3LexerBREAK:              highlighter.TokenTypeKeyword,
		python3.Python3LexerASYNC:              highlighter.TokenTypeKeyword,
		python3.Python3LexerAWAIT:              highlighter.TokenTypeKeyword,
		python3.Python3LexerNEWLINE:            highlighter.TokenTypeText,
		python3.Python3LexerNAME:               highlighter.TokenTypeText,
		python3.Python3LexerSTRING_LITERAL:     highlighter.TokenTypeText,
		python3.Python3LexerBYTES_LITERAL:      highlighter.TokenTypeText,
		python3.Python3LexerDECIMAL_INTEGER:    highlighter.TokenTypeNumberInteger,
		python3.Python3LexerOCT_INTEGER:        highlighter.TokenTypeKeyword,
		python3.Python3LexerHEX_INTEGER:        highlighter.TokenTypeKeyword,
		python3.Python3LexerBIN_INTEGER:        highlighter.TokenTypeKeyword,
		python3.Python3LexerFLOAT_NUMBER:       highlighter.TokenTypeKeyword,
		python3.Python3LexerIMAG_NUMBER:        highlighter.TokenTypeKeyword,
		python3.Python3LexerDOT:                highlighter.TokenTypeKeyword,
		python3.Python3LexerELLIPSIS:           highlighter.TokenTypeKeyword,
		python3.Python3LexerSTAR:               highlighter.TokenTypeOperator,
		python3.Python3LexerOPEN_PAREN:         highlighter.TokenTypeText,
		python3.Python3LexerCLOSE_PAREN:        highlighter.TokenTypeText,
		python3.Python3LexerCOMMA:              highlighter.TokenTypeText,
		python3.Python3LexerCOLON:              highlighter.TokenTypeText,
		python3.Python3LexerSEMI_COLON:         highlighter.TokenTypeText,
		python3.Python3LexerPOWER:              highlighter.TokenTypeText,
		python3.Python3LexerASSIGN:             highlighter.TokenTypeText,
		python3.Python3LexerOPEN_BRACK:         highlighter.TokenTypeText,
		python3.Python3LexerCLOSE_BRACK:        highlighter.TokenTypeText,
		python3.Python3LexerOR_OP:              highlighter.TokenTypeText,
		python3.Python3LexerXOR:                highlighter.TokenTypeText,
		python3.Python3LexerAND_OP:             highlighter.TokenTypeText,
		python3.Python3LexerLEFT_SHIFT:         highlighter.TokenTypeText,
		python3.Python3LexerRIGHT_SHIFT:        highlighter.TokenTypeText,
		python3.Python3LexerADD:                highlighter.TokenTypeOperator,
		python3.Python3LexerMINUS:              highlighter.TokenTypeOperator,
		python3.Python3LexerDIV:                highlighter.TokenTypeOperator,
		python3.Python3LexerMOD:                highlighter.TokenTypeOperator,
		python3.Python3LexerIDIV:               highlighter.TokenTypeText,
		python3.Python3LexerNOT_OP:             highlighter.TokenTypeText,
		python3.Python3LexerOPEN_BRACE:         highlighter.TokenTypeText,
		python3.Python3LexerCLOSE_BRACE:        highlighter.TokenTypeText,
		python3.Python3LexerLESS_THAN:          highlighter.TokenTypeOperator,
		python3.Python3LexerGREATER_THAN:       highlighter.TokenTypeOperator,
		python3.Python3LexerEQUALS:             highlighter.TokenTypeOperator,
		python3.Python3LexerGT_EQ:              highlighter.TokenTypeOperator,
		python3.Python3LexerLT_EQ:              highlighter.TokenTypeOperator,
		python3.Python3LexerNOT_EQ_1:           highlighter.TokenTypeOperator,
		python3.Python3LexerNOT_EQ_2:           highlighter.TokenTypeOperator,
		python3.Python3LexerAT:                 highlighter.TokenTypeText,
		python3.Python3LexerARROW:              highlighter.TokenTypeText,
		python3.Python3LexerADD_ASSIGN:         highlighter.TokenTypeText,
		python3.Python3LexerSUB_ASSIGN:         highlighter.TokenTypeText,
		python3.Python3LexerMULT_ASSIGN:        highlighter.TokenTypeText,
		python3.Python3LexerAT_ASSIGN:          highlighter.TokenTypeText,
		python3.Python3LexerDIV_ASSIGN:         highlighter.TokenTypeText,
		python3.Python3LexerMOD_ASSIGN:         highlighter.TokenTypeText,
		python3.Python3LexerAND_ASSIGN:         highlighter.TokenTypeText,
		python3.Python3LexerOR_ASSIGN:          highlighter.TokenTypeText,
		python3.Python3LexerXOR_ASSIGN:         highlighter.TokenTypeText,
		python3.Python3LexerLEFT_SHIFT_ASSIGN:  highlighter.TokenTypeText,
		python3.Python3LexerRIGHT_SHIFT_ASSIGN: highlighter.TokenTypeText,
		python3.Python3LexerPOWER_ASSIGN:       highlighter.TokenTypeText,
		python3.Python3LexerIDIV_ASSIGN:        highlighter.TokenTypeText,
		python3.Python3LexerSKIP_:              highlighter.TokenTypeText,
		python3.Python3LexerUNKNOWN_CHAR:       highlighter.TokenTypeText,
		python3.Python3LexerCOMMENT:            highlighter.TokenTypeComment,
	}

	return Python3Lexer{
		symbolicMap: symbolicMap,
	}

}
