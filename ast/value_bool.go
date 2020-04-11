package ast

import "github.com/skanehira/go-json/token"

type Bool struct {
	Value bool
	Type  token.TokenType
}

func (b Bool) String() string {
	if b.Value == true {
		return "true"
	}
	return "false"
}

func (b Bool) TokenType() token.TokenType {
	return b.Type
}
