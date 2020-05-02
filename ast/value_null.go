package ast

type Null struct {
	ValueType ValueType
}

func (n Null) Type() ValueType {
	return n.ValueType
}

func (n Null) String() string {
	return "null"
}
