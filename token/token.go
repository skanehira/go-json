package token

type TokenType string

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	LBRACE   = "{"
	RBRACE   = "}"
	LBRACKET = "["
	RBRACKET = "]"
	COMMA    = ","
	COLON    = ":"

	TRUE  = "TRUE"
	FALSE = "FALSE"
	NULL  = "NULL"

	OBJECT = "OBJECT"
	ARRAY  = "ARRAY"
	INT    = "INT"

	LETTER = "LETTER"
)

type Token struct {
	Type    TokenType
	Literal string
}
