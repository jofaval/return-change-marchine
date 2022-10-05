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

func getChangeReturn(amount int, change []ChangeType) Dictionary {
	remainingAmount := amount
	totalChange := Dictionary{}

	changeLen := len(change)
	for changeIndex := 0; changeIndex < changeLen; changeIndex++ {
		currentChange := change[changeIndex]
		parsedValue := currentChange.value

		if currentChange.units == EUROS || currentChange.units == DOLLARS {
			fmt.Println("Has to parse units")
			parsedValue *= 100
		}

		if parsedValue > remainingAmount {
			fmt.Println("Value is too high to compute", "parsed", parsedValue, "remaining", remainingAmount)
			continue
		}

		amountOfValue := remainingAmount / parsedValue
		if amountOfValue > currentChange.amount {
			fmt.Println("There's not enough amount of this currency to supply")
			return nil
		}

		totalChange[currentChange.name] = amountOfValue
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

func prettyFormatChange(rawChange Dictionary) string {
	if rawChange == nil {
		return "No change was available"
	}

	formatted := []string{}

	for key, value := range rawChange {
		splitted := strings.Split(key, ".")

		amount := splitted[0]
		units := splitted[1]

		formatted = append(formatted, fmt.Sprintf("%d of %s %s", value, amount, units))
	}

	return strings.Join(formatted, "\n")
}

func main() {
	change := initChange()

	// input := readInput("Enter text (introduce the value in cents please): ")
	input := "235101"

	// parse amount
	parsedAmount, err := strconv.Atoi(input)
	fmt.Println("raw", input, "parsed", parsedAmount, "error", err)
	fmt.Println()

	result := getChangeReturn(parsedAmount, change)
	fmt.Println()
	fmt.Println("Total result")
	fmt.Println(prettyFormatChange(result))
}
