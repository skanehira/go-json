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

func TestArray(t *testing.T) {
	a := Array{
		Values: []Value{},
		Type:   token.Array,
	}

	want := "[]"
	if a.String() != want {
		t.Fatalf("want=%s, got=%s", want, a.String())
	}

	a = Array{
		Type: token.Array,
		Values: []Value{
			Object{
				Type: token.Object,
				Elements: []Element{
					{
						Name: "name",
						Value: String{
							Type:  token.String,
							Value: "gorilla",
						},
					},
				},
			},
			Array{
				Type: token.Array,
				Values: []Value{
					Int{
						Type:  token.Int,
						Value: 1,
					},
				},
			},
			Int{
				Type:  token.Int,
				Value: 10,
			},
			String{
				Type:  token.String,
				Value: "gorilla",
			},
			Bool{
				Type:  token.Bool,
				Value: false,
			},
			Null{
				Type: token.Null,
			},
		},
	}

	want = `[{"name":"gorilla"},[1],10,"gorilla",false,null]`

	if a.String() != want {
		t.Fatalf("want=%s, got=%s", want, a.String())
	}

	if a.TokenType() != token.Array {
		t.Fatalf("want=%s, got=%s", token.String, a.TokenType().String())
	}
}
