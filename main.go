package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/FollowTheProcess/monkey/repl"
)

var (
	version = "dev"
	commit  = ""
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the Monkey programming language!\n", user.Username)
	fmt.Printf("Version: %s\n", version)
	fmt.Printf("Commit: %s\n", commit)
	fmt.Println("Type some commands...")

	repl.Start(os.Stdin, os.Stdout)
}
