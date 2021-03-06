package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"

	"programminglang/constants"
	"programminglang/helpers"
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

		langInterpreter.Init(userInput, false)
		result := langInterpreter.Interpret()

		if result != nil {
			helpers.ColorPrint(constants.LightYellow, 1, 1, result)
		}
	}
}

func interpretFile(langInterpreter interpreter.Interpreter, fileName string) {
	file, err := os.Open(fileName)

	if err != nil {
		fmt.Printf("File '%s' does not exist in the current directory\n", fileName)
		os.Exit(1)
	}

	fileData, err := ioutil.ReadAll(file)

	if err != nil {
		fmt.Printf("Failed to read file '%s'\n", fileName)
		os.Exit(1)
	}

	langInterpreter.Init(string(fileData), false)

	result := langInterpreter.Interpret()

	if result != nil {
		helpers.ColorPrint(constants.LightYellow, 1, 1, result)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	langInterpreter := interpreter.Interpreter{}
	langInterpreter.InitConcrete()

	args := os.Args

	if len(args) == 1 {
		getUserInput(reader, langInterpreter)
	} else {
		interpretFile(langInterpreter, args[1])
	}
}
