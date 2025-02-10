package main

import "fmt"

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go goOne(ch1)
	go goTwo(ch2)

	select {
	case msg1 := <-ch1:
		fmt.Println(msg1)
	case msg2 := <-ch2:
		fmt.Println(msg2)
	}

	ch3 := make(chan string)
	ch4 := make(chan string)

	go goOne(ch3)
	go goTwo(ch4)

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch3:
			fmt.Println(msg1)
		case msg2 := <-ch4:
			fmt.Println(msg2)
		}
	}

}

func goOne(ch chan string) {
	ch <- "From goOne goroutine"
}

func goTwo(ch chan string) {
	ch <- "From goTwo goroutine"
}
