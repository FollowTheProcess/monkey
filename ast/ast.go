// Package ast implements the Abstract Syntax Tree for Monkey
package ast

import (
	"bytes"

	"github.com/FollowTheProcess/monkey/lexer"
)

// Node represents a generic AST node
type Node interface {
	TokenLiteral() string
	String() string
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

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
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

func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}

	out.WriteString(";")

	return out.String()
}

// Identifier refers to a variable assignment
// e.g. 'x = 5;'
type Identifier struct {
	Token lexer.Token // The 'IDENT' token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

func (i *Identifier) String() string {
	return i.Value
}

// ReturnStatement is our object responsible for e.g. 'return true'
type ReturnStatement struct {
	Token       lexer.Token // The 'RETURN' token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")

	return out.String()
}

// ExpressionStatement is our object responsible for e.g. 'x = 5;'
// being a valid statement
type ExpressionStatement struct {
	Token      lexer.Token // The first token of the expression
	Expression Expression
}

func (es *ExpressionStatement) statementNode()       {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }

func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

// IntegerLiteral is out object responsible for e.g. '5;'
type IntegerLiteral struct {
	Token lexer.Token
	Value int
}

func (il *IntegerLiteral) expressionNode()      {}
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }
func (il *IntegerLiteral) String() string       { return il.Token.Literal }

// PrefixExpression is our object responsible for e.g. '!true;'
type PrefixExpression struct {
	Token    lexer.Token // The prefix token e.g. '!'
	Operator string
	Right    Expression
}

func (pe *PrefixExpression) expressionNode()      {}
func (pe *PrefixExpression) TokenLiteral() string { return pe.Token.Literal }

func (pe *PrefixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")

	return out.String()
}

// InfixExpression is our object responsible for e.g. '5 != 10;'
type InfixExpression struct {
	Token    lexer.Token // The operator token, e.g. '+'
	Left     Expression
	Operator string
	Right    Expression
}

func (ie *InfixExpression) expressionNode()      {}
func (ie *InfixExpression) TokenLiteral() string { return ie.Token.Literal }

func (ie *InfixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString(" " + ie.Operator + " ")
	out.WriteString(ie.Right.String())
	out.WriteString(")")

	return out.String()
}
