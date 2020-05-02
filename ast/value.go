package ast

type ValueType int

const (
	Unknown ValueType = iota
	IntType
	FloatType
	StringType
	TrueType
	FalseType
	NullType
	ObjectType
	ArrayType
)

type Value interface {
	Type() ValueType
	String() string
}
