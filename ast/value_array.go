package ast

import (
	"strings"

	"github.com/skanehira/go-json/token"
)

type Array struct {
	Values []Value
	Type   token.TokenType
}

func (a Array) String() string {
	var out strings.Builder
	out.WriteString("[")

	values := []string{}
	for _, value := range a.Values {
		values = append(values, value.String())
	}
	out.WriteString(strings.Join(values, ","))
	out.WriteString("]")
	return out.String()
}

func (a Array) TokenType() token.TokenType {
	return a.Type
}
