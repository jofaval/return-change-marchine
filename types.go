package main

type Dictionary map[string]interface{}

// unit types
const (
	CENTS   = iota
	EUROS   = iota
	DOLLARS = iota
)

// input type
type ChangeType struct {
	name   string
	value  int
	units  int
	amount int
}
