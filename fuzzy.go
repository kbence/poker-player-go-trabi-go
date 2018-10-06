package main

var Cfg *Config = nil

func init() {
	Cfg = NewConfig()
}

var stackScale = FloatScale{}.Domain(0.0, 3000.0).Range(1.0, 0.0)
var raiseScale = FloatScale{}.Domain(0.0, 1000.0).Range(0.0, 1.0)
var handScale = FloatScale{}.Domain(0.0, 10.0).Range(0.0, 1.0)

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

type FuzzyDecider struct {
	config     *Config
	game       *Game
	raiseCurve func(float64) float64
	stackCurve func(float64) float64
	handCurve  func(float64) float64
	log        *Logger
}

func NewFuzzyDecider(config *Config, game *Game) *FuzzyDecider {
	decider := &FuzzyDecider{config: config, game: game}
	decider.stackCurve = Curve(config.Curves.StackCurve)
	decider.raiseCurve = Curve(config.Curves.RaiseCurve)
	decider.handCurve = Curve(config.Curves.HandCurve)
	decider.log = NewLogger(game.GameID)

	return decider
}

func (d *FuzzyDecider) Next() Decision {
	stackValue := stackScale.Scale((float64)(d.game.Player().Stack))
	raiseValue := raiseScale.Scale((float64)(d.game.MinimumRaise))
	handValue := handScale.Scale(
		ValueOfCards(d.game.EffectiveCards()) - ValueOfCards(d.game.CommunityCards))
	confidence := d.stackCurve(stackValue) * d.raiseCurve(raiseValue) * d.handCurve(handValue)

	d.log.Debugf("stackValue=%f, raiseValue=%f, handValue=%f -> confidence=%f\n",
		stackValue, raiseValue, handValue, confidence)

	switch {
	case confidence > d.config.ConfidenceLevels.AllIn:
		d.log.Debugf("all in with confidence %f", confidence)
		return ALL_IN
	case confidence > d.config.ConfidenceLevels.Raise:
		d.log.Debugf("raise with confidence %f", confidence)
		return RAISE
	case confidence > d.config.ConfidenceLevels.Call:
		d.log.Debugf("call with confidence %f", confidence)
		return CALL
	}
	return FOLD
}
