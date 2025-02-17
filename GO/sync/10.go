package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) == 4 {
		fmt.Println("First Argument:", args[1])
		fmt.Println("Second Argument:", args[2])
		fmt.Println("Third Argument:", args[3])
	}

	userName := flag.String("username", "", "client username")

	printResult := flag.Bool("print", false, "print value(boolian)")

	repeatCount := flag.Int("count", 1, "Repeat count number")

	flag.Parse()

	fmt.Println("User name: ", *userName)
	fmt.Println("Print result: ", *printResult)
	fmt.Println("Repeat count: ", *repeatCount)
}

// go run main.go -username=John -print=true -count=5
