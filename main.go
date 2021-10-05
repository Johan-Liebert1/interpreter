package main

import (
	"bufio"
	"fmt"
	"os"

	"interpreter/interpreter"
	"interpreter/types"
)

func getUserInput(reader *bufio.Reader) {
	for {
		fmt.Printf(">>> ")

		line, _, _ := reader.ReadLine()

		userInput := string(line)

		if _, ok := types.ExitCodes[userInput]; ok {
			os.Exit(0)
		}

		// fmt.Printf("%s\n\n", line)

		langInterpreter := interpreter.Interpreter{}
		langInterpreter.Init(userInput)
		result := langInterpreter.Interpret()

		fmt.Println(result)

	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	getUserInput(reader)
}
