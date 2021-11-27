package lexer

import "testing"

func TestNextToken(t *testing.T) {
	input := `=+(){},;`

	tests := []struct {
		expectedType    TokenType
		expectedLiteral string
	}{
		{ASSIGN, "="},
		{PLUS, "+"},
		{LPAREN, "("},
		{RPAREN, ")"},
		{LBRACE, "{"},
		{RBRACE, "}"},
		{COMMA, ","},
		{SEMICOLON, ";"},
		{EOF, ""},
	}

	l := New(input)

	for _, tt := range tests {
		t.Run(string(tt.expectedType), func(t *testing.T) {
			tok := l.NextToken()

			if tok.Type != tt.expectedType {
				t.Fatalf("wrong token type: got %q, wanted %q", tok.Type, tt.expectedType)
			}

			if tok.Literal != tt.expectedLiteral {
				t.Fatalf("wrong token literal: got %q, wanted %q", tok.Literal, tt.expectedLiteral)
			}
		})
	}
}
