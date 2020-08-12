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

	INT   = "INT"
	FLOAT = "FLOAT"

	LETTER = "LETTER"
)

type Token struct {
	Type    TokenType
	Literal string
}
