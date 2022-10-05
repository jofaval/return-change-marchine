package main

import (
	"testing"
)

func TestGetChangeReturn(t *testing.T) {
	got := GetChangeReturn(100, []ChangeType{
		{
			name:   "1.EUR",
			value:  1,
			units:  EUROS,
			amount: 1,
		},
	})
	if got["1.EUR"] != 1 {
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
	if got["1.EUR"] != 3 {
		t.Error("Expected to not have enough")
	}
}

func TestPrettyFormatChangeNoInput(t *testing.T) {
	got := PrettyFormatChange(nil)
	if got != "No change was available" {
		t.Error("Didn't return the fallback")
	}
}

func TestPrettyFormatChange(t *testing.T) {
	got := PrettyFormatChange(Dictionary{
		"3.EUR": 3,
	})

	if got != "3 of 3 EUR" {
		t.Error("Didn't match the expected output")
	}
}
