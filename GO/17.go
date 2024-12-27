package main

import "fmt"

func main() {
	var a = make([]int, 0, 10)
	fmt.Println(a)
	a[0] = 1
	fmt.Println(a)

	var b = new([]int)
	(*b)[0] = 1
	fmt.Println(*b)
}
