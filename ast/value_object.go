package ast

import (
	"strings"

	"github.com/skanehira/go-json/token"
)

type Object struct {
	Elements map[string]Value
	Type     token.TokenType
}

func (o Object) String() string {
	var out strings.Builder
	out.WriteString("{")

	elements := []string{}
	for name, node := range o.Elements {
		e := "\"" + name + "\":" + node.String()
		elements = append(elements, e)
	}

	out.WriteString(strings.Join(elements, ","))
	out.WriteString("}")
	return out.String()
}

func (o Object) TokenType() token.TokenType {
	return o.Type
}
