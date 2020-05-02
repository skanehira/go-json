package parser

import "testing"

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
			t.Fatalf("[case %d] unexpected err: want=%#v, got=%#v", i, tt.err, err)
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
		err   error
	}{
		{
			input: []byte(`{"a":"b"}`),
			want:  `{"a":"b"}`,
			err:   nil,
		},
		{
			input: []byte(`{}`),
			want:  `{}`,
			err:   nil,
		},
	}

	for i, tt := range tests {
		p := NewParser(tt.input).Init()
		v, err := p.ParseObject()
		if err != tt.err {
			t.Fatalf("[case %d] unexpected err: want=%#v, got=%#v", i, tt.err, err)
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
	}

	for i, tt := range tests {
		p := NewParser(tt.input).Init()
		_, got := p.ParseObject()
		if got.Error() != tt.want {
			t.Fatalf("[case %d] unexpected value: want=%s, got=%s", i, tt.want, got)
		}
	}

}
