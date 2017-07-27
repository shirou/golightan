package languages

import (
	xml "github.com/shirou/antlr-grammars-v4-go/xml"

	"github.com/shirou/golightan"
)

func NewXMLTokenMap() TokenMap {
	return TokenMap{
		symbolicMap: newXMLSymbolicMap(),
		ruleMap:     newXMLRuleMap(),
	}
}

func newXMLRuleMap() RuleMap {
	return RuleMap{
		Rule{xml.XMLParserRULE_attribute, xml.XMLParserName}:     golightan.TokenTypeNameAttribute,
		Rule{xml.XMLParserRULE_attribute, xml.XMLParserSTRING}:   golightan.TokenTypeString,
		Rule{xml.XMLParserRULE_attribute, xml.XMLParserEQUALS}:   golightan.TokenTypeText,
		Rule{xml.XMLParserRULE_element, xml.XMLParserSTRING}:     golightan.TokenTypeNameTag,
		Rule{xml.XMLParserRULE_prolog, xml.XMLParserXMLDeclOpen}: golightan.TokenTypeNameTag,
	}
}

func newXMLSymbolicMap() TypeMap {
	return TypeMap{
		xml.XMLParserName:          golightan.TokenTypeNameTag,
		xml.XMLParserOPEN:          golightan.TokenTypeNameTag,
		xml.XMLParserCLOSE:         golightan.TokenTypeNameTag,
		xml.XMLParserSPECIAL_CLOSE: golightan.TokenTypeNameTag,
		xml.XMLParserSLASH_CLOSE:   golightan.TokenTypeNameTag,
		xml.XMLParserSLASH:         golightan.TokenTypeNameTag,
	}
}
