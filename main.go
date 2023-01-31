package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/aitumik/interpreter/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! this is an interpreter\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
