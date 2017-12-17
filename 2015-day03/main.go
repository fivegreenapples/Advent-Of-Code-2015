package main

import (
	"fmt"
	"os"
)

func main() {
	testFailure := false
	for t, expected := range tests {
		numLocations := findUniqueLocations(t, coord{0, 0})
		if numLocations != expected {
			fmt.Printf("Unique location test failed for %v. Expected %d, got %d.\n", t, expected, numLocations)
			testFailure = true
		}
	}
	for t, expected := range testsRoboSanta {
		numLocations := findUniqueLocationsWithRoboSanta(t, coord{0, 0})
		if numLocations != expected {
			fmt.Printf("Robo santa test failed for %v. Expected %d, got %d.\n", t, expected, numLocations)
			testFailure = true
		}
	}
	if testFailure {
		fmt.Printf("\nStopping. Test failures.\n\n")
		os.Exit(1)
	}

	fmt.Printf("Santa visits %d different houses.\n", findUniqueLocations(input, coord{0, 0}))
	fmt.Printf("Santa & Robo-Santa visit %d different houses.\n", findUniqueLocationsWithRoboSanta(input, coord{0, 0}))
}

func findUniqueLocations(moves string, current coord) uint {
	locations := map[coord]int{}
	locations[current]++
	for _, r := range moves {
		current.applyDelta(makeDeltaFromRune(r))
		locations[current]++
	}
	return uint(len(locations))
}
func findUniqueLocationsWithRoboSanta(moves string, current coord) uint {
	locations := map[coord]int{}
	locations[current]++
	roboCurrent := current
	for i, r := range moves {
		if i%2 == 0 {
			current.applyDelta(makeDeltaFromRune(r))
			locations[current]++
		} else {
			roboCurrent.applyDelta(makeDeltaFromRune(r))
			locations[roboCurrent]++
		}
	}
	return uint(len(locations))
}
