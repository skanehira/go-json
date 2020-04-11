package ast

import (
	"testing"

	"github.com/skanehira/go-json/token"
)

func TestString(t *testing.T) {
	want := `"a"`
	str := String{
		Value: "a",
	}

	if str.String() != want {
		t.Fatalf("want=%s, got=%s", want, str.String())
	}
}

func TestObject(t *testing.T) {
	o := Object{
		Type: token.Object,
	}

	want := "{}"
	if o.String() != "{}" {
		t.Fatalf("want=%s, got=%s", want, o.String())
	}

	o = Object{
		Elements: []Element{
			{
				Name: "name",
				Value: String{
					Type:  token.String,
					Value: "gorilla",
				},
			},
			{
				Name: "age",
				Value: Int{
					Type:  token.Int,
					Value: 1,
				},
			},
			{
				Name: "is_gorilla",
				Value: Bool{
					Type:  token.Bool,
					Value: true,
				},
			},
			{
				Name: "is_human",
				Value: Bool{
					Type:  token.Bool,
					Value: false,
				},
			},
			{
				Name: "banana",
				Value: Null{
					Type: token.Null,
				},
			},
		},
	}

	want = `{"name":"gorilla","age":1,"is_gorilla":true,"is_human":false,"banana":null}`

	if o.String() != want {
		t.Fatalf("want=%s, got=%s", want, o.String())
	}
}

func TestArray(t *testing.T) {
	a := Array{
		Values: []Value{},
		Type:   token.Array,
	}

	want := "[]"
	if a.String() != want {
		t.Fatalf("want=%s, got=%s", want, a.String())
	}
}
