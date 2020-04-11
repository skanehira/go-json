package ast

import (
	"strconv"

	"github.com/skanehira/go-json/token"
)

type Int struct {
	Value int
	Type  token.TokenType
}

func (i Int) String() string {
	return strconv.Itoa(i.Value)
}

func (i Int) TokenType() token.TokenType {
	return i.Type
}
