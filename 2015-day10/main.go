package main

import "fmt"

func main() {
	testEncoded := encodeMultiple([]byte(testInput), 5)
	fmt.Printf("Test part 1 - after 5 encodings, the result is %s having length %d.\n", testEncoded, len(testEncoded))

	part1Encoded := encodeMultiple([]byte(input), 40)
	fmt.Printf("Part 1 - after 40 encodings, the result has length %d.\n", len(part1Encoded))

	part2Encoded := encodeMultiple([]byte(input), 50)
	fmt.Printf("Part 2 - after 50 encodings, the result has length %d.\n", len(part2Encoded))
}

func encodeMultiple(in []byte, numRepeats int) string {
	out := in
	for numRepeats > 0 {
		// fmt.Println(numRepeats)
		out = encode(out)
		numRepeats--
	}
	return string(out)
}

func encode(in []byte) []byte {
	if len(in) == 0 {
		return []byte{}
	}

	strconvMap := map[int]byte{
		1: '1',
		2: '2',
		3: '3',
	}

	out := []byte{}
	prevChar := in[0]
	runningCount := 1
	for i := 1; i < len(in); i++ {
		c := in[i]
		if c == prevChar {
			runningCount++
			continue
		}

		out = append(out, strconvMap[runningCount], prevChar)
		runningCount = 1
		prevChar = c
	}
	out = append(out, strconvMap[runningCount], prevChar)
	return out
}
