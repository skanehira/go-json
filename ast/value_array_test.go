package ast

import (
	"testing"

	"github.com/skanehira/go-json/token"
)

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
