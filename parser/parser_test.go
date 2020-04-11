package parser

import (
	"testing"

	"github.com/skanehira/go-json/ast"
)

func TestParseString(t *testing.T) {
	input := []byte(`"gorilla"`)
	p := &Parser{
		input:        input,
		position:     0,
		readPosition: 0,
	}

	value, ok := p.Parse().(ast.String)
	if !ok {
		t.Fatalf("want=%T, got=%T", ast.String{}, p.Parse())
	}

	want := "gorilla"
	if value.Value != want {
		t.Fatalf("want=%s, got=%s", want, value.Value)
	}
}

func TestParseInt(t *testing.T) {
	p := &Parser{
		input:        []byte(`12345`),
		position:     0,
		readPosition: 0,
	}

	value, ok := p.Parse().(ast.Int)
	if !ok {
		t.Fatalf("want=%T, got=%T", ast.String{}, p.Parse())
	}

	want := 12345
	if value.Value != want {
		t.Fatalf("want=%d, got=%d", want, value.Value)
	}
}
