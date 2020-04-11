package ast

import "github.com/skanehira/go-json/token"

type Null struct {
	Type token.TokenType
}

func (n Null) String() string {
	return "null"
}

func (n Null) TokenType() token.TokenType {
	return n.Type
}
