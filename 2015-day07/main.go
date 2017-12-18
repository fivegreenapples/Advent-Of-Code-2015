package main

import "fmt"

func main() {

	testCircuit := makeCircuit()
	testCircuit.parseAllInstructions(testInput)
	fmt.Printf("Test circuit:\n%v\n", testCircuit)

	part1Circuit := makeCircuit()
	part1Circuit.parseAllInstructions(input)
	fmt.Printf("Part1 circuit:\n%v\n", part1Circuit)
	fmt.Printf("Signal on wire a is %d\n", part1Circuit.valueForWire(wireIdentifier("a")))

	wireAValue := part1Circuit.valueForWire(wireIdentifier("a"))
	part1Circuit.reset()
	part1Circuit.setWireToValue(wireIdentifier("b"), wireAValue)
	fmt.Printf("Part2 circuit:\n%v\n", part1Circuit)
	fmt.Printf("Signal on wire a is now %d\n", part1Circuit.valueForWire(wireIdentifier("a")))

}
