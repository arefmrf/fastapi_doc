package main

import (
	"fmt"
	"regexp"
)

func main() {
	sampleRegex := regexp.MustCompile("[.]")
	match := sampleRegex.Match([]byte("."))
	fmt.Println(match)

	sampleRegex = regexp.MustCompile("[abc]")
	match = sampleRegex.Match([]byte("abcd"))
	fmt.Printf("For abcd: %t\n", match)

	match = sampleRegex.Match([]byte("1abc23"))
	fmt.Printf("For 1abc23: %t\n", match)

	match = sampleRegex.Match([]byte("abc"))
	fmt.Printf("For abc: %t\n", match)
}
