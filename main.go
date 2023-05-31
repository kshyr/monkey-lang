package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/kshyr/monkey-lang/repl"
)

func main() {
	_, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("monkey!\n")
	repl.Start(os.Stdin, os.Stdout)
}
