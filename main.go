package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Dictionary map[string]interface{}

// unit types
const (
	CENTS   = iota
	EUROS   = iota
	DOLLARS = iota
)

// TODO: focus on the solution, the aim for the best solution where it knows the available change
// and can even be able to acknowledge that no change is available, too high or not precise

// input type
type ChangeType struct {
	name  string
	value int
	units int
}

func readInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	text, _ := reader.ReadString('\n')

	// convert CRLF to LF
	text = strings.Replace(text, "\n", "", -1)
	text = strings.Replace(text, "\r", "", -1)

	return text
}

func getChangeReturn(amount int, change []ChangeType) Dictionary {
	remainingAmount := amount
	totalChange := Dictionary{}

	changeLen := len(change)
	for i := 0; i < changeLen; i++ {
		currentChange := change[i]
		parsedValue := currentChange.value

		if currentChange.units == EUROS || currentChange.units == DOLLARS {
			fmt.Println("Has to parse units")
			parsedValue *= 100
		}

		if parsedValue > remainingAmount {
			fmt.Println("Value is too high to compute", "parsed", parsedValue, "remaining", remainingAmount)
			continue
		}

		totalChange[currentChange.name] = remainingAmount / parsedValue
		remainingAmount = remainingAmount % parsedValue

		if remainingAmount <= 0 {
			fmt.Println("There's no remaining amount, early escaping, hopefully")
			break
		}
	}

	// If no change, or not enough, available, "raise an exception"
	if len(totalChange) == 0 || remainingAmount > 0 {
		return nil
	}

	return totalChange
}

func main() {
	change := []ChangeType{
		{
			name:  "500.EUR",
			value: 500,
			units: EUROS,
		},
		{
			name:  "200.EUR",
			value: 200,
			units: EUROS,
		},
		{
			name:  "100.EUR",
			value: 100,
			units: EUROS,
		},
		{
			name:  "50.EUR",
			value: 50,
			units: EUROS,
		},
		{
			name:  "20.EUR",
			value: 20,
			units: EUROS,
		},
		{
			name:  "10.EUR",
			value: 10,
			units: EUROS,
		},
		{
			name:  "5.EUR",
			value: 5,
			units: EUROS,
		},
		{
			name:  "2.EUR",
			value: 2,
			units: EUROS,
		},
		{
			name:  "1.EUR",
			value: 1,
			units: EUROS,
		},
		{
			name:  "50.CENTS",
			value: 50,
			units: CENTS,
		},
		{
			name:  "20.CENTS",
			value: 20,
			units: CENTS,
		},
		{
			name:  "10.CENTS",
			value: 10,
			units: CENTS,
		},
		{
			name:  "5.CENTS",
			value: 5,
			units: CENTS,
		},
		{
			name:  "2.CENTS",
			value: 2,
			units: CENTS,
		},
		{
			name:  "1.CENTS",
			value: 1,
			units: CENTS,
		},
	}
	input := readInput("Enter text (introduce the value in cents please): ")
	// input := "235"

	// parse amount
	parsedAmount, err := strconv.Atoi(input)
	fmt.Println("raw", input, "parsed", parsedAmount, "error", err)

	result := getChangeReturn(parsedAmount, change)
	fmt.Println("Total result", result)
}
