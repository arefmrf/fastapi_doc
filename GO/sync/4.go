package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	pool := sync.Pool{
		New: func() interface{} {
			fmt.Println("Creating a new object")
			return "New Object"
		},
	}

	pool.Put("Preloaded Object")

	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			obj := pool.Get().(string)
			fmt.Printf("Goroutine %d got: %s\n", id, obj)
			time.Sleep(100 * time.Millisecond)
			pool.Put(fmt.Sprintf("Goroutine %d's Object", id))
		}(i)
	}

	wg.Wait()
	time.Sleep(1000 * time.Millisecond)
	for i := 0; i < 5; i++ {
		obj := pool.Get().(string)
		fmt.Println("Main goroutine got:", obj)
	}
}
