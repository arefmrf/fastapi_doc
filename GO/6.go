package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func xSlice() {
	var a []int = []int{1, 2, 3, 4, 5}
	fmt.Println(a)
	var b []int = make([]int, 5)
	fmt.Println(b)
	c := []int{1, 2, 3, 4, 5}
	fmt.Println(c)
	fmt.Println("len:", len(c), "cap: ", cap(c))

	d := c[0:2]
	fmt.Println(d)
	c[0] = 7
	fmt.Println(d)

	var e [3]int = [3]int{1, 2, 3}
	f := e[:2]
	g := append(f, 4, 5, 6, 7)
	/* Go creates a new underlying array with enough capacity to accommodate the new elements.*/
	fmt.Println(g)
	h := append(g, f...)
	fmt.Println(h)
	fmt.Println("#############################################")
	i := []int{1, 2, 3, 4, 5}
	fmt.Println("Before removal:", i)
	// Remove the element at index 2
	index := 2
	copy(i[index:], i[index+1:]) // Shift elements left
	fmt.Println("After removal:", i)
	i = i[:len(i)-1] // Resize slice
	fmt.Println("After removal:", i)

}

func shareUnderlyingArray(slice1, slice2 []int) bool {
	ptr1 := (*reflect.SliceHeader)(unsafe.Pointer(&slice1)).Data
	ptr2 := (*reflect.SliceHeader)(unsafe.Pointer(&slice2)).Data
	return ptr1 == ptr2
}

func main() {
	xSlice()
}
