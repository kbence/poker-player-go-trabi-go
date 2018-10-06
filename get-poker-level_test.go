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
	level := GetPokerLevel(cards)
	assert.Equal(t, 2, level, "couple")
}

func TestTwoPairs(t *testing.T) {
	cards := []Card{
		Card{Rank: "2", Suit: "clubs"},
		Card{Rank: "2", Suit: "spades"},
		Card{Rank: "3", Suit: "clubs"},
		Card{Rank: "3", Suit: "spades"},
	}
	level := GetPokerLevel(cards)
	assert.Equal(t, 3, level, "two pairs")

	cards = []Card{
		Card{Rank: "2", Suit: "clubs"},
		Card{Rank: "2", Suit: "spades"},
		Card{Rank: "3", Suit: "clubs"},
		Card{Rank: "3", Suit: "spades"},
		Card{Rank: "4", Suit: "clubs"},
		Card{Rank: "4", Suit: "spades"},
	}
	level = GetPokerLevel(cards)
	assert.Equal(t, 3, level, "if 3 pairs")
}

func TestDrill(t *testing.T) {
	cards := []Card{
		Card{Rank: "2", Suit: "clubs"},
		Card{Rank: "2", Suit: "spades"},
		Card{Rank: "2", Suit: "hearts"},
	}
	level := GetPokerLevel(cards)
	assert.Equal(t, 4, level, "Drill")
}
