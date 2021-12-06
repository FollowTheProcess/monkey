// Package repl implements the Monkey REPL (Read Eval Print Loop)
package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/FollowTheProcess/monkey/compiler"
	"github.com/FollowTheProcess/monkey/lexer"
	"github.com/FollowTheProcess/monkey/parser"
	"github.com/FollowTheProcess/monkey/vm"
)

const PROMPT = ">> "

// Start will start a REPL, currently only exited with ctrl + c
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprint(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParseErrors(out, p.Errors())
			continue
		}

		// Here's the new bit, add the compiler into the mix!
		comp := compiler.New()
		err := comp.Compile(program)
		if err != nil {
			fmt.Fprintf(out, "Uh oh! This doesn't compile:\n %s\n", err)
			continue
		}

		machine := vm.New(comp.ByteCode())
		err = machine.Run()
		if err != nil {
			fmt.Fprintf(out, "Uh oh! Can't execute the bytecode:\n %s\n", err)
			continue
		}

		stackTop := machine.StackTop()
		fmt.Fprintln(out, stackTop.Inspect())
	}
}

func printParseErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		fmt.Fprintf(out, "\t%s\n", msg)
	}
}
