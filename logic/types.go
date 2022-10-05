package logic

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

// return result
type ResultChangeType struct {
	value  int
	units  int
	amount int
}

type MapResultChangeType map[string]ResultChangeType
