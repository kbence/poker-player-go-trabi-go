package main

import "testing"

func TestValue(t *testing.T) {
	c := Card{"K", "diamonds"}
	calculated := c.Value()
	expected := 8.0
	assertEquals(calculated, expected, t)
}

func TestHoleValue_Q_hearts_8_hearts(t *testing.T) {
	q := Card{"Q", "hearts"}
	eight := Card{"8", "hearts"}
	hole := Hole([]Card{q, eight})
	hole.Value()

	assertEquals(hole.Value(), 4, t)
}

func assertEquals(calculated float64, expected float64, t *testing.T) {
	if calculated != expected {
		t.Errorf("Value was incorrect, got: %f, want: %f.", calculated, expected)
	}
}
