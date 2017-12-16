package main

import (
	"fmt"
	"os"
)

func main() {

	testFailure := false
	for directions, expectedFloor := range testInputs {
		possibleFloor := calcFloor(directions)
		if possibleFloor != expectedFloor {
			fmt.Printf("Test failed for %s. Expecting %d, got %d\n", directions, expectedFloor, possibleFloor)
			testFailure = true
		}
	}
	if testFailure {
		fmt.Printf("\nStopping. Test failures.\n\n")
		os.Exit(1)
	}

	part1Floor := calcFloor(input)
	fmt.Printf("Part1: Santa's floor is %d\n", part1Floor)

	basementTestFailure := false
	for directions, expected := range basementEntryTests {
		possible := calcBasementEntry(directions)
		if possible != expected {
			fmt.Printf("Basement test failed for %s. Expecting %d, got %d\n", directions, expected, possible)
			basementTestFailure = true
		}
	}
	if basementTestFailure {
		fmt.Printf("\nStopping. Test failures.\n\n")
		os.Exit(1)
	}

	part2 := calcBasementEntry(input)
	fmt.Printf("Part1: Santa enters basement at instruction number %d\n", part2)
}

func calcFloor(directions string) int {
	currentFloor := 0
	for _, d := range directions {
		if d == '(' {
			currentFloor++
		} else if d == ')' {
			currentFloor--
		} else {
			panic(fmt.Errorf("unhandled direction %c", d))
		}
	}
	return currentFloor
}

func calcBasementEntry(directions string) int {
	currentFloor := 0
	for i, d := range directions {
		if d == '(' {
			currentFloor++
		} else if d == ')' {
			currentFloor--
		} else {
			panic(fmt.Errorf("unhandled direction %c", d))
		}
		if currentFloor < 0 {
			return i + 1
		}
	}
	return -1
}
