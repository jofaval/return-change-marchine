package logic

import (
	"fmt"
	"return-change-machine/utils"
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
			utils.DebugPrint("Has to parse units")
			parsedValue *= 100
		}

		utils.DebugPrint(parsedValue)
		if changeWouldBeTooHigh(parsedValue, remainingAmount) {
			utils.DebugPrint("Value is too high to compute", "parsed", parsedValue, "remaining", remainingAmount)
			continue
		}

		amountOfValue := remainingAmount / parsedValue
		if cantMakeUpForAmount(amountOfValue, currentChange.amount) {
			// Early return
			// utils.DebugPrint.Println("There's not enough amount of this currency to supply")
			// return nil
			utils.DebugPrint("There's not enough amount of this currency to supply, it will attempt to compensate with lower value currency")

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
			utils.DebugPrint("There's no remaining amount, early escaping, hopefully")
			break
		}
	}

	// If no change, or not enough, available, "raise an exception"
	if thereWasntEnoughChange(totalChange, remainingAmount) {

		utils.DebugPrint("There wasn't enough change")
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
