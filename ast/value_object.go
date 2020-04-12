package ast

import (
	"fmt"
	"strings"

	"github.com/skanehira/go-json/token"
)

type Element struct {
	Name  string
	Value Value
}

type Object struct {
	Elements []Element
	Type     token.TokenType
}

func (o Object) String() string {
	var out strings.Builder
	out.WriteString("{")

	elements := []string{}
	for _, e := range o.Elements {
		if e.Value == nil {
			continue
		}
		elements = append(elements, fmt.Sprintf("\"%s\":%s", e.Name, e.Value.String()))
	}

	out.WriteString(strings.Join(elements, ","))

	out.WriteString("}")
	return out.String()
}

func (o Object) TokenType() token.TokenType {
	return o.Type
}
