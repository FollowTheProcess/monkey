package vm

import (
	"fmt"
	"testing"

	"github.com/FollowTheProcess/monkey/ast"
	"github.com/FollowTheProcess/monkey/compiler"
	"github.com/FollowTheProcess/monkey/lexer"
	"github.com/FollowTheProcess/monkey/object"
	"github.com/FollowTheProcess/monkey/parser"
)

type vmTestCase struct {
	input    string
	expected interface{}
}

func TestIntegerArithmetic(t *testing.T) {
	tests := []vmTestCase{
		{"1", 1},
		{"2", 2},
		{"1 + 2", 3},
	}

	runVmTests(t, tests)
}

func runVmTests(t *testing.T, tests []vmTestCase) {
	t.Helper()

	for _, tt := range tests {
		program := parse(tt.input)

		comp := compiler.New()
		err := comp.Compile(program)
		if err != nil {
			t.Fatalf("compiler error: %s", err)
		}

		vm := New(comp.ByteCode())
		err = vm.Run()
		if err != nil {
			t.Fatalf("vm error: %s", err)
		}

		stackElem := vm.StackTop()

		testExpectedObject(t, tt.expected, stackElem)
	}
}

func parse(input string) *ast.Program {
	l := lexer.New(input)
	p := parser.New(l)
	return p.ParseProgram()
}

func testIntegerObject(expected int, actual object.Object) error {
	result, ok := actual.(*object.Integer)
	if !ok {
		return fmt.Errorf("object is not an Integer, got %T (%+v)", actual, actual)
	}

	if result.Value != expected {
		return fmt.Errorf("object has wrong value: got %d, wanted %d", result.Value, expected)
	}

	return nil
}

func testExpectedObject(t *testing.T, expected interface{}, actual object.Object) {
	t.Helper()

	switch expected := expected.(type) {
	case int:
		err := testIntegerObject(expected, actual)
		if err != nil {
			t.Errorf("testIntegerObject failed: %s", err)
		}
	}
}
