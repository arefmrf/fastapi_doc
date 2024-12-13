package main

import "fmt"

func qwe(x, y int) (a, b int) {
	a = x - y
	b = x + y
	return
}

func sumNumbers(numbers ...int) int {
	sum := 0
	for _, value := range numbers {
		sum += value
	}
	return sum
}

func main() {
	fmt.Println(qwe(1, 2))
	fmt.Println(sumNumbers())
	fmt.Println(sumNumbers(1, 2, 3))
}
