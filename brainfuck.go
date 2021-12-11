package brainfuck

import (
	"bufio"
	"fmt"
	"os"
)

var PrintOutput = false
var MemorySize = 30_000

func Execute(input []byte) (output []byte, err error) {
	// brainfuck code only
	program := make([]byte, 0, len(input))
	//  double lookup to match loop positions faster
	loop := make(map[int]int)
	// for keeping track of brackets during parsing
	stack := make([]int, 0)
	for i, b := range input {
		switch b {
		//    +	  ,   -   .   <   >
		case 43, 44, 45, 46, 60, 62:
			program = append(program, b)
		case 91: // [
			// push bracket position in program
			stack = append(stack, len(program))
			program = append(program, b)
		case 93: // ]
			if len(stack) == 0 {
				return nil, fmt.Errorf("unmatched closing bracket at position: %d", i)
			}
			n := len(stack) - 1
			// double lookup
			loop[stack[n]] = len(program) // start -> end
			loop[len(program)] = stack[n] // end -> start
			stack = stack[:n]             // pop bracket position
			program = append(program, b)
		}
	}
	// check that the stack is empty
	if len(stack) > 0 {
		return nil, fmt.Errorf("unmatched opening bracket at position: %d", stack[0])
	}

	// initialize main program memory and pointer
	ptr := 0
	memory := make([]byte, MemorySize)

	// main program loop
	for i := 0; i < len(program); i++ {
		switch program[i] {
		case 43: // +
			memory[ptr]++
		case 45: // -
			memory[ptr]--
		case 60: // <
			ptr--
		case 62: // >
			ptr++
		case 44: // ,
			reader := bufio.NewReader(os.Stdin)
			if memory[ptr], err = reader.ReadByte(); err != nil {
				return nil, err
			}
		case 46: // .
			if PrintOutput {
				fmt.Print(string(memory[ptr]))
			}
			output = append(output, memory[ptr])
		case 91: // [
			if memory[ptr] == 0 {
				i = loop[i] // start -> end
			}
		case 93: // ]
			if memory[ptr] != 0 {
				i = loop[i] // end -> start
			}
		}
	}
	return output, nil
}
