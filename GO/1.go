package main

import (
	"fmt"
	"math/rand"
)

// func add(x int, y int) int {
func add(x, y int) int {
	return x * y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(rand.Intn(10))

	fmt.Println(add(42, 13))

	a, b := swap("hello", "world")
	fmt.Println(a, b)

	var c string
	var d string
	c, d = swap("hello", "world")
	fmt.Println(c, d)

	fmt.Println(split(17))
}
