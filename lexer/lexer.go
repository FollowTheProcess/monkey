// Package lexer implements the Monkey semantic Tokens and Lexer
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

// peekChar returns the character at l.readPosition
// i.e. it lets us look ahead by 1 character so we can detect
// things like == and !=
func (l *Lexer) peekChar() rune {
	if l.readPosition >= l.input.RuneCount() {
		return 0
	}

	return l.input.At(l.readPosition)
}

// NextToken looks at the current character under examination, emits the appropriate Token
// and reads the next character (thus advancing the indexes)
// If 0 is found, will emit an EOF
func (l *Lexer) NextToken() Token {
	var token Token
	l.skipWhiteSpace()

	switch l.ch {
	case '=':
		// Look ahead to see if we have a '=='
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			token = Token{Type: EQ, Literal: string(ch) + string(l.ch)}
		} else {
			// If not, must just be a normal '='
			token = newToken(ASSIGN, l.ch)
		}
	case '+':
		token = newToken(PLUS, l.ch)
	case '-':
		token = newToken(MINUS, l.ch)
	case '!':
		// Look ahead to see if we have a '!='
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			token = Token{Type: NOTEQ, Literal: string(ch) + string(l.ch)}
		} else {
			// If not, must just be a normal '!'
			token = newToken(BANG, l.ch)
		}
	case '*':
		token = newToken(ASTERISK, l.ch)
	case '/':
		token = newToken(SLASH, l.ch)
	case '<':
		token = newToken(LT, l.ch)
	case '>':
		token = newToken(GT, l.ch)
	case ';':
		token = newToken(SEMICOLON, l.ch)
	case ',':
		token = newToken(COMMA, l.ch)
	case '(':
		token = newToken(LPAREN, l.ch)
	case ')':
		token = newToken(RPAREN, l.ch)
	case '{':
		token = newToken(LBRACE, l.ch)
	case '}':
		token = newToken(RBRACE, l.ch)
	case '"':
		token.Type = STRING
		token.Literal = l.readString()
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

// readNumber reads l.ch so long as it a valid digit
// and advances the index until it reaches a non-digit character
// upon which it will return the string of valid digits
func (l *Lexer) readNumber() string {
	position := l.position
	for unicode.IsDigit(l.ch) {
		l.readChar()
	}

	return l.input.Slice(position, l.position)
}

func (l *Lexer) readString() string {
	position := l.position + 1
	for {
		l.readChar()
		if l.ch == '"' || l.ch == 0 {
			break
		}
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
