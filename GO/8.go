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

func recursive1(n int) int {
	if n == 1 {
		return 1
	}
	return n * recursive1(n-1)
}

func anonymousFunc() {
	x := func(a, b int) int {
		return a + b
	}
	fmt.Println(x(2, 3))
	y := func(a, b int) int {
		return a + b
	}(2, 3)
	fmt.Println(y)
}

func main() {
	fmt.Println(qwe(1, 2))
	fmt.Println(sumNumbers())
	fmt.Println(sumNumbers(1, 2, 3))
	fmt.Println(recursive1(5))
	anonymousFunc()
}
