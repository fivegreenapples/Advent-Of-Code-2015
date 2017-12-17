package main

import (
	"fmt"
	"os"
)

func main() {
	testFailure := false
	for t, expected := range testsPart1 {
		calculated := decideQualityPart1(t)
		if calculated != expected {
			fmt.Printf("Part 1 quality test failed for %s. Expected %v, got %v.\n", t, expected, calculated)
			testFailure = true
		}
	}
	for t, expected := range testsPart2 {
		calculated := decideQualityPart2(t)
		if calculated != expected {
			fmt.Printf("Part 2 quality test failed for %s. Expected %v, got %v.\n", t, expected, calculated)
			testFailure = true
		}
	}
	if testFailure {
		fmt.Printf("\nStopping. Test failures.\n\n")
		os.Exit(1)
	}

	part1NiceCount := countNiceStrings(input, decideQualityPart1)
	fmt.Printf("Part 1 - nice count is %d\n", part1NiceCount)
	part2NiceCount := countNiceStrings(input, decideQualityPart2)
	fmt.Printf("Part 2 - nice count is %d\n", part2NiceCount)
}

func countNiceStrings(in []string, decider func(string) quality) int {
	count := 0
	for _, s := range input {
		if decider(s) == nice {
			count++
		}
	}
	return count
}

func decideQualityPart2(in string) quality {

	hasSeparatedDouble := false
	hasDuplicatedPair := false
	pairCounts := map[string]int{}

	var prevRune rune
	var prevPrevRune rune
	var immediatelyPriorPair string

	for i, r := range in {

		if r == prevPrevRune {
			hasSeparatedDouble = true
		}

		if i >= 1 {
			pair := string(prevRune) + string(r)
			if pair != immediatelyPriorPair {
				pairCounts[pair]++
				if pairCounts[pair] >= 2 {
					hasDuplicatedPair = true
				}
				immediatelyPriorPair = pair
			} else {
				immediatelyPriorPair = ""
			}
		}

		prevPrevRune = prevRune
		prevRune = r
	}
	if hasSeparatedDouble && hasDuplicatedPair {
		return nice
	}
	return naughty

}

func decideQualityPart1(in string) quality {

	vowelCount := 0
	hasDouble := false
	hasBadPair := false // ab, cd, pq, or xy

	var prevRune rune

	for _, r := range in {

		if r == prevRune {
			hasDouble = true
		}

		if r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u' {
			vowelCount++
		}

		if r == 'b' && prevRune == 'a' {
			hasBadPair = true
		}
		if r == 'd' && prevRune == 'c' {
			hasBadPair = true
		}
		if r == 'q' && prevRune == 'p' {
			hasBadPair = true
		}
		if r == 'y' && prevRune == 'x' {
			hasBadPair = true
		}

		prevRune = r
	}
	if vowelCount >= 3 && hasDouble && !hasBadPair {
		return nice
	}
	return naughty
}
