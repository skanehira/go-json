package ast

type Array struct {
	ValueType ValueType
	Values    []Value
}

func (a Array) Type() ValueType {
	return a.ValueType
}

func (a Array) String() string {
	value := "["

	for _, v := range a.Values {
		value += v.String()
	}

	value += "]"
	return value
}
