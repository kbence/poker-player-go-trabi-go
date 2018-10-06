package main

import _ "github.com/kbence/poker-player-go-trabi-go/fuzzy"

// VERSION provides a short description of the player's current version
// The string specified here will be shown for live.leanpoker.org games
const VERSION = "Deciding go player"

// PokerPlayer is a struct to organize player methods
type PokerPlayer struct{}

// NewPokerPlayer creates a new instance of *PokerPlayer
func NewPokerPlayer() *PokerPlayer {
	return &PokerPlayer{}
}

// BetRequest handles the main betting logic. The return value of this
// function will be used to decide whether the player want to fold,
// call, raise or do an all-in; more information about this behaviour
// can be found here: http://leanpoker.org/player-api
func (p *PokerPlayer) BetRequest(state *Game) int {
	value := (Hole)(state.Player().HoleCards).Value()

	if value >= 12.0 {
		return state.MinimumRaiseValue()
	}

	if value >= 9.0 {
		return state.CallValue()
	}

	return 0
}

// Showdown is called at the end of every round making it possible to
// e.g. collect statistics or log end results of the games
func (p *PokerPlayer) Showdown(state *Game) {

}

// Version returns the version string that will be shown on the UI at
// live.leanpoker.org
func (p *PokerPlayer) Version() string {
	return VERSION
}
