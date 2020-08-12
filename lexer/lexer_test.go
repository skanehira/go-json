package lexer

import (
	"testing"

	"github.com/skanehira/go-json/token"
)

func TestNextToken(t *testing.T) {
	input := `{"types": [true, false, null, "string", 11, 55.55]}`

	tests := []struct {
		wantType    token.TokenType
		wantLiteral string
	}{
		{token.LBRACE, "{"},
		{token.LETTER, "types"},
		{token.COLON, ":"},
		{token.LBRACKET, "["},
		{token.TRUE, "true"},
		{token.COMMA, ","},
		{token.FALSE, "false"},
		{token.COMMA, ","},
		{token.NULL, "null"},
		{token.COMMA, ","},
		{token.LETTER, "string"},
		{token.COMMA, ","},
		{token.INT, "11"},
		{token.COMMA, ","},
		{token.FLOAT, "55.55"},
		{token.RBRACKET, "]"},
		{token.RBRACE, "}"},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.wantType {
			t.Fatalf("tests[%d] - token type wrong. want=%q, got=%q", i, tt.wantType, tok.Type)
		}
		if tok.Literal != tt.wantLiteral {
			t.Fatalf("tests[%d] - token literal wrong. want=%q, got=%q", i, tt.wantLiteral, tok.Literal)
		}
	}
}
