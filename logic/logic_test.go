package logic

import (
	"testing"
)

func TestShouldScaleUpUnitIsEuros(t *testing.T) {
	got := shouldScaleUpUnit(EUROS)
	if !got {
		t.Error("Failed")
	}
}

func TestShouldScaleUpUnitIsDollars(t *testing.T) {
	got := shouldScaleUpUnit(DOLLARS)
	if !got {
		t.Error("Failed")
	}
}

func TestShouldScaleUpUnitIsNeither(t *testing.T) {
	got := shouldScaleUpUnit(CENTS)
	if got {
		t.Error("Failed")
	}
}

func TestChangeWouldBeTooHigh(t *testing.T) {
	got := changeWouldBeTooHigh(20, 5)
	if !got {
		t.Error("Failed")
	}
}

func TestChangeWouldBeTooHighButIsnt(t *testing.T) {
	got := changeWouldBeTooHigh(5, 20)
	if got {
		t.Error("Failed")
	}
}

func TestCantMakeUpForAmount(t *testing.T) {
	got := cantMakeUpForAmount(5, 20)
	if got {
		t.Error("Failed")
	}
}

func TestCantMakeUpForAmountButActuallyCan(t *testing.T) {
	got := cantMakeUpForAmount(20, 5)
	if !got {
		t.Error("Failed")
	}
}

func TestThereWasntEnoughChangeNoChange(t *testing.T) {
	got := thereWasntEnoughChange(MapResultChangeType{}, 20)
	if !got {
		t.Error("Failed")
	}
}

func TestThereWasntEnoughChangeThereStillRemaining(t *testing.T) {
	got := thereWasntEnoughChange(nil, 20)
	if !got {
		t.Error("Failed")
	}
}

func TestThereWasntEnoughChangeThereWas(t *testing.T) {
	got := thereWasntEnoughChange(MapResultChangeType{
		"1.EUR": {
			value:  1,
			amount: 50,
			units:  EUROS,
		},
	}, 10)
	if !got {
		t.Error("Failed")
	}
}

func TestAlreadyReturnedChangeAmountIsZero(t *testing.T) {
	got := alreadyReturnedChange(0)
	if !got {
		t.Error("Failed")
	}
}

func TestAlreadyReturnedChangeAmountIsNegative(t *testing.T) {
	got := alreadyReturnedChange(-1)
	if !got {
		t.Error("Failed")
	}
}

func TestAlreadyReturnedChangeButDidnt(t *testing.T) {
	got := alreadyReturnedChange(10)
	if got {
		t.Error("Failed")
	}
}

func TestGetChangeReturn(t *testing.T) {
	got := GetChangeReturn(100, []ChangeType{
		{
			name:   "1.EUR",
			value:  1,
			units:  EUROS,
			amount: 1,
		},
	})
	if got["1.EUR"].amount != 1 {
		t.Error("Expected to get 1 euro")
	}
}

func TestGetChangeReturnNotEnoughAmount(t *testing.T) {
	got := GetChangeReturn(300, []ChangeType{
		{
			name:   "1.EUR",
			value:  1,
			units:  EUROS,
			amount: 1,
		},
	})
	if len(got) != 0 {
		t.Error("Expected to not have enough")
	}
}

func TestGetChangeReturnHasTooMuch(t *testing.T) {
	got := GetChangeReturn(300, []ChangeType{
		{
			name:   "500.CENTS",
			value:  1,
			units:  CENTS,
			amount: 1,
		},
	})
	if len(got) != 0 {
		t.Error("Expected to not have enough")
	}
}

func TestGetChangeReturnPaysWithLowerCurrency(t *testing.T) {
	got := GetChangeReturn(300, []ChangeType{
		{
			name:   "5.EUR",
			value:  5,
			units:  EUROS,
			amount: 1,
		},
		{
			name:   "1.EUR",
			value:  1,
			units:  EUROS,
			amount: 100,
		},
	})
	if got["1.EUR"].amount != 3 {
		t.Error("Expected to not have enough")
	}
}

func TestGetChangeReturnPaysAndStillHasMoneyLeft(t *testing.T) {
	got := GetChangeReturn(500, []ChangeType{
		{
			name:   "5.CENTS",
			value:  5,
			units:  CENTS,
			amount: 1000,
		},
		{
			name:   "1.EUR",
			value:  1,
			units:  EUROS,
			amount: 100,
		},
	})
	if got["5.CENTS"].amount != 100 {
		t.Error("Expected to still have money left")
	}
}

func TestPrettyFormatChangeNoInput(t *testing.T) {
	got := PrettyFormatChange(nil)
	if got != "No change was available" {
		t.Error("Didn't return the fallback")
	}
}

func TestPrettyFormatChangeEuros(t *testing.T) {
	got := PrettyFormatChange(MapResultChangeType{
		"3.EUR": {
			amount: 3,
			units:  EUROS,
			// Mulitiplied by 100 because it's the result of parsing euros to cents
			value: 300,
		},
	})

	if got != "3 of 3 EUROS" {
		t.Error("Didn't match the expected output")
	}
}

func TestPrettyFormatChangeDollars(t *testing.T) {
	got := PrettyFormatChange(MapResultChangeType{
		"3.DOLLARS": {
			amount: 3,
			units:  DOLLARS,
			// Mulitiplied by 100 because it's the result of parsing euros to cents
			value: 300,
		},
	})

	if got != "3 of 3 DOLLARS" {
		t.Error("Didn't match the expected output")
	}
}

func TestPrettyFormatChangeCents(t *testing.T) {
	got := PrettyFormatChange(MapResultChangeType{
		"3.CENTS": {
			amount: 3,
			units:  CENTS,
			// Mulitiplied by 100 because it's the result of parsing euros to cents
			value: 3,
		},
	})

	if got != "3 of 3 CENTS" {
		t.Error("Didn't match the expected output")
	}
}

func TestInitChange(t *testing.T) {
	got := InitChange()
	if len(got) <= 1 {
		t.Error("Didn't load enough change elements")
	}
}
