package ast

import "github.com/skanehira/go-json/token"

type String struct {
	Value string
	Type  token.TokenType
}

func (s String) String() string {
	return "\"" + s.Value + "\""
}

func (s String) TokenType() token.TokenType {
	return s.Type
}
