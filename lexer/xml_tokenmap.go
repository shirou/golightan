package lexer

import (
	xml "github.com/shirou/antlr-grammars-v4-go/xml"

	"github.com/shirou/highlighter"
)

func NewXMLTokenMap() TokenMap {
	return TokenMap{
		symbolicMap: newXMLSymbolicMap(),
		ruleMap:     newXMLRuleMap(),
	}
}

func newXMLRuleMap() RuleMap {
	return RuleMap{
		Rule{xml.XMLParserRULE_attribute, xml.XMLParserName}:     highlighter.TokenTypeNameAttribute,
		Rule{xml.XMLParserRULE_attribute, xml.XMLParserSTRING}:   highlighter.TokenTypeString,
		Rule{xml.XMLParserRULE_attribute, xml.XMLParserEQUALS}:   highlighter.TokenTypeText,
		Rule{xml.XMLParserRULE_element, xml.XMLParserSTRING}:     highlighter.TokenTypeNameTag,
		Rule{xml.XMLParserRULE_prolog, xml.XMLParserXMLDeclOpen}: highlighter.TokenTypeNameTag,
	}
}

func newXMLSymbolicMap() TypeMap {
	return TypeMap{
		xml.XMLParserName:          highlighter.TokenTypeNameTag,
		xml.XMLParserOPEN:          highlighter.TokenTypeNameTag,
		xml.XMLParserCLOSE:         highlighter.TokenTypeNameTag,
		xml.XMLParserSPECIAL_CLOSE: highlighter.TokenTypeNameTag,
		xml.XMLParserSLASH_CLOSE:   highlighter.TokenTypeNameTag,
		xml.XMLParserSLASH:         highlighter.TokenTypeNameTag,
	}
}
