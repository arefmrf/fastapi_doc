package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	ch2 := make(chan int, 3)

	go send(ch)
	go send(ch2)
	go send(ch2)
	go send(ch2)
	time.Sleep(time.Second * 1)
	go receive(ch)
	go receive(ch2)

	time.Sleep(time.Second * 1)
	fmt.Println("-------------------------")
	fmt.Println("-------------------------")
	ch3 := make(chan int, 4)
	send(ch3)
	send(ch3)
	go sum(ch3)
	send(ch3)
	close(ch3) // if we don't close channel:  range ch won't stop
	time.Sleep(time.Second * 1)
}

func sum(ch chan int) {
	sum := 0
	for val := range ch {
		sum += val
	}
	fmt.Printf("Sum: %d\n", sum)
}

func send(ch chan<- int) {
	fmt.Println("Sending value to channel with capacity: ", cap(ch))
	ch <- cap(ch)
}

func receive(ch <-chan int) {
	val := <-ch
	fmt.Printf("Value Received=%d in receive function. remain length: %d\n", val, len(ch))
}
