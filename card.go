package main

import (
	"log"
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

func (c Card) Position() int {
	switch c.Rank {
	case "A":
		return 14
	case "K":
		return 13
	case "Q":
		return 12
	case "J":
		return 11
	default:
		x, _ := strconv.Atoi(c.Rank)
		return x
	}
}

func(c Card) gap(other Card)  int {
	p1, p2:= c.Position(), other.Position()
	return int(math.Abs(float64(p1-p2)))
}

type Hole []Card

func (h Hole) Value() float64 {
	a, b := h[0], h[1]
	value := math.Max(a.Value(), b.Value())
	if(a.Value()== b.Value()) {value*=2}
	inSuit := a.Suit == b.Suit
	if inSuit {value+=2}
	g:=a.gap(b)

	switch g {
	case 0:break
	case 1:value -= 1
	case 2:value-=2
	case 3:value-=4
	default:
		value-=5
	}

	if (a.Position()<12 && b.Position()<12 && g<=1) {value+=1}

	rounded := math.Round(value)
	log.Printf("%s %s %s %s [v1=%v v2=%v g=%v %v] => %v", a.Rank ,a.Suit, b.Rank, a.Suit,
		a.Value(), b.Value(), g, inSuit,
		rounded)
	return rounded
}
