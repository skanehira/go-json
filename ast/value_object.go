package ast

import "fmt"

type Object struct {
	ValueType ValueType
	Key       string
	Value     Value
}

func (t Object) Type() ValueType {
	return t.ValueType
}

func (t Object) String() string {
	if t.Key == "" && t.Value == nil {
		return "{}"
	}
	return fmt.Sprintf("{\"%s\":%s}", t.Key, t.Value.String())
}
