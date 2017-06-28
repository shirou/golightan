package lexer

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"

	sqlite "github.com/shirou/antlr-grammars-v4-go/sqlite"

	"github.com/shirou/highlighter"
)

type SQLiteLexer struct {
	lexer       antlr.Lexer
	ruleMap     TypeMap
	literalMap  TypeMap
	symbolicMap TypeMap
}

func (l SQLiteLexer) Tokenize(input antlr.CharStream) (highlighter.Tokens, error) {
	le := sqlite.NewSQLiteLexer(input)
	return CommonTokenize(le, l.symbolicMap)
}

func NewSQLiteLexer() Lexer {
	symbolicMap := TypeMap{
		sqlite.SQLiteParserSCOL:                highlighter.TokenTypeText,
		sqlite.SQLiteParserDOT:                 highlighter.TokenTypeText,
		sqlite.SQLiteParserOPEN_PAR:            highlighter.TokenTypePunctuation,
		sqlite.SQLiteParserCLOSE_PAR:           highlighter.TokenTypePunctuation,
		sqlite.SQLiteParserCOMMA:               highlighter.TokenTypeText,
		sqlite.SQLiteParserASSIGN:              highlighter.TokenTypeText,
		sqlite.SQLiteParserSTAR:                highlighter.TokenTypeOperator,
		sqlite.SQLiteParserPLUS:                highlighter.TokenTypeOperator,
		sqlite.SQLiteParserMINUS:               highlighter.TokenTypeOperator,
		sqlite.SQLiteParserTILDE:               highlighter.TokenTypeText,
		sqlite.SQLiteParserPIPE2:               highlighter.TokenTypeText,
		sqlite.SQLiteParserDIV:                 highlighter.TokenTypeOperator,
		sqlite.SQLiteParserMOD:                 highlighter.TokenTypeOperator,
		sqlite.SQLiteParserLT2:                 highlighter.TokenTypeText,
		sqlite.SQLiteParserGT2:                 highlighter.TokenTypeText,
		sqlite.SQLiteParserAMP:                 highlighter.TokenTypeText,
		sqlite.SQLiteParserPIPE:                highlighter.TokenTypeText,
		sqlite.SQLiteParserLT:                  highlighter.TokenTypeOperator,
		sqlite.SQLiteParserLT_EQ:               highlighter.TokenTypeOperator,
		sqlite.SQLiteParserGT:                  highlighter.TokenTypeOperator,
		sqlite.SQLiteParserGT_EQ:               highlighter.TokenTypeOperator,
		sqlite.SQLiteParserEQ:                  highlighter.TokenTypeOperator,
		sqlite.SQLiteParserNOT_EQ1:             highlighter.TokenTypeOperator,
		sqlite.SQLiteParserNOT_EQ2:             highlighter.TokenTypeText,
		sqlite.SQLiteParserK_ABORT:             highlighter.TokenTypeText,
		sqlite.SQLiteParserK_ACTION:            highlighter.TokenTypeText,
		sqlite.SQLiteParserK_ADD:               highlighter.TokenTypeText,
		sqlite.SQLiteParserK_AFTER:             highlighter.TokenTypeText,
		sqlite.SQLiteParserK_ALL:               highlighter.TokenTypeText,
		sqlite.SQLiteParserK_ALTER:             highlighter.TokenTypeText,
		sqlite.SQLiteParserK_ANALYZE:           highlighter.TokenTypeText,
		sqlite.SQLiteParserK_AND:               highlighter.TokenTypeText,
		sqlite.SQLiteParserK_AS:                highlighter.TokenTypeText,
		sqlite.SQLiteParserK_ASC:               highlighter.TokenTypeText,
		sqlite.SQLiteParserK_ATTACH:            highlighter.TokenTypeText,
		sqlite.SQLiteParserK_AUTOINCREMENT:     highlighter.TokenTypeText,
		sqlite.SQLiteParserK_BEFORE:            highlighter.TokenTypeText,
		sqlite.SQLiteParserK_BEGIN:             highlighter.TokenTypeText,
		sqlite.SQLiteParserK_BETWEEN:           highlighter.TokenTypeText,
		sqlite.SQLiteParserK_BY:                highlighter.TokenTypeText,
		sqlite.SQLiteParserK_CASCADE:           highlighter.TokenTypeText,
		sqlite.SQLiteParserK_CASE:              highlighter.TokenTypeText,
		sqlite.SQLiteParserK_CAST:              highlighter.TokenTypeText,
		sqlite.SQLiteParserK_CHECK:             highlighter.TokenTypeText,
		sqlite.SQLiteParserK_COLLATE:           highlighter.TokenTypeText,
		sqlite.SQLiteParserK_COLUMN:            highlighter.TokenTypeText,
		sqlite.SQLiteParserK_COMMIT:            highlighter.TokenTypeText,
		sqlite.SQLiteParserK_CONFLICT:          highlighter.TokenTypeText,
		sqlite.SQLiteParserK_CONSTRAINT:        highlighter.TokenTypeText,
		sqlite.SQLiteParserK_CREATE:            highlighter.TokenTypeKeyword,
		sqlite.SQLiteParserK_CROSS:             highlighter.TokenTypeText,
		sqlite.SQLiteParserK_CURRENT_DATE:      highlighter.TokenTypeText,
		sqlite.SQLiteParserK_CURRENT_TIME:      highlighter.TokenTypeText,
		sqlite.SQLiteParserK_CURRENT_TIMESTAMP: highlighter.TokenTypeText,
		sqlite.SQLiteParserK_DATABASE:          highlighter.TokenTypeKeyword,
		sqlite.SQLiteParserK_DEFAULT:           highlighter.TokenTypeText,
		sqlite.SQLiteParserK_DEFERRABLE:        highlighter.TokenTypeText,
		sqlite.SQLiteParserK_DEFERRED:          highlighter.TokenTypeText,
		sqlite.SQLiteParserK_DELETE:            highlighter.TokenTypeText,
		sqlite.SQLiteParserK_DESC:              highlighter.TokenTypeText,
		sqlite.SQLiteParserK_DETACH:            highlighter.TokenTypeText,
		sqlite.SQLiteParserK_DISTINCT:          highlighter.TokenTypeText,
		sqlite.SQLiteParserK_DROP:              highlighter.TokenTypeText,
		sqlite.SQLiteParserK_EACH:              highlighter.TokenTypeText,
		sqlite.SQLiteParserK_ELSE:              highlighter.TokenTypeText,
		sqlite.SQLiteParserK_END:               highlighter.TokenTypeText,
		sqlite.SQLiteParserK_ESCAPE:            highlighter.TokenTypeText,
		sqlite.SQLiteParserK_EXCEPT:            highlighter.TokenTypeText,
		sqlite.SQLiteParserK_EXCLUSIVE:         highlighter.TokenTypeText,
		sqlite.SQLiteParserK_EXISTS:            highlighter.TokenTypeText,
		sqlite.SQLiteParserK_EXPLAIN:           highlighter.TokenTypeText,
		sqlite.SQLiteParserK_FAIL:              highlighter.TokenTypeText,
		sqlite.SQLiteParserK_FOR:               highlighter.TokenTypeText,
		sqlite.SQLiteParserK_FOREIGN:           highlighter.TokenTypeText,
		sqlite.SQLiteParserK_FROM:              highlighter.TokenTypeKeyword,
		sqlite.SQLiteParserK_FULL:              highlighter.TokenTypeText,
		sqlite.SQLiteParserK_GLOB:              highlighter.TokenTypeText,
		sqlite.SQLiteParserK_GROUP:             highlighter.TokenTypeText,
		sqlite.SQLiteParserK_HAVING:            highlighter.TokenTypeText,
		sqlite.SQLiteParserK_IF:                highlighter.TokenTypeText,
		sqlite.SQLiteParserK_IGNORE:            highlighter.TokenTypeText,
		sqlite.SQLiteParserK_IMMEDIATE:         highlighter.TokenTypeText,
		sqlite.SQLiteParserK_IN:                highlighter.TokenTypeText,
		sqlite.SQLiteParserK_INDEX:             highlighter.TokenTypeText,
		sqlite.SQLiteParserK_INDEXED:           highlighter.TokenTypeText,
		sqlite.SQLiteParserK_INITIALLY:         highlighter.TokenTypeText,
		sqlite.SQLiteParserK_INNER:             highlighter.TokenTypeText,
		sqlite.SQLiteParserK_INSERT:            highlighter.TokenTypeText,
		sqlite.SQLiteParserK_INSTEAD:           highlighter.TokenTypeText,
		sqlite.SQLiteParserK_INTERSECT:         highlighter.TokenTypeText,
		sqlite.SQLiteParserK_INTO:              highlighter.TokenTypeText,
		sqlite.SQLiteParserK_IS:                highlighter.TokenTypeText,
		sqlite.SQLiteParserK_ISNULL:            highlighter.TokenTypeText,
		sqlite.SQLiteParserK_JOIN:              highlighter.TokenTypeText,
		sqlite.SQLiteParserK_KEY:               highlighter.TokenTypeKeyword,
		sqlite.SQLiteParserK_LEFT:              highlighter.TokenTypeText,
		sqlite.SQLiteParserK_LIKE:              highlighter.TokenTypeText,
		sqlite.SQLiteParserK_LIMIT:             highlighter.TokenTypeText,
		sqlite.SQLiteParserK_MATCH:             highlighter.TokenTypeText,
		sqlite.SQLiteParserK_NATURAL:           highlighter.TokenTypeText,
		sqlite.SQLiteParserK_NO:                highlighter.TokenTypeText,
		sqlite.SQLiteParserK_NOT:               highlighter.TokenTypeOperatorWord,
		sqlite.SQLiteParserK_NOTNULL:           highlighter.TokenTypeOperatorWord,
		sqlite.SQLiteParserK_NULL:              highlighter.TokenTypeKeyword,
		sqlite.SQLiteParserK_OF:                highlighter.TokenTypeText,
		sqlite.SQLiteParserK_OFFSET:            highlighter.TokenTypeText,
		sqlite.SQLiteParserK_ON:                highlighter.TokenTypeText,
		sqlite.SQLiteParserK_OR:                highlighter.TokenTypeText,
		sqlite.SQLiteParserK_ORDER:             highlighter.TokenTypeText,
		sqlite.SQLiteParserK_OUTER:             highlighter.TokenTypeText,
		sqlite.SQLiteParserK_PLAN:              highlighter.TokenTypeText,
		sqlite.SQLiteParserK_PRAGMA:            highlighter.TokenTypeText,
		sqlite.SQLiteParserK_PRIMARY:           highlighter.TokenTypeKeyword,
		sqlite.SQLiteParserK_QUERY:             highlighter.TokenTypeText,
		sqlite.SQLiteParserK_RAISE:             highlighter.TokenTypeText,
		sqlite.SQLiteParserK_RECURSIVE:         highlighter.TokenTypeText,
		sqlite.SQLiteParserK_REFERENCES:        highlighter.TokenTypeText,
		sqlite.SQLiteParserK_REGEXP:            highlighter.TokenTypeText,
		sqlite.SQLiteParserK_REINDEX:           highlighter.TokenTypeText,
		sqlite.SQLiteParserK_RELEASE:           highlighter.TokenTypeText,
		sqlite.SQLiteParserK_RENAME:            highlighter.TokenTypeText,
		sqlite.SQLiteParserK_REPLACE:           highlighter.TokenTypeText,
		sqlite.SQLiteParserK_RESTRICT:          highlighter.TokenTypeText,
		sqlite.SQLiteParserK_RIGHT:             highlighter.TokenTypeText,
		sqlite.SQLiteParserK_ROLLBACK:          highlighter.TokenTypeText,
		sqlite.SQLiteParserK_ROW:               highlighter.TokenTypeText,
		sqlite.SQLiteParserK_SAVEPOINT:         highlighter.TokenTypeText,
		sqlite.SQLiteParserK_SELECT:            highlighter.TokenTypeKeyword,
		sqlite.SQLiteParserK_SET:               highlighter.TokenTypeText,
		sqlite.SQLiteParserK_TABLE:             highlighter.TokenTypeNameClass,
		sqlite.SQLiteParserK_TEMP:              highlighter.TokenTypeText,
		sqlite.SQLiteParserK_TEMPORARY:         highlighter.TokenTypeText,
		sqlite.SQLiteParserK_THEN:              highlighter.TokenTypeText,
		sqlite.SQLiteParserK_TO:                highlighter.TokenTypeText,
		sqlite.SQLiteParserK_TRANSACTION:       highlighter.TokenTypeText,
		sqlite.SQLiteParserK_TRIGGER:           highlighter.TokenTypeText,
		sqlite.SQLiteParserK_UNION:             highlighter.TokenTypeText,
		sqlite.SQLiteParserK_UNIQUE:            highlighter.TokenTypeText,
		sqlite.SQLiteParserK_UPDATE:            highlighter.TokenTypeText,
		sqlite.SQLiteParserK_USING:             highlighter.TokenTypeText,
		sqlite.SQLiteParserK_VACUUM:            highlighter.TokenTypeText,
		sqlite.SQLiteParserK_VALUES:            highlighter.TokenTypeText,
		sqlite.SQLiteParserK_VIEW:              highlighter.TokenTypeText,
		sqlite.SQLiteParserK_VIRTUAL:           highlighter.TokenTypeText,
		sqlite.SQLiteParserK_WHEN:              highlighter.TokenTypeText,
		sqlite.SQLiteParserK_WHERE:             highlighter.TokenTypeText,
		sqlite.SQLiteParserK_WITH:              highlighter.TokenTypeText,
		sqlite.SQLiteParserK_WITHOUT:           highlighter.TokenTypeText,
		sqlite.SQLiteParserK_TEXT:              highlighter.TokenTypeNameClass,
		sqlite.SQLiteParserK_INT:               highlighter.TokenTypeNameClass,
		sqlite.SQLiteParserK_INTEGER:           highlighter.TokenTypeNameClass,
		sqlite.SQLiteParserK_CHAR:              highlighter.TokenTypeNameClass,
		sqlite.SQLiteParserK_REAL:              highlighter.TokenTypeNameClass,
		sqlite.SQLiteParserK_BLOB:              highlighter.TokenTypeNameClass,
		sqlite.SQLiteParserIDENTIFIER:          highlighter.TokenTypeText,
		sqlite.SQLiteParserNUMERIC_LITERAL:     highlighter.TokenTypeNumberInteger,
		sqlite.SQLiteParserBIND_PARAMETER:      highlighter.TokenTypeText,
		sqlite.SQLiteParserSTRING_LITERAL:      highlighter.TokenTypeText,
		sqlite.SQLiteParserBLOB_LITERAL:        highlighter.TokenTypeText,
		sqlite.SQLiteParserSINGLE_LINE_COMMENT: highlighter.TokenTypeText,
		sqlite.SQLiteParserMULTILINE_COMMENT:   highlighter.TokenTypeText,
		sqlite.SQLiteParserSPACES:              highlighter.TokenTypeWhitespace,
		sqlite.SQLiteParserUNEXPECTED_CHAR:     highlighter.TokenTypeText,
	}

	return SQLiteLexer{
		symbolicMap: symbolicMap,
	}

}
