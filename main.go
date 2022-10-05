package main

import (
	"bufio"
	"fmt"
	"os"
	"return-change-machine/config"
	"return-change-machine/logic"
	"strconv"
	"strings"
)

func readInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	text, _ := reader.ReadString('\n')

	// convert CRLF to LF
	text = strings.Replace(text, "\n", "", -1)
	text = strings.Replace(text, "\r", "", -1)

	return text
}

func main() {
	change := logic.InitChange()

	var input string

	if config.DEBUG_MODE {
		input = "235101"
	} else {
		input = readInput("Enter text (introduce the value in cents please): ")
	}

	// parse amount
	parsedAmount, err := strconv.Atoi(input)
	if config.DEBUG_MODE {
		fmt.Println("raw", input, "parsed", parsedAmount, "error", err)
	}
	fmt.Println()

	result := logic.GetChangeReturn(parsedAmount, change)
	fmt.Println()
	fmt.Println("Total result")
	fmt.Println(logic.PrettyFormatChange(result))
}
