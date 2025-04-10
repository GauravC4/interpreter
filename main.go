package main

import (
	"fmt"
	"os"

	"github.com/GauravC4/interpreter/repl"
)

func main() {
	fmt.Printf("Hello! This is a simple interpreter!\n")
	fmt.Printf("Type your commands here\n")
	repl.Start(os.Stdin, os.Stdout)
}
