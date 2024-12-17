package main

import "fmt"

func qPointer() {
	i, j := 42, 2701

	p := &i
	fmt.Println("*p: ", *p)
	*p = 21
	fmt.Println("i: ", i)
	fmt.Println("*p: ", *p)

	p = &j
	*p = *p / 37
	fmt.Println(j)
}

func qStruct() {
	type vertex struct {
		X int
		Y int
		Z []int
		Q func()
	}
	fmt.Println(vertex{1, 2, []int{1, 2, 3}, qPointer})
	q := vertex{1, 2, []int{1, 2, 3}, qPointer}
	q.Q()
}

func main() {
	//qPointer()
	qStruct()
}
