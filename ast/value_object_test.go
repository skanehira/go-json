package ast

import (
	"testing"

	"github.com/skanehira/go-json/token"
)

func TestObject(t *testing.T) {
	o := Object{
		Type: token.Object,
	}

	want := "{}"
	if o.String() != "{}" {
		t.Fatalf("want=%s, got=%s", want, o.String())
	}

	o = Object{
		Type: token.Object,
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

	if o.TokenType() != token.Object {
		t.Fatalf("want=%s, got=%s", token.String, o.TokenType().String())
	}
}
