package main

import (
	"bufio"
	"fmt"
	"os"
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
	change := initChange()

	// input := readInput("Enter text (introduce the value in cents please): ")
	input := "235101"

	// parse amount
	parsedAmount, err := strconv.Atoi(input)
	fmt.Println("raw", input, "parsed", parsedAmount, "error", err)
	fmt.Println()

	result := GetChangeReturn(parsedAmount, change)
	fmt.Println()
	fmt.Println("Total result")
	fmt.Println(PrettyFormatChange(result))
}
