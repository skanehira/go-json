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
		Elements: map[string]Value{
			"name": String{
				Type:  token.String,
				Value: "gorilla",
			},
			"age": Int{
				Type:  token.Int,
				Value: 1,
			},
		},
	}

	want := `{"name":"gorilla","age":1}`

	if o.String() != want {
		t.Fatalf("want=%s, got=%s", want, o.String())
	}
}
