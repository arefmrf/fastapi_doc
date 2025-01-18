package main

import (
	"fmt"
	"sync"
	"time"
)

var count int

func main() {
	mu := new(sync.Mutex)

	go increment(mu)
	go increment(mu)
	go increment(mu)
	go increment(mu)

	time.Sleep(time.Second)
}

func increment(mu *sync.Mutex) {
	mu.Lock()
	defer mu.Unlock()
	count++
	fmt.Printf("Incrementing: %d\n", count)
}
