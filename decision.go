package main

type Decision uint

const (
	FOLD Decision = iota
	CALL
	RAISE
	ALL_IN
)
