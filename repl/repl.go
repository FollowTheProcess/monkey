// Package repl implements the Monkey REPL (Read Eval Print Loop)
package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/FollowTheProcess/monkey/eval"
	"github.com/FollowTheProcess/monkey/lexer"
	"github.com/FollowTheProcess/monkey/object"
	"github.com/FollowTheProcess/monkey/parser"
)

const PROMPT = ">> "

// Start will start a REPL, currently only exited with ctrl + c
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

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

		evaluated := eval.Eval(program, env)
		if evaluated != nil {
			fmt.Fprintln(out, evaluated.Inspect())
		}
	}
}

func printParseErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		fmt.Fprintf(out, "\t%s\n", msg)
	}
}
