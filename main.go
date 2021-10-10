package main

import (
	"bufio"
	"fmt"
	"os"

	"programminglang/interpreter"
	"programminglang/types"
)

func getUserInput(reader *bufio.Reader, langInterpreter interpreter.Interpreter) {
	for {
		fmt.Printf(">>> ")

		line, _, _ := reader.ReadLine()

		userInput := string(line)

		if _, ok := types.ExitCodes[userInput]; ok {
			os.Exit(0)
		}

		if len(userInput) == 0 {
			continue
		}

		langInterpreter.Init(userInput)
		result := langInterpreter.Interpret()

		fmt.Println(result)

	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	langInterpreter := interpreter.Interpreter{}
	langInterpreter.InitConcrete()

	getUserInput(reader, langInterpreter)
}
