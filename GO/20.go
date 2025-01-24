package main

import (
	"fmt"
	"https://golang.org/x/exp/constraints"
)

// ~ means even for types that have int as their underlying type
func GenericAdd[T ~int | float64](a, b T) T {
	return a + b
}

type MyInt int

type numeric interface {
	~int | float64
}

type numeric2 interface {
	constraints.Integer | constants.Float
}

func iGenericAdd[T numeric](a, b T) T {
	return a + b
}

func iGenericAdd2[T numeric2](a, b T) T {
	return a + b
}

func main() {
	fmt.Println(GenericAdd[int](1, 2))
	fmt.Println(GenericAdd[float64](1.1, 2.2))
	q := MyInt(1)
	w := MyInt(2)
	fmt.Println(GenericAdd(q, w))
	fmt.Println(GenericAdd[MyInt](q, w))

	fmt.Println(iGenericAdd[MyInt](q, w))

	fmt.Println(iGenericAdd2(1, 2))
}
