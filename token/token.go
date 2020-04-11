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

var tokenMap = map[TokenType]string{
	Int:    "Int",
	String: "String",
	Bool:   "Bool",
	Object: "Object",
	Array:  "Array",
	Null:   "Null",
}

func (t TokenType) String() string {
	return tokenMap[t]
}

type Token struct {
	Type  TokenType
	Value string
}

func (t Token) String() string {
	return t.Value
}
