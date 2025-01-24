package main

import (
	"fmt"
	"sync"
)

func printIt(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(s)
}

func main() {
	/* words := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i"}
	//wg := sync.WaitGroup  why its not working? why var x struct works?
	var wg sync.WaitGroup
	wg.Add(len(words))
	for i, word := range words {
		go printIt(fmt.Sprintf("%d: %s", i, word), &wg)
	}
	wg.Wait() */
}
