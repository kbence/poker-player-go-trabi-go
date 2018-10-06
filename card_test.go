package main

import "testing"

func TestValue(t *testing.T) {
	c := Card{"K", "diamonds"}
	calculated := c.Value()
	expected := 8.0
	if calculated != expected {
		t.Errorf("Value was incorrect, got: %f, want: %f.", calculated, expected)
	}
}
