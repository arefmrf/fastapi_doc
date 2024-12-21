package main

import "fmt"

func forLoop() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	fmt.Println("======================")
	y := 0
	for y < 5 {
		fmt.Println(y)
		y++
	}
	fmt.Println("======================")
	j := 0
	for {
		j++
		if j == 1 {
			continue
		}
		if j == 5 {
			break
		}
		fmt.Println(j)
	}
}

func xArray() {
	var i [5]int
	fmt.Println(i)

	var j [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(j)

	z := [5]int{1, 2, 3, 4, 5}
	fmt.Println(z)

	y := [...]int{1, 2, 3, 4, 5}
	fmt.Println(y)
	fmt.Println("length of y is: ", len(y))
	fmt.Println("index 0 of y is: ", y[0])
	y[0] = 6
	fmt.Println("index 0 of y is: ", y[0])
	fmt.Println("======================")
	for index, value := range y {
		fmt.Println("index: ", index, "value: ", value)
	}
	fmt.Println("======================")
	// return index
	for ff := range y {
		fmt.Println("index: ", ff, "value: ", y[ff])
	}
	fmt.Println("======================")
	//var q [3][2]int = [3][2]int{{0, 1}, {2, 3}, {4, 5}}
	q := [3][2]int{{0, 1}, {2, 3}, {4, 5}}
	fmt.Println(q)
	fmt.Println(q[0])
	fmt.Println(q[0][0])
}
func main() {
	//forLoop()
	xArray()
}
