package main

import (
	"fmt"
	"os"

	"github.com/kaneta1992/writing-an-interpreter-in-go/repl"
)

func main() {
	fmt.Printf("Hello! This is the Monkey programming language!\n")
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
