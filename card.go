package main

import (
	"math"
	"strconv"
)

// Card describes the properties of a card in the game state JSON
type Card struct {
	// Rank of the card. Possible values are numbers 2-10 and J,Q,K,A
	Rank string `json:"rank"`

	// Suit of the card. Possible values are: clubs,spades,hearts,diamonds
	Suit string `json:"suit"`
}

func (c Card) Value() float64 {
	switch c.Rank {
	case "A":
		return 10.0
	case "K":
		return 8.0
	case "Q":
		return 7.0
	case "J":
		return 6.0
	default:
		x, _ := strconv.ParseFloat(c.Rank, 64)
		return x / 2.0
	}
}

type Hole []Card

func (h Hole) Value() float64 {
	a, b := h[0], h[1]
	value := math.Max(a.Value(), b.Value())
	if(a.Value()== b.Value()) {value*=2}
	if(a.Suit==b.Suit){value+=2}
	return value
}
