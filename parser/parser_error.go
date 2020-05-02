package parser

import "errors"

var (
	ErrInvalidNextToken = errors.New("invalid next token")
	ErrInvalidToken     = errors.New("invalid token")
)
