package ast

import "fmt"

type Float struct {
	ValueType ValueType
	Value     float64
}

func (f *Float) Type() ValueType {
	return f.ValueType
}

func (f *Float) String() string {
	return fmt.Sprintf("%f", f.Value)
}
