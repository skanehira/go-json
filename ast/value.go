package ast

import "github.com/skanehira/go-json/token"

type Value interface {
	String() string
	TokenType() token.TokenType
}
