package lexer

import (
	"testing"

	"sisi/token"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;`

	tests := []struct {
		expectedType    token.TokenType
		exceptedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. excepted=%q, got=%q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.exceptedLiteral {
			t.Fatalf("tests[%d] - literal worng. excepted=%q, got=%q", i, tt.exceptedLiteral, tok.Literal)
		}
	}
}
