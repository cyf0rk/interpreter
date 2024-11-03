package main

import (
	"os"

	"github.com/cyf0rk/interpreter/cmd/repl"
)

func main() {
	repl.Start(os.Stdin, os.Stdout)
}
