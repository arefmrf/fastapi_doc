package main

import (
	"fmt"
	"math/rand"
	"time"
)

func fill(ch chan<- int, max int, done chan<- struct{}, id int) {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	for i := 0; i < cap(ch); i++ {
		fmt.Println("adding to channel. id: ", id)
		ch <- r.Intn(max)
	}
	done <- struct{}{}
}

func main() {
	numChan := make(chan int, 5)
	done := make(chan struct{}, 2)
	fmt.Println("======== ", len(numChan), cap(numChan))
	go fill(numChan, 100, done, 1)
	go fill(numChan, 100, done, 2)

	go func() {
		<-done // Wait for the first goroutine to finish
		<-done // Wait for the second goroutine to finish
		fmt.Println("***closing***")
		close(numChan)
	}()

	for num := range numChan {
		fmt.Println(num)
	}
	fmt.Println("======== ", len(numChan), cap(numChan))
}
