package parser

import (
	"testing"

	"github.com/skanehira/go-json/ast"
)

func TestParserString(t *testing.T) {
	tests := []struct {
		input []byte
		want  string
		err   error
	}{
		{
			input: []byte(`"a"`),
			want:  `"a"`,
			err:   nil,
		},
		{
			input: []byte(`""`),
			want:  `""`,
			err:   nil,
		},
		{
			input: []byte(`"`),
			want:  `""`,
			err:   ErrInvalidNextToken,
		},
		{
			input: []byte(`"a`),
			want:  `""`,
			err:   ErrInvalidNextToken,
		},
	}

	for i, tt := range tests {
		p := NewParser(tt.input).Init()
		v, err := p.ParseString()
		if err != tt.err {
			t.Fatalf("[case %d] unexpected err: want=%s, got=%s", i, tt.err, err)
		}

		got := v.String()
		if got != tt.want {
			t.Fatalf("[case %d] unexpected value: want=%#v, got=%#v", i, tt.want, got)
		}
	}
}

func TestParseObject(t *testing.T) {
	tests := []struct {
		input []byte
		want  string
	}{
		{
			input: []byte(`{"a":"b"}`),
			want:  `{"a":"b"}`,
		},
		{
			input: []byte(`{}`),
			want:  `{}`,
		},
	}

	for i, tt := range tests {
		p := NewParser(tt.input).Init()
		v, err := p.ParseObject()
		if err != nil {
			t.Fatalf("[case %d] unexpected err: want=nil, got=%#v", i, err)
		}

		got := v.String()
		if got != tt.want {
			t.Fatalf("[case %d] unexpected value: want=%#v, got=%#v", i, tt.want, got)
		}
	}
}

func TestParseObjectFailed(t *testing.T) {
	tests := []struct {
		input []byte
		want  string
	}{
		{
			input: []byte(`{"a:"b"}`),
			want:  `current token=b: invalid token`,
		},
		{
			input: []byte(`{a:"b"}`),
			want:  `next token=a: invalid next token`,
		},
		{
			input: []byte(`{:"b"}`),
			want:  `next token=:: invalid next token`,
		},
		{
			input: []byte(`{"b":}`),
			want:  `current token=}: invalid token`,
		},
		{
			input: []byte(`{"b"`),
			want:  `current token=null(0): invalid token`,
		},
		{
			input: []byte(`{"b":`),
			want:  `current token=null(0): invalid token`,
		},
	}

	for i, tt := range tests {
		p := NewParser(tt.input).Init()
		_, got := p.ParseObject()
		if got.Error() != tt.want {
			t.Fatalf("[case %d] unexpected value: want=%#v, got=%#v", i, tt.want, got)
		}
	}
}

func TestParseInteger(t *testing.T) {
	tests := []struct {
		input []byte
		want  int64
	}{
		{
			input: []byte(`123`),
			want:  123,
		},
		{
			input: []byte(`999999`),
			want:  999999,
		},
	}

	for i, tt := range tests {
		p := NewParser(tt.input).Init()
		got, err := p.ParseInteger()
		if err != nil {
			t.Fatalf("[case %d] unexpected error: want=nil, got=%s", i, err)
		}

		value := got.(ast.Integer)
		if value.Value != tt.want {
			t.Fatalf("[case %d] unexpected value: want=%d, got=%d", i, tt.want, value.Value)
		}
	}
}

func TestParseIntegerFailed(t *testing.T) {
	tests := []struct {
		input []byte
		want  string
	}{
		{
			input: []byte(`a`),
			want:  `failed to convert "a" to integer`,
		},
		{
			input: []byte(`1a`),
			want:  `failed to convert "1a" to integer`,
		},
	}

	for i, tt := range tests {
		p := NewParser(tt.input).Init()
		_, err := p.ParseInteger()

		if err.Error() != tt.want {
			t.Fatalf("[case %d] unexpected error: want=%s, got=%s", i, tt.want, err.Error())
		}
	}

}
