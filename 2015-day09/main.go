package main

import "fmt"
import "sort"

func main() {
	testAllDistances := expandInput(testInput)
	testShortestJourney, testLongestJourney := findShortestAndLongestJourney(testAllDistances)
	fmt.Printf("Part 1 test: shortest journey is %d\n", testShortestJourney)
	fmt.Printf("Part 1 test: longest journey is %d\n", testLongestJourney)

	part1AllDistances := expandInput(input)
	part1ShortestJourney, part1LongestJourney := findShortestAndLongestJourney(part1AllDistances)
	fmt.Printf("Part 1: shortest journey is %d\n", part1ShortestJourney)
	fmt.Printf("Part 1: longest journey is %d\n", part1LongestJourney)

}

func expandInput(in map[city]distancesByCity) map[city]distancesByCity {
	out := map[city]distancesByCity{}
	for c, distances := range in {
		if _, found := out[c]; !found {
			out[c] = distancesByCity{}
		}
		for cc, d := range distances {
			if _, found := out[cc]; !found {
				out[cc] = distancesByCity{}
			}
			out[c][cc] = d
			out[cc][c] = d
		}
	}
	return out
}
func findShortestAndLongestJourney(cityDistances map[city]distancesByCity) (int, int) {
	allJourneys := findShortestJourneysByStart(cityDistances)

	journeyMins := []int{}
	journeyMaxes := []int{}
	for _, d := range allJourneys {
		journeyMins = append(journeyMins, d[0])
		journeyMaxes = append(journeyMaxes, d[1])
	}
	sort.Ints(journeyMins)
	sort.Sort(sort.Reverse(sort.IntSlice(journeyMaxes)))
	return journeyMins[0], journeyMaxes[0]
}
func findShortestJourneysByStart(cityDistances map[city]distancesByCity) map[city][2]int {

	journeysByStart := map[city][2]int{}

	for c := range cityDistances {
		destinations := []city{}
		for cc := range cityDistances {
			if cc == c {
				continue
			}
			destinations = append(destinations, cc)
		}
		min, max := findMinMaxJourneyFrom(c, destinations, cityDistances)
		journeysByStart[c] = [...]int{min, max}
	}

	return journeysByStart

}
func findMinMaxJourneyFrom(from city, to []city, cityDistances map[city]distancesByCity) (min, max int) {
	if len(to) == 0 {
		return 0, 0
	}
	if len(to) == 1 {
		return cityDistances[from][to[0]], cityDistances[from][to[0]]
	}

	journeyMins := []int{}
	journeyMaxes := []int{}
	for i, c := range to {
		remainingDestinations := append([]city{}, to[:i]...)
		remainingDestinations = append(remainingDestinations, to[i+1:]...)

		initialLength := cityDistances[from][c]
		remainingMin, remainingMax := findMinMaxJourneyFrom(c, remainingDestinations, cityDistances)
		journeyMins = append(journeyMins, initialLength+remainingMin)
		journeyMaxes = append(journeyMaxes, initialLength+remainingMax)
	}

	sort.Ints(journeyMins)
	sort.Sort(sort.Reverse(sort.IntSlice(journeyMaxes)))
	return journeyMins[0], journeyMaxes[0]
}
