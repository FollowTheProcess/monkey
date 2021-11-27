package lexer

// Lexer is our semantic Lexer, it proceeds character by character
// emitting tokens as it sees them
type Lexer struct {
	input        string
	position     int  // Current position in input (points to current character)
	readPosition int  // Current reading position in input (points to next character)
	ch           byte // Current character under examination
}

// New constructs and returns a new Lexer
// the Lexer is initialised by reading the first rune in input
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// readChar reads the next character to be lexed or a 0 (ASCII "NUL")
// if we have reached the end of the input stream
//
// If we are not at the end of the stream, the current character is set
// and the current position and the position of the next character are advanced
func (l *Lexer) readChar() {
	// If we haven't read anything yet or if we've reached the end of the
	// input stream, set current char to 0
	// otherwise, index the input with current read position to grab the character
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	// Advance our current and next positions
	l.position = l.readPosition
	l.readPosition += 1
}

// NextToken looks at the current character under evaluation and
// returns the appropriate semantic Token for that character
// in the case of 0 (ASCII "NUL"), it will return an EOF Token
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

// newToken constructs and returns a Token
func newToken(t TokenType, ch byte) Token {
	return Token{Type: t, Literal: string(ch)}
}
