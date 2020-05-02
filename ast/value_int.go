package ast

import "strconv"

type Integer struct {
	ValueType ValueType
	Value     int64
}

func (i Integer) Type() ValueType {
	return i.ValueType
}

func (i Integer) String() string {
	return strconv.Itoa(int(i.Value))
}
