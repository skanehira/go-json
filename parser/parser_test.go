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
			t.Fatalf("[case %d] unexpected err: want=%s, got=%s", i, tt.err, err)
		}

		got := v.String()
		if got != tt.want {
			t.Fatalf("[case %d] unexpected value: want=%#v, got=%#v", i, tt.want, got)
		}
	}
}
