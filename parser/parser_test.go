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
		t.Fatalf("want=%T, got=%T", ast.Int{}, p.Parse())
	}

	want := 12345
	if value.Value != want {
		t.Fatalf("want=%d, got=%d", want, value.Value)
	}
}

func TestParseObject(t *testing.T) {
	tests := []struct {
		input []byte
		want  string
	}{
		{
			input: []byte(`{"a": "b","c": 1}`),
			want:  `{"a":"b","c":1}`,
		},
		{
			input: []byte(`{"a": 1,"b": 1}`),
			want:  `{"a":1,"b":1}`,
		},
		{
			input: []byte(`{"a": {"b": "c", "a": 1}}`),
			want:  `{"a":{"b":"c","a":1}}`,
		},
	}

	for i, tt := range tests {
		p := &Parser{
			input:        tt.input,
			position:     0,
			readPosition: 0,
		}

		value, ok := p.Parse().(ast.Object)

		if !ok {
			t.Fatalf("case[%d] want=%T, got=%T", i, ast.Object{}, p.Parse())
		}

		if value.String() != tt.want {
			t.Fatalf("case[%d], want=%s, got=%s", i, tt.want, value.String())
		}
	}

}
