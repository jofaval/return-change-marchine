package main

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
	name   string
	value  int
	units  int
	amount int
}
