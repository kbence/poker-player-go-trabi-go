package main

var Cfg *Config = nil

func init() {
	Cfg = NewConfig()
}

type DecisionType uint

const (
	ALL_IN DecisionType = iota
)

type Decision struct {
	Decision DecisionType
}

type FuzzyDecider struct {
}

func NewFuzzyDecider() *FuzzyDecider {
	decider := &FuzzyDecider{}

	return decider
}

func (d *FuzzyDecider) Next(state *Game) Decision {
	return Decision{Decision: ALL_IN}
}
