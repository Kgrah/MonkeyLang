package lexer

import (
	"testing"

	"github.com/kgrah/MonkeyLang/token"
)

func TestNextToken(test *testing.T) {
	input := `=+(){},;`

	tests := []struct {
		expectType      token.TokenType
		expectedLiteral string
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

	l := NewLexer(input)
	for i, tt := range tests {
		token := l.NextToken()

		if token.Type != tt.expectType {
			test.Fatalf("test[%d] - tokentype wrong. Expected=%q, got=%q", i, tt.expectType, token.Type)
		}

		if token.Literal != tt.expectedLiteral {
			test.Fatalf("test[%d] - literal wrong. Expected=%q, got =%q", i, tt.expectedLiteral, token.Literal)
		}

	}

}
