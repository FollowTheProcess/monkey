package lexer

import (
	"unicode"

	"golang.org/x/exp/utf8string"
)

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

// readChar reads the next character in the input stream
// and advances our position markers
// If we have not read anything or we are at the end of the input
// it will set the current character to 0 (ASCII "NUL")
func (l *Lexer) readChar() {
	if l.readPosition >= l.input.RuneCount() {
		// Nothing to do, set to 0
		l.ch = 0
	} else {
		// Grab the character at the readPosition and store it
		// in l.ch
		l.ch = l.input.At(l.readPosition)
	}

	// Advance the indexes
	l.position = l.readPosition
	l.readPosition += 1
}

// NextToken looks at the current character under examination, emits the appropriate Token
// and reads the next character (thus advancing the indexes)
// If 0 is found, will emit an EOF
func (l *Lexer) NextToken() Token {
	var token Token
	l.skipWhiteSpace()

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
	default:
		switch {
		case unicode.IsLetter(l.ch):
			token.Literal = l.readIdentifier()
			token.Type = LookupIdent(token.Literal)
			// Early return as readIdentifier calls readChar repeatedly
			// so it does not need to be called again later
			return token

		case unicode.IsDigit(l.ch):
			token.Literal = l.readNumber()
			token.Type = INT
			// Another early return as readNumber will also repeatedly call
			// readChar
			return token

		default:
			token = newToken(ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return token
}

// readIdentifier reads l.ch so long as it is a valid utf-8 letter
// and advances the index until it reaches a non-letter character
// upon which it will return the string of valid letters i.e. the identifier
// it's just read
func (l *Lexer) readIdentifier() string {
	position := l.position
	for unicode.IsLetter(l.ch) {
		l.readChar()
	}

	return l.input.Slice(position, l.position)
}

func (l *Lexer) readNumber() string {
	position := l.position
	for unicode.IsDigit(l.ch) {
		l.readChar()
	}

	return l.input.Slice(position, l.position)
}

// skipWhiteSpace allows us to easily skip all whitespace characters
func (l *Lexer) skipWhiteSpace() {
	for unicode.IsSpace(l.ch) {
		l.readChar()
	}
}

// newToken constructs and returns a Token
func newToken(t TokenType, ch rune) Token {
	return Token{Type: t, Literal: string(ch)}
}
