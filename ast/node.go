package ast

import (
	"github.com/skanehira/go-json/token"
)

type Node interface {
	String() string
	TokenType() token.TokenType
}
