package main

import (
	"fmt"
	"strings"
)

func shouldScaleUpUnit(units int) bool {
	return units == EUROS || units == DOLLARS
}

func changeWouldBeTooHigh(parsedValue int, remainingAmount int) bool {
	return parsedValue > remainingAmount
}

func cantMakeUpForAmount(amountOfValue int, amount int) bool {
	return amountOfValue > amount
}

func thereWasntEnoughChange(totalChange MapResultChangeType, remainingAmount int) bool {
	return len(totalChange) == 0 || remainingAmount > 0
}

func alreadyReturnedChange(remainingAmount int) bool {
	return remainingAmount <= 0
}

func GetChangeReturn(amount int, change []ChangeType) MapResultChangeType {
	remainingAmount := amount
	totalChange := MapResultChangeType{}

	changeLen := len(change)
	for changeIndex := 0; changeIndex < changeLen; changeIndex++ {
		currentChange := change[changeIndex]
		parsedValue := currentChange.value

		if shouldScaleUpUnit(currentChange.units) {
			if DEBUG_MODE {
				fmt.Println("Has to parse units")
			}
			parsedValue *= 100
		}

		if DEBUG_MODE {
			fmt.Println(parsedValue)
		}
		if changeWouldBeTooHigh(parsedValue, remainingAmount) {
			if DEBUG_MODE {
				fmt.Println("Value is too high to compute", "parsed", parsedValue, "remaining", remainingAmount)
			}
			continue
		}

		amountOfValue := remainingAmount / parsedValue
		if cantMakeUpForAmount(amountOfValue, currentChange.amount) {
			// Early return
			// if DEBUG_MODE {
			// 	fmt.Println("There's not enough amount of this currency to supply")
			// }
			// return nil

			if DEBUG_MODE {
				fmt.Println("There's not enough amount of this currency to supply, it will attempt to compensate with lower value currency")
			}

			amountOfValue = currentChange.amount
			remainingAmount = remainingAmount - parsedValue*currentChange.amount
		} else {
			remainingAmount = remainingAmount % parsedValue
		}

		totalChange[currentChange.name] = ResultChangeType{
			value:  parsedValue,
			amount: amountOfValue,
			units:  currentChange.units,
		}

		if alreadyReturnedChange(remainingAmount) {
			if DEBUG_MODE {
				fmt.Println("There's no remaining amount, early escaping, hopefully")
			}
			break
		}
	}

	// If no change, or not enough, available, "raise an exception"
	if thereWasntEnoughChange(totalChange, remainingAmount) {
		if DEBUG_MODE {
			fmt.Println("There wasn't enough change")
		}
		return nil
	}

	return totalChange
}

func PrettyFormatChange(rawChange MapResultChangeType) string {
	if rawChange == nil {
		return "No change was available"
	}

	formatted := []string{}

	for _, returnChangeResult := range rawChange {
		parsedUnit := ""
		parsedValue := returnChangeResult.value

		switch returnChangeResult.units {
		case EUROS:
			parsedUnit = "EUROS"
			parsedValue /= 100
		case DOLLARS:
			parsedUnit = "DOLLARS"
			parsedValue /= 100
		case CENTS:
			parsedUnit = "CENTS"
		}

		formatted = append(formatted, fmt.Sprintf("%d of %d %s", returnChangeResult.amount, parsedValue, parsedUnit))
	}

	return strings.Join(formatted, "\n")
}
