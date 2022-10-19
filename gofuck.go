package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("usage: gofuck <filename>")
	}

	code, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatalf("failed to open file")
	}

	loops := make([]int, 0, 1024)
	tape := make([]byte, 30000)
	ptr := 0
	skip := 0

	for pos := 0; pos < len(code); pos++ {

		switch code[pos] {
		case '+':
			tape[ptr]++
		case '-':
			tape[ptr]--
		case '.':
			fmt.Printf("%c", tape[ptr])
		case ',':
			fmt.Printf("Enter a single Character: ")
			reader := bufio.NewReader(os.Stdin)
			char, _, err := reader.ReadRune()

			if err != nil {
				fmt.Println(err)
			}

			tape[ptr] = uint8(char)

		case '<':
			ptr--
		case '>':
			ptr++
		case '[':
			if tape[ptr] == 0 {
				skip++
				for skip > 0 {
					pos++
					if code[pos] == '[' {
						skip++
					} else if code[pos] == ']' {
						skip--
					}
				}
			} else {
				loops = append(loops, pos)
			}
		case ']':
			if tape[ptr] == 0 {
				loops = loops[:len(loops)-1]
			} else {
				pos = loops[len(loops)-1]
			}
		}
	}
}
