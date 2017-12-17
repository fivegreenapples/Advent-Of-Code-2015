package main

import "github.com/fivegreenapples/adventofcode2015/2015-day06/day6"
import "fmt"

func main() {
	theGridPart1 := day6.MakeGrid(1000, 1000)
	theGridPart1.ApplyInstructionList(input)
	fmt.Printf("Part 1 - number of lights now on is %d\n", theGridPart1.NumberOfLightsTurnedOn())

	theGridPart2 := day6.MakeGrid2(1000, 1000)
	theGridPart2.ApplyInstructionList(input)
	fmt.Printf("Part 2 - total brightness is %d\n", theGridPart2.TotalBrightness())
}
