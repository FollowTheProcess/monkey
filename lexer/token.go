package lexer

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers and literals
	IDENT = "IDENT"
	INT   = "INT"

	// Operators
	ASSIGN = "="
	PLUS   = "+"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

type TokenType string

// Token represents a lexical token for monkey
type Token struct {
	Type    TokenType // The type of token
	Literal string    // The token's literal string value
}

// keywords are the monkey keywords mapped to their Tokens
var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

// LookupIdent checks to see if 'ident' is an accepted keyword
// returning it's token if it is
// If not, it will return a generic IDENT which can be any user-defined identifier
func LookupIdent(ident string) TokenType {
	if token, ok := keywords[ident]; ok {
		return token
	}
	return IDENT
}
