package main

import "fmt"

func main() {
	fmt.Printf("Part 1 test: difference in lengths totals %d\n", findLengthDiffForInputs(testInput))
	fmt.Printf("Part 1: difference in lengths totals %d\n", findLengthDiffForInputs(input))

	fmt.Printf("Part 2 test: difference in lengths totals %d\n", findLengthDiffForInputs(encodeAll(testInput)))
	fmt.Printf("Part 2: difference in lengths totals %d\n", findLengthDiffForInputs(encodeAll(input)))

}

func encodeAll(in []string) []string {
	out := []string{}
	for _, s := range in {
		out = append(out, encode(s))
	}
	return out
}

func encode(in string) string {
	out := "\""
	for _, r := range in {
		if r == '\\' {
			out += "\\\\"
		} else if r == '"' {
			out += "\\\""
		} else {
			out += string(r)
		}
	}
	out += "\""
	return out
}

func findLengthDiffForInputs(in []string) int {
	diff := 0
	for _, s := range in {
		raw, memory := findLengths(s)
		diff += (raw - memory)
	}
	return diff
}

type parseState int

const (
	begin parseState = iota
	normal
	escapeBackslash
	escapeHex16
	escapeHex1
	end
)

func findLengths(in string) (raw, memory int) {
	raw = len(in)
	state := begin
	for _, r := range in {
		switch state {
		case begin:
			if r != '"' {
				panic("string doesn't start with double quotes: " + in)
			}
			state = normal
		case end:
			panic("parsing finished but not reached end of input: " + in)
		case escapeBackslash:
			if r == '\\' || r == '"' {
				memory++
				state = normal
			} else if r == 'x' {
				state = escapeHex16
			} else {
				panic("parsing failed after backslash: " + in)
			}
		case escapeHex16:
			if (r >= '0' && r <= '9') || (r >= 'a' && r <= 'f') {
				state = escapeHex1
			} else {
				panic("parsing failed after \\x: " + in)
			}
		case escapeHex1:
			if (r >= '0' && r <= '9') || (r >= 'a' && r <= 'f') {
				memory++
				state = normal
			} else {
				panic("parsing failed after \\x[0-9]: " + in)
			}
		case normal:
			if r == '\\' {
				state = escapeBackslash
			} else if r == '"' {
				state = end
			} else {
				memory++
			}
		default:
			panic("unahndled parsing state")
		}
	}

	if state != end {
		panic("premature end of input" + in)
	}
	return
}
