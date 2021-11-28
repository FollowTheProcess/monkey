package parser

import (
	"testing"

	"github.com/FollowTheProcess/monkey/ast"
	"github.com/FollowTheProcess/monkey/lexer"
)

func TestLetStatements(t *testing.T) {
	input := `
	let x = 5;
	let y = 10;
	let foobar = 838383;
	`

	lex := lexer.New(input)
	p := New(lex)

	program := p.ParseProgram()
	checkParserErrors(t, p)
	if program == nil {
		t.Fatal("ParseProgram() returned nil")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements, got %d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifer string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		assertLetStatement(t, stmt, tt.expectedIdentifer)
	}
}

func assertLetStatement(t *testing.T, stmt ast.Statement, expectedIdent string) {
	t.Helper()
	if stmt.TokenLiteral() != "let" {
		t.Errorf("Let statement TokenLiteral not 'let', got %q", stmt.TokenLiteral())
	}

	letStmt, ok := stmt.(*ast.LetStatement)
	if !ok {
		t.Errorf("statement is not an *ast.LetStatement, got %T", stmt)
	}

	if letStmt.Name.Value != expectedIdent {
		t.Errorf("letStmt.Name.Value not %q, got %q", expectedIdent, letStmt.Name.Value)
	}

	if letStmt.Name.TokenLiteral() != expectedIdent {
		t.Errorf("statement name not %q, got %q", expectedIdent, letStmt.Name.TokenLiteral())
	}
}

func checkParserErrors(t *testing.T, p *Parser) {
	t.Helper()
	errors := p.Errors()

	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}
