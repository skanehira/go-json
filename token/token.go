package token

type TokenType int

const (
	_ TokenType = iota
	Int
	String
	Bool
	Object
	Array
	Null
)

type Token struct {
	Type  TokenType
	Value string
}

func (t Token) String() string {
	return t.Value
}
