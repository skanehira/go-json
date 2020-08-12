package ast

import (
	"bytes"
	"fmt"
	"strconv"
)

type Value interface {
	String() string
}

type JSON struct {
	Values []Value
}

func (j *JSON) String() string {
	var out bytes.Buffer
	for _, v := range j.Values {
		out.WriteString(v.String())
	}
	return out.String()
}

type Element struct {
	Key   string
	Value Value
}

func (el Element) String() string {
	var out bytes.Buffer
	out.WriteString("\"" + el.Key + "\"")
	out.WriteString(":")
	out.WriteString(el.Value.String())
	return out.String()
}

type Object struct {
	Elements []Element
}

func (o Object) String() string {
	var out bytes.Buffer
	out.WriteString("{")
	for i, e := range o.Elements {
		out.WriteString(e.String())
		if len(o.Elements)-1 > i {
			out.WriteString(",")
		}
	}
	out.WriteString("}")
	return out.String()
}

type String struct {
	Value string
}

func (s String) String() string {
	return "\"" + s.Value + "\""
}

type Array struct {
	Values []Value
}

func (a Array) String() string {
	var out bytes.Buffer
	out.WriteString("[")

	for i, v := range a.Values {
		out.WriteString(v.String())
		if len(a.Values)-1 > i {
			out.WriteString(",")
		}
	}

	out.WriteString("]")
	return out.String()
}

type Int struct {
	Value int64
}

func (i Int) String() string {
	return fmt.Sprintf("%d", i.Value)
}

type Bool struct {
	Value bool
}

func (b Bool) String() string {
	return strconv.FormatBool(b.Value)
}

type Null struct {
}

func (n Null) String() string {
	return "null"
}
