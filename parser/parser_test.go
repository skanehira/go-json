package parser

import "testing"

func TestParserString(t *testing.T) {
	tests := []struct {
		input []byte
		want  string
		err   error
	}{
		{
			input: []byte(`"string"`),
			want:  `"string"`,
			err:   nil,
		},
		{
			input: []byte(`""`),
			want:  `""`,
			err:   nil,
		},
		{
			input: []byte(`"a" `),
			want:  `"a"`,
			err:   nil,
		},
		{
			input: []byte(`	"a"`),
			want: `"a"`,
			err:  nil,
		},
		{
			input: []byte(` "a "`),
			want:  `"a "`,
			err:   nil,
		},
		{
			input: []byte(``),
			want:  `""`,
			err:   ErrInvalidNextToken,
		},
		{
			input: []byte(`"`),
			want:  `""`,
			err:   ErrInvalidNextToken,
		},
		{
			input: []byte(`a`),
			want:  `""`,
			err:   ErrInvalidNextToken,
		},
		{
			input: []byte(`"a`),
			want:  `""`,
			err:   ErrInvalidNextToken,
		},
		{
			input: []byte(`a"`),
			want:  `""`,
			err:   ErrInvalidNextToken,
		},
	}

	for i, tt := range tests {
		p := NewParser(tt.input).Init()
		v, err := p.ParseString()
		if err != tt.err {
			t.Fatalf("[case %d] unexpected err: want=nil, got=%s", i, err)
		}

		got := v.String()
		if got != tt.want {
			t.Fatalf("[case %d] unexpected value: want=%#v, got=%#v", i, tt.want, got)
		}
	}
}
