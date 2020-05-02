package ast

type String struct {
	ValueType ValueType
	Value     string
}

func (s String) Type() ValueType {
	return s.ValueType
}

func (s String) String() string {
	return "\"" + s.Value + "\""
}
