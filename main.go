package main

import (
	"bufio"
	"fmt"
	"os"

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

		fmt.Printf("%s\n\n", line)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	getUserInput(reader)
}
