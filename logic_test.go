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
