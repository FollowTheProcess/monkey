package vm

import (
	"fmt"

	"github.com/FollowTheProcess/monkey/code"
	"github.com/FollowTheProcess/monkey/compiler"
	"github.com/FollowTheProcess/monkey/object"
)

const StackSize = 2048

type VM struct {
	constants    []object.Object
	instructions code.Instructions

	// The canonical stack for the VM
	stack []object.Object
	// The stack pointer, always points to the next value
	// top of stack is stack[sp - 1]
	sp int
}

func New(bytecode *compiler.ByteCode) *VM {
	return &VM{
		instructions: bytecode.Instructions,
		constants:    bytecode.Constants,

		stack: make([]object.Object, StackSize),
		sp:    0,
	}
}

func (vm *VM) StackTop() object.Object {
	if vm.sp == 0 {
		return nil
	}
	return vm.stack[vm.sp-1]
}

func (vm *VM) Run() error {
	for i := 0; i < len(vm.instructions); i++ {
		op := code.Opcode(vm.instructions[i])

		switch op {
		case code.OpConstant:
			constIndex := code.ReadUint16(vm.instructions[i+1:])
			i += 2
			err := vm.push(vm.constants[constIndex])
			if err != nil {
				return err
			}

		case code.OpAdd:
			// Note: this assumes the right hand value was the last one
			// to be pushed onto the stack
			right := vm.pop()
			left := vm.pop()
			leftValue := left.(*object.Integer).Value
			rightValue := right.(*object.Integer).Value

			result := leftValue + rightValue
			if err := vm.push(&object.Integer{Value: result}); err != nil {
				return err
			}
		}
	}

	return nil
}

func (vm *VM) push(obj object.Object) error {
	if vm.sp >= StackSize {
		return fmt.Errorf("STACK OVERFLOW!!")
	}

	vm.stack[vm.sp] = obj
	vm.sp++

	return nil
}

func (vm *VM) pop() object.Object {
	o := vm.stack[vm.sp-1]
	vm.sp--
	return o
}
