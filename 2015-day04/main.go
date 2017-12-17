package main

import (
	"crypto/md5"
	"fmt"
	"os"
)

func main() {
	testFailure := false
	for t, expected := range tests {
		_, hashNum := findFirstHashWithPrefix(t, "00000")
		if hashNum != expected {
			fmt.Printf("Hash test failed for %v. Expected %d, got %d.\n", t, expected, hashNum)
			testFailure = true
		}
	}
	if testFailure {
		fmt.Printf("\nStopping. Test failures.\n\n")
		os.Exit(1)
	}

	_, part1HashNum := findFirstHashWithPrefix(input, "00000")
	fmt.Printf("PentaZero Hash number is %d.\n", part1HashNum)
	_, part2HashNum := findFirstHashWithPrefix(input, "000000")
	fmt.Printf("HexaZero Hash number is %d.\n", part2HashNum)
}

func findFirstHashWithPrefix(key string, prefix string) (string, uint) {
	i := 1
	for {
		hashInput := fmt.Sprintf("%s%d", key, i)
		md5Sum := fmt.Sprintf("%x", md5.Sum([]byte(hashInput)))
		md5SumPrefix := md5Sum[0:len(prefix)]
		if string(md5SumPrefix) == prefix {
			return string(md5Sum[:]), uint(i)
		}
		i++
	}
}
