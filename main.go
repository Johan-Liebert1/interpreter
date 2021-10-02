package main

import (
	"bufio"
	"fmt"
	"os"

	"interpreter/interpreter"
	"interpreter/types"
)

func parsedInput(input string) {
	for char := range input {
		switch char {
		}
	}
}

func getUserInput(reader *bufio.Reader) {
	for {
		fmt.Printf(">>> ")

		line, _, _ := reader.ReadLine()

		userInput := string(line)

		if _, ok := types.ExitCodes[userInput]; ok {
			os.Exit(0)
		}

		fmt.Printf("%s\n\n", line)

		lexer := interpreter.Interpreter{}
		lexer.Init(userInput)

		fmt.Println(lexer)

	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	getUserInput(reader)
}
