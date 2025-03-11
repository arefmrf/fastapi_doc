package main

import "fmt"

func xAdd(x, y int, z ...int) (a, b int) {
	a = x + y
	b = x - y
	for _, z := range z {
		a += z
		b -= z
	}
	return
}

func main() {
	a := 7
	fmt.Println(a)
	b := &a
	fmt.Println(*b)
	q, w := xAdd(40, 2, 1, 2, 3, 4)
	fmt.Println(q, w)
	q, w = xAdd(40, 2)
	fmt.Println(q, w)
}
