// Package ast implements the Abstract Syntax Tree for Monkey
package ast

import "github.com/FollowTheProcess/monkey/lexer"

// Node represents a generic AST node
type Node interface {
	TokenLiteral() string
}

// Statement represents a statement, i.e. a piece of syntax
// that does not produce a value
// e.g. 'let x = 5;'
type Statement interface {
	Node
	statementNode() // Dummy to ensure we only use this for statements
}

// Expression represents an expression, i.e. a piece of syntax
// that produces a value
// e.g. '5 + 5;'
type Expression interface {
	Node
	expressionNode() // Dummy to ensure we only use this for expressions
}

// Program is the root AST node for the entire program
type Program struct {
	Statements []Statement
}

// TokenLiteral satisfies the ast.Node interface for our Program
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

// LetStatement is our object responsible for e.g. 'let x = 5;'
type LetStatement struct {
	Token lexer.Token // The 'LET' token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

// Identifier refers to a variable assignment
// e.g. 'x = 5;'
type Identifier struct {
	Token lexer.Token // The 'IDENT' token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
