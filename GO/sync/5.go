package main

import (
	"fmt"
	"sync"
	"time"
)

var sharedResource = make(map[string]interface{})

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	locker := sync.Mutex{}
	condition := sync.NewCond(&locker)

	go waitForResourceUpdate(&wg, condition, "rsc1")
	go waitForResourceUpdate(&wg, condition, "rsc2")
	time.Sleep(2 * time.Second)
	// this one writes changes to sharedResource
	condition.L.Lock()
	sharedResource["rsc1"] = "a string"
	sharedResource["rsc2"] = 123456
	condition.Broadcast()
	condition.L.Unlock()

	wg.Wait()
}

// waitForResourceUpdate waits for a signal that a resource changed and prints it.
func waitForResourceUpdate(wg *sync.WaitGroup, cond *sync.Cond, key string) {
	defer wg.Done()
	cond.L.Lock()
	for len(sharedResource) == 0 {
		cond.Wait()
	}
	fmt.Println("Resource", key, ":", sharedResource[key])
	cond.L.Unlock()
}
