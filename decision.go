package main

type DecisionType uint

const (
	FOLD DecisionType = iota
	CALL
	RAISE
	ALL_IN
)

type Decision struct {
	Decision DecisionType
	Amount   int
}
