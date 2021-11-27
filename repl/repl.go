// Package repl implements the Monkey REPL (Read Eval Print Loop)
package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/FollowTheProcess/monkey/lexer"
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

		for token := l.NextToken(); token.Type != lexer.EOF; token = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", token)
		}
	}
}
