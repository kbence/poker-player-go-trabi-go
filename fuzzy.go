package main

var Cfg *Config = nil

func init() {
	Cfg = NewConfig()
}

type FuzzyDecider struct {
	config     *Config
	game       *Game
	raiseCurve func(float64) float64
	stackCurve func(float64) float64
}

var stackScale = FloatScale{}.Domain(0.0, 3000.0).Range(0.0, 0.1)
var raiseScale = FloatScale{}.Domain(0.0, 1000.0).Range(0.0, 0.1)

func EaseCubicOut(v float64) float64 {
	return v * v * v
}

func EaseCubicIn(v float64) float64 {
	return 1.0 - v*v*v
}

func Lerp(a, b, t float64) float64 {
	return (1.0-t)*a + t*b
}

func Curve(c float64) func(float64) float64 {
	return func(v float64) float64 {
		return Lerp(EaseCubicOut(v), EaseCubicIn(v), (c+1.0)/2.0)
	}
}

func NewFuzzyDecider(config *Config, game *Game) *FuzzyDecider {
	decider := &FuzzyDecider{config: config, game: game}
	decider.stackCurve = Curve(config.StackCurve)
	decider.raiseCurve = Curve(config.RaiseCurve)

	return decider
}

func (d *FuzzyDecider) Next() Decision {
	return Decision{Decision: ALL_IN}
}
