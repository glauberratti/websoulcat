package input

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var READER = bufio.NewReader(os.Stdin)

func readInput() (string, error) {
	inputStr := ""
	inputStr, err := READER.ReadString('\n')
	if err != nil {
		log.Println("Error to read input (input.readInput()):", err)
		return "", err
	}

	inputStr = strings.TrimSuffix(inputStr, "\n")
	inputStr = strings.TrimSuffix(inputStr, "\r")
	return inputStr, nil
}

func ReadInputWithQuestion(question string) (string, error) {
	inputStr := ""
	var err error = nil

	for len(inputStr) == 0 || err != nil {
		if len(question) > 0 {
			fmt.Print(question)
		}

		inputStr, err = readInput()
		if err != nil {
			return "", err
		}
		if len(inputStr) == 0 {
			fmt.Println("")
			fmt.Println("Valor não pode ser vazio")
		}
	}

	return inputStr, err
}
