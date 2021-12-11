package main

import (
	"fmt"
	"os"

	"github.com/phrozen/brainfuck"
)

func main() {
	if len(os.Args) < 2 {
		panic("You must specify and input file")
	}

	input, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	output, err := brainfuck.Execute(input)
	if err != nil {
		panic(err)
	}

	fmt.Print(string(output))
}
