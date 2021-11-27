package lexer

type Lexer struct {
	input        string
	position     int  // Current position in input (points to current char)
	readPosition int  // Current reading position (points to next char)
	ch           byte // Current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
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

func newToken(t TokenType, ch byte) Token {
	return Token{Type: t, Literal: string(ch)}
}
