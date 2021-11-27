package lexer

import "golang.org/x/exp/utf8string"

// Lexer is our semantic Lexer, it proceeds character by character emitting
// semantic Tokens as it sees them
// It is identical to the one in the book other than I have tried to make it
// support UTF-8
type Lexer struct {
	input        *utf8string.String // So we can support UTF-8
	position     int                // Current position in input (points to current char)
	readPosition int                // Current reading position (points to next char)
	ch           rune               // Current char under examination
}

// New constructs and returns a new Lexer and initialises
// it by reading the first character
func New(input string) *Lexer {
	l := &Lexer{input: utf8string.NewString(input)}
	l.readChar()
	return l
}

// readChar
func (l *Lexer) readChar() {
	if l.readPosition >= l.input.RuneCount() {
		l.ch = 0
	} else {
		l.ch = l.input.At(l.readPosition)
	}

	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() Token {
	var token Token

	switch l.ch {
	case '=':
		token = newToken(ASSIGN, l.ch)
	case ';':
		token = newToken(SEMICOLON, l.ch)
	case '(':
		token = newToken(LPAREN, l.ch)
	case ')':
		token = newToken(RPAREN, l.ch)
	case ',':
		token = newToken(COMMA, l.ch)
	case '+':
		token = newToken(PLUS, l.ch)
	case '{':
		token = newToken(LBRACE, l.ch)
	case '}':
		token = newToken(RBRACE, l.ch)
	case 0:
		token.Literal = ""
		token.Type = EOF
	}

	l.readChar()
	return token
}

func newToken(t TokenType, ch rune) Token {
	return Token{Type: t, Literal: string(ch)}
}
