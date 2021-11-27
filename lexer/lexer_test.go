package lexer

import "testing"

func TestNextToken(t *testing.T) {
	input := `=+(){},;`

	tests := []struct {
		name            string
		expectedType    TokenType
		expectedLiteral string
	}{
		{
			name:            "assignment",
			expectedType:    ASSIGN,
			expectedLiteral: "=",
		},
		{
			name:            "plus",
			expectedType:    PLUS,
			expectedLiteral: "+",
		},
		{
			name:            "left parenthesis",
			expectedType:    LPAREN,
			expectedLiteral: "(",
		},
		{
			name:            "right parenthesis",
			expectedType:    RPAREN,
			expectedLiteral: ")",
		},
		{
			name:            "comma",
			expectedType:    COMMA,
			expectedLiteral: ",",
		},
		{
			name:            "semicolon",
			expectedType:    SEMICOLON,
			expectedLiteral: ";",
		},
		{
			name:            "end of file",
			expectedType:    EOF,
			expectedLiteral: "",
		},
	}

	l := New(input)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			token := l.NextToken()

			if token.Type != tt.expectedType {
				t.Fatalf("wrong token type, got %q, wanted %q", token.Type, tt.expectedType)
			}

			if token.Literal != tt.expectedLiteral {
				t.Fatalf("wrong token literal, got %q, wanted %q", token.Literal, tt.expectedLiteral)
			}
		})
	}
}
