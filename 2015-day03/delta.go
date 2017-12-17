package main

import (
	"fmt"
)

type delta struct {
	dX int
	dY int
}

func makeDeltaFromRune(r rune) delta {
	switch r {
	case '^':
		return delta{0, 1}
	case '>':
		return delta{1, 0}
	case '<':
		return delta{-1, 0}
	case 'v':
		return delta{0, -1}
	default:
		panic(fmt.Errorf("can't create delta - unhandled rune '%v'", r))
	}
}
