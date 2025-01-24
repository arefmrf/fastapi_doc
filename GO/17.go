package main

import "fmt"

func makeNew() {
	var a = make([]int, 1, 10)
	fmt.Println(a)
	a[0] = 1
	fmt.Println(a)

	var b = new([]int) // new create point
	/* allocates memory for the slice header (a struct with pointer, length, and capacity)
	but does not allocate memory for the underlying array.
	Without an underlying array, accessing or modifying elements results in a panic.*/
	*b = make([]int, 1, 10)
	(*b)[0] = 1
	fmt.Println(*b)
}

func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		// v and x are type T, which has the comparable
		// constraint, so we can use == here.
		if v == x {
			return i
		}
	}
	return -1
}

type List[T any] struct {
	next *List[T]
	val  T
}

func workWithList() {
	a := List[int]{
		next: &List[int]{},
		val:  1,
	}
	fmt.Println(a)
	fmt.Println(a.next)
}

func main() {
	makeNew()
	//
	//si := []int{10, 20, 15, -10}
	//fmt.Println(Index(si, 15))
	//
	//ss := []string{"foo", "bar", "baz"}
	//fmt.Println(Index(ss, "hello"))

	//workWithList()
}
