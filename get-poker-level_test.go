package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCouple(t *testing.T) {
	cards := []Card{
		Card{Rank: "2", Suit: "clubs"},
		Card{Rank: "2", Suit: "spades"},
	}
	level, cardLevel := GetPokerLevel(cards)
	assert.Equal(t, 2, level, "couple")
	assert.Equal(t, 1, cardLevel, "couple card level")
}

func TestTwoPairs(t *testing.T) {
	cards := []Card{
		Card{Rank: "2", Suit: "clubs"},
		Card{Rank: "2", Suit: "spades"},
		Card{Rank: "3", Suit: "clubs"},
		Card{Rank: "3", Suit: "spades"},
	}
	level, cardLevel := GetPokerLevel(cards)
	assert.Equal(t, 3, level, "two pairs")
	assert.Equal(t, 1, cardLevel, "two pairs card level")

	cards = []Card{
		Card{Rank: "2", Suit: "clubs"},
		Card{Rank: "2", Suit: "spades"},
		Card{Rank: "3", Suit: "clubs"},
		Card{Rank: "3", Suit: "spades"},
		Card{Rank: "4", Suit: "clubs"},
		Card{Rank: "4", Suit: "spades"},
	}
	level, cardLevel = GetPokerLevel(cards)
	assert.Equal(t, 3, level, "if 3 pairs")
	assert.Equal(t, 1, cardLevel, "if 3 pairs card level")
}

func TestDrill(t *testing.T) {
	cards := []Card{
		Card{Rank: "2", Suit: "clubs"},
		Card{Rank: "2", Suit: "spades"},
		Card{Rank: "2", Suit: "hearts"},
	}
	level, cardLevel := GetPokerLevel(cards)
	assert.Equal(t, 4, level, "drill")
	assert.Equal(t, 1, cardLevel, "drill card level")
}
