package main

import "fmt"

func first() string {
	defer fmt.Println("world")
	return "Hello"
}

func main() {
	// defer -->last-in-first-out order
	defer fmt.Println("world2")
	a := first()
	fmt.Println(a)
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered. Error:\n", r)
		}
	}()
	defer fmt.Println("world3")

	fmt.Println("done")
}
