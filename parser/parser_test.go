package parser

import (
	"fmt"
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

func TestReturnStatements(t *testing.T) {
	input := `
	return 5;
	return 10;
	return 993322;
	`

	lex := lexer.New(input)
	p := New(lex)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements, got %d", len(program.Statements))
	}

	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("stmt not *ast.ReturnStatement, got %T", stmt)
		}
		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("returnStmt.TokenLiteral not 'return', got %q", returnStmt.TokenLiteral())
		}
	}
}

func TestIdentifierExpression(t *testing.T) {
	input := "foobar;"

	lex := lexer.New(input)
	p := New(lex)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("wrong number of statements, got %d, wanted %d", len(program.Statements), 1)
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("statement is not an ast.ExpressionStatement, got %T", program.Statements[0])
	}

	ident, ok := stmt.Expression.(*ast.Identifier)
	if !ok {
		t.Fatalf("expression not *ast.Identifier, got %T", stmt.Expression)
	}

	if ident.Value != "foobar" {
		t.Errorf("wrong ident.Value, got %s, wanted %s", ident.Value, "foobar")
	}

	if ident.TokenLiteral() != "foobar" {
		t.Errorf("wrong TokenLiteral, got %s, wanted %s", ident.TokenLiteral(), "foobar")
	}
}

func TestIntegerLiteralExpression(t *testing.T) {
	input := "5;"

	lex := lexer.New(input)
	p := New(lex)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("wrong number of statements, got %d, wanted %d", len(program.Statements), 1)
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("statement not an ExpressionStatement, got %T", stmt)
	}

	literal, ok := stmt.Expression.(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("expression not IntegerLiteral, got %T", stmt.Expression)
	}

	if literal.Value != 5 {
		t.Errorf("wrong literal.Value, got %d, wanted %d", literal.Value, 5)
	}

	if literal.TokenLiteral() != "5" {
		t.Errorf("wrong TokenLiteral, got %s, wanted %s", literal.TokenLiteral(), "5")
	}
}

func TestParsingPrefixExpressions(t *testing.T) {
	tests := []struct {
		input    string
		operator string
		value    int
	}{
		{"!5;", "!", 5},
		{"-15;", "-", 15},
	}

	for _, tt := range tests {
		lex := lexer.New(tt.input)
		p := New(lex)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("wrong number of statements, got %d, wanted %d", len(program.Statements), 1)
		}

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("statement is not an ExpressionStatement, got %T", program.Statements[0])
		}

		exp, ok := stmt.Expression.(*ast.PrefixExpression)
		if !ok {
			t.Fatalf("statement is not a PrefixExpression, got %T", stmt.Expression)
		}

		if exp.Operator != tt.operator {
			t.Fatalf("wrong operator, got %s, wanted %s", exp.Operator, tt.operator)
		}

		assertIntegerLiteral(t, exp.Right, tt.value)

	}
}

func TestParsingInfixExpressions(t *testing.T) {
	tests := []struct {
		input    string
		left     int
		operator string
		right    int
	}{
		{"5 + 5;", 5, "+", 5},
		{"5 - 5;", 5, "-", 5},
		{"5 * 5;", 5, "*", 5},
		{"5 / 5;", 5, "/", 5},
		{"5 > 5;", 5, ">", 5},
		{"5 < 5;", 5, "<", 5},
		{"5 == 5;", 5, "==", 5},
		{"5 != 5;", 5, "!=", 5},
	}

	for _, tt := range tests {
		lex := lexer.New(tt.input)
		p := New(lex)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("wrong number of statements, got %d, wanted %d", len(program.Statements), 1)
		}

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("statement is not ExpressionStatement, got %T", program.Statements[0])
		}

		exp, ok := stmt.Expression.(*ast.InfixExpression)
		if !ok {
			t.Fatalf("expression is not InfixExpression, got %T", stmt.Expression)
		}

		assertIntegerLiteral(t, exp.Left, tt.left)

		if exp.Operator != tt.operator {
			t.Fatalf("wrong operator, got %s, wanted %s", exp.Operator, tt.operator)
		}

		assertIntegerLiteral(t, exp.Right, tt.right)
	}
}

func TestOperatorPrecedenceParsing(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			"-a * b",
			"((-a) * b)",
		},
		{
			"!-a",
			"(!(-a))",
		},
		{
			"a + b + c",
			"((a + b) + c)",
		},
		{
			"a + b - c",
			"((a + b) - c)",
		},
		{
			"a * b * c",
			"((a * b) * c)",
		},
		{
			"a * b / c",
			"((a * b) / c)",
		},
		{
			"a + b / c",
			"(a + (b / c))",
		},
		{
			"a + b * c + d / e - f",
			"(((a + (b * c)) + (d / e)) - f)",
		},
		{
			"3 + 4; -5 * 5",
			"(3 + 4)((-5) * 5)",
		},
		{
			"5 > 4 == 3 < 4",
			"((5 > 4) == (3 < 4))",
		},
		{
			"5 < 4 != 3 > 4",
			"((5 < 4) != (3 > 4))",
		},
		{
			"3 + 4 * 5 == 3 * 1 + 4 * 5",
			"((3 + (4 * 5)) == ((3 * 1) + (4 * 5)))",
		},
		{
			"3 + 4 * 5 == 3 * 1 + 4 * 5",
			"((3 + (4 * 5)) == ((3 * 1) + (4 * 5)))",
		},
	}

	for _, tt := range tests {
		lex := lexer.New(tt.input)
		p := New(lex)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		actual := program.String()

		if actual != tt.expected {
			t.Errorf("got %s, wanted %s", actual, tt.expected)
		}
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

func assertIntegerLiteral(t *testing.T, il ast.Expression, value int) {
	t.Helper()

	i, ok := il.(*ast.IntegerLiteral)
	if !ok {
		t.Errorf("not an *ast.IntegerLiteral, got %T", il)
	}

	if i.Value != value {
		t.Errorf("wrong integer value, got %d, wanted %d", i.Value, value)
	}

	if i.TokenLiteral() != fmt.Sprintf("%d", value) {
		t.Errorf("wrong TokenLiteral, got %s, wanted %s", i.TokenLiteral(), fmt.Sprintf("%d", value))
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
