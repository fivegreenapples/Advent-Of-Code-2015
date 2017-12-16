package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {

	testFailure := false
	for t, expected := range testInputs {
		calcArea, calcRibbon := calculateRequiredPaperAndRibbon(t)
		if calcArea != expected[0] {
			fmt.Printf("Paper area test failed for %v. Expected %d, got %d.\n", t, expected[0], calcArea)
			testFailure = true
		}
		if calcRibbon != expected[1] {
			fmt.Printf("Ribbon length test failed for %v. Expected %d, got %d.\n", t, expected[1], calcRibbon)
			testFailure = true
		}
	}
	if testFailure {
		fmt.Printf("\nStopping. Test failures.\n\n")
		os.Exit(1)
	}

	totalArea, totalRibbon := calculateAllResources(input)
	fmt.Printf("Part1 - total area required is %d\n", totalArea)
	fmt.Printf("Part2 - total ribbon required is %d\n", totalRibbon)
}

func calculateRequiredPaperAndRibbon(dims dimensions) (paper, ribbon uint) {
	sort.Slice(dims[:], func(i, j int) bool {
		return dims[i] < dims[j]
	})
	a1 := dims[0] * dims[1]
	a2 := dims[0] * dims[2]
	a3 := dims[1] * dims[2]
	p1 := dims[0] + dims[1] + dims[0] + dims[1]
	v := dims[0] * dims[1] * dims[2]
	return a1 + a1 + a1 + a2 + a2 + a3 + a3, p1 + v
}

func calculateAllResources(allDims []dimensions) (paper, ribbon uint) {
	for _, d := range allDims {
		thisPaper, thisRibbon := calculateRequiredPaperAndRibbon(d)
		paper += thisPaper
		ribbon += thisRibbon
	}
	return
}
