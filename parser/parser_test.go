package parser

import (
	"testing"

	"github.com/skanehira/go-json/lexer"
)

func TestParseObject(t *testing.T) {
	tests := []struct {
		want string
	}{
		{`{}`},
		{`{"name":"gorilla"}`},
		{`{"name":"gorilla","age":27,"like":"banan"}`},
		{`{"banana":100}`},
		{`{"likes":["cat","banana","dog"]}`},
	}

	for _, tt := range tests {
		l := lexer.New(tt.want)
		p := New(l)

		o := p.parseObject()

		if o.String() != tt.want {
			t.Fatalf("wront object ast. want=%s, got=%s", tt.want, o.String())
		}
	}
}

func TestParseArray(t *testing.T) {
	tests := []struct {
		want string
	}{
		{`[]`},
		{`["string",null,true,false,10]`},
		{`[{"name":"gorilla"},["string",true,false,10],10,true,false,10,"string"]`},
		{`[{"name":"gorilla"},[[]]]`},
	}

	for _, tt := range tests {
		l := lexer.New(tt.want)
		p := New(l)

		a := p.parseArray()

		if a.String() != tt.want {
			t.Fatalf("wront object ast. want=%s, got=%s", tt.want, a.String())
		}
	}
}
