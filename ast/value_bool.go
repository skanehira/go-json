package ast

import "strconv"

type Bool struct {
	ValueType ValueType
	Value     bool
}

func (b Bool) Type() ValueType {
	return b.ValueType
}

func (b Bool) String() string {
	return strconv.FormatBool(b.Value)
}
