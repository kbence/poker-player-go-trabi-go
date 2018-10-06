package main

import "log"

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

var stackScale = FloatScale{}.Domain(0.0, 3000.0).Range(1.0, 0.0)
var raiseScale = FloatScale{}.Domain(0.0, 1000.0).Range(0.0, 1.0)
var handScale = FloatScale{}.Domain(0.0, 7.0).Range(0.0, 1.0)

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

func ValueOfCards(cards []Card) float64 {
	handLevel, cardLevel := GetPokerLevel(cards)

	return (float64)(handLevel) + (float64)(cardLevel)/15.0
}

func NewFuzzyDecider(config *Config, game *Game) *FuzzyDecider {
	decider := &FuzzyDecider{config: config, game: game}
	decider.stackCurve = Curve(config.Curves.StackCurve)
	decider.raiseCurve = Curve(config.Curves.RaiseCurve)

	return decider
}

func (d *FuzzyDecider) Next() Decision {
	stackValue := stackScale.Scale((float64)(d.game.Player().Stack))
	raiseValue := raiseScale.Scale((float64)(d.game.MinimumRaise))
	handValue := ValueOfCards(d.game.EffectiveCards()) - ValueOfCards(d.game.CommunityCards)
	confidence := stackValue * raiseValue * handValue

	log.Printf("stackValue=%f, raiseValue=%f, handValue=%f -> confidence=%f\n",
		stackValue, raiseValue, handValue, confidence)

	switch {
	case confidence > d.config.ConfidenceLevels.AllIn:
		log.Printf("all in with confidence %f", confidence)
		return ALL_IN
	case confidence > d.config.ConfidenceLevels.Raise:
		log.Printf("raise with confidence %f", confidence)
		return RAISE
	case confidence > d.config.ConfidenceLevels.Call:
		log.Printf("call with confidence %f", confidence)
		return CALL
	}
	return FOLD
}
