package ast

import (
	"testing"

	"github.com/skanehira/go-json/token"
)

func TestString(t *testing.T) {
	str := String{
		Type:  token.String,
		Value: "",
	}

	want := "\"\""
	if str.String() != want {
		t.Fatalf("want=%s, got=%s", want, str.String())
	}

	want = `"a"`
	str = String{
		Type:  token.String,
		Value: "a",
	}

	if str.String() != want {
		t.Fatalf("want=%s, got=%s", want, str.String())
	}

	if str.TokenType() != token.String {
		t.Fatalf("want=%s, got=%s", token.String, str.TokenType().String())
	}
}
