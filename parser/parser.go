// Package parser implements the recursive descent, 'Pratt' parser for
// the Monkey programming language
package parser

import (
	"fmt"

	"github.com/FollowTheProcess/monkey/ast"
	"github.com/FollowTheProcess/monkey/lexer"
)

// Parser is our parser, it contains a lexer and fields to track
// the current and next tokens as emitted from the lexer
type Parser struct {
	l            *lexer.Lexer
	currentToken lexer.Token
	peekToken    lexer.Token
	errors       []string
}

// New constructs and returns a new parser
func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}

	// Read two tokens so currentToken and peekToken are both set
	p.nextToken()
	p.nextToken()

	return p
}

// Errors returns a list of all parsing errors
// no errors will return an empty slice
func (p *Parser) Errors() []string {
	return p.errors
}

// nextToken sets the internal token state of the parser
// to reflext the next tokenised state as emitted from the lexer
func (p *Parser) nextToken() {
	// Set current token to the last token emitted from the lexer
	p.currentToken = p.peekToken
	// Get a new token from the lexer and set the new peek to that
	p.peekToken = p.l.NextToken()
}

// ParseProgram is the entry point to the parser
func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	// Parse the current token and keep getting new tokens
	// until we reach an EOF
	for !p.currentToken.Is(lexer.EOF) {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	return program
}

// parseStatement is a helper that will parse the current statement
// depending on what type of statement it is
func (p *Parser) parseStatement() ast.Statement {
	switch p.currentToken.Type {
	case lexer.LET:
		return p.parseLetStatement()
	case lexer.RETURN:
		return p.parseReturnStatement()
	default:
		return nil
	}
}

// parseLetStatement should be called on the 'let' branch of the main parser switch statement
// It will return nil if the statement is not of the correct form
// e.g. 'let x = 5;'
func (p *Parser) parseLetStatement() *ast.LetStatement {
	// Construct the LetStatement AST node with the token we're currently sitting on
	stmt := &ast.LetStatement{Token: p.currentToken}

	// First we expect the next token to be an IDENT
	// e.g. 'x'
	if !p.expectPeek(lexer.IDENT) {
		return nil
	}

	// Now we know we have an Identifier, construct the node
	stmt.Name = &ast.Identifier{Token: p.currentToken, Value: p.currentToken.Literal}

	// The next thing to expect is an ASSIGN '='
	if !p.expectPeek(lexer.ASSIGN) {
		return nil
	}

	// TODO: We're skipping the expressions until we encounter a semicolon
	// this will get fixed once we know how to parse expressions
	for !p.currentToken.Is(lexer.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

// parseReturnStatement handles e.g. 'return true'
func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.currentToken}

	p.nextToken()

	// TODO: Skip expressions until we encounter a semicolon
	for !p.currentToken.Is(lexer.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

// expectPeek checks if the peek (next) token is of type 't'
// if it is, it will advance the token and return true
// else return false
func (p *Parser) expectPeek(t lexer.TokenType) bool {
	if p.peekToken.Is(t) {
		p.nextToken()
		return true
	}

	p.peekError(t)
	return false
}

// peekError adds an error to the list of parser errors when the peeked
// token is not of the correct type
func (p *Parser) peekError(t lexer.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}
