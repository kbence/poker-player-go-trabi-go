package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHighCard(t *testing.T) {
	cards := []Card{
		Card{Rank: "K", Suit: "clubs"},
		Card{Rank: "3", Suit: "spades"},
	}
	level, cardLevel := GetPokerLevel(cards)
	assert.Equal(t, 1, level, "high card")
	assert.Equal(t, 13, cardLevel, "high card level")
}

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

func TestFlush(t *testing.T) {
	cards := []Card{
		Card{Rank: "2", Suit: "clubs"},
		Card{Rank: "3", Suit: "clubs"},
		Card{Rank: "4", Suit: "clubs"},
		Card{Rank: "5", Suit: "clubs"},
		Card{Rank: "6", Suit: "clubs"},
	}
	level, cardLevel := GetPokerLevel(cards)
	assert.Equal(t, 6, level, "flush")
	assert.Equal(t, 1, cardLevel, "flush card level")
}

func TestPoker(t *testing.T) {
	cards := []Card{
		Card{Rank: "3", Suit: "clubs"},
		Card{Rank: "3", Suit: "spades"},
		Card{Rank: "3", Suit: "hearts"},
		Card{Rank: "3", Suit: "diamonds"},
		Card{Rank: "6", Suit: "clubs"},
	}
	level, cardLevel := GetPokerLevel(cards)
	assert.Equal(t, 8, level, "poker")
	assert.Equal(t, 1, cardLevel, "poker card level")
}

func TestRoyalFlush(t *testing.T) {
	cards := []Card{
		Card{Rank: "10", Suit: "clubs"},
		Card{Rank: "J", Suit: "clubs"},
		Card{Rank: "Q", Suit: "clubs"},
		Card{Rank: "K", Suit: "clubs"},
		Card{Rank: "A", Suit: "clubs"},
	}
	level, cardLevel := GetPokerLevel(cards)
	assert.Equal(t, 10, level, "royal flush")
	assert.Equal(t, 1, cardLevel, "royal flush card level")
}
