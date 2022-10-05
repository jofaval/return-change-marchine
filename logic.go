package main

import (
	"fmt"
	"strings"
)

func GetChangeReturn(amount int, change []ChangeType) Dictionary {
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
			// Early return
			// fmt.Println("There's not enough amount of this currency to supply")
			// return nil

			fmt.Println("There's not enough amount of this currency to supply, it will attempt to compensate with lower value currency")
			amountOfValue = parsedValue * currentChange.amount
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

func PrettyFormatChange(rawChange Dictionary) string {
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
