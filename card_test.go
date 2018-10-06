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

func TestHoleValue_SameSuitAdds2(t *testing.T) {
	q := Card{"Q", "hearts"}
	sameSuitEight := Card{"8", "hearts"}
	otherSuitEight := Card{"8", "diamonds"}

	sameSuitValue := Hole([]Card{q, sameSuitEight}).Value()
	otherSuitValue := Hole([]Card{q, otherSuitEight}).Value()
	assertEquals(sameSuitValue-otherSuitValue, 2, t)
}

func TestHoleValue_QKAPair(t *testing.T) {
	q1 := Card{"Q", "hearts"}
	q2 := Card{"Q", "diamonds"}

	assertEquals(Hole([]Card{q1, q2}).Value(), 14, t)
}

func TestHoleValue_LowerThanQKAPair(t *testing.T) {
	q1 := Card{"J", "hearts"}
	q2 := Card{"J", "diamonds"}

	assertEquals(Hole([]Card{q1, q2}).Value(), 13, t)
}

func TestHoleValue_SmallPair(t *testing.T) {
	q1 := Card{"4", "hearts"}
	q2 := Card{"4", "diamonds"}

	assertEquals(Hole([]Card{q1, q2}).Value(), 6, t)
}

func assertEquals(calculated float64, expected float64, t *testing.T) {
	if calculated != expected {
		t.Errorf("Value was incorrect, got: %f, want: %f.", calculated, expected)
	}
}
