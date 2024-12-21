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
	q.X = 10
	fmt.Println(q.X)
	fmt.Println("=============================")
	var w *vertex
	w = &q
	w.Y = 12
	fmt.Println(w.Y)
	fmt.Println(q.Y)

	e := &vertex{X: 1}
	r := *e
	fmt.Println(e)
	fmt.Println(r)
}

func Pic(dx, dy int) [][]uint8 {
	a := make([][]uint8, dy)
	for y := 0; y < dy; y++ {
		a[y] = make([]uint8, dx)
		for x := 0; x < dx; x++ {
			a[y][x] = uint8((x + y) / 2)
		}
	}
	return a
}

func Pic2(dx, dy int) [][]uint8 {
	// Initialize the outer slice
	var a [][]uint8
	for y := 0; y < dy; y++ {
		var row []uint8
		for x := 0; x < dx; x++ {
			// Append each value to the row
			row = append(row, uint8((x+y)/2))
		}
		// Append the row to the outer slice
		a = append(a, row)
	}
	return a
}

func main() {
	//qPointer()
	//qStruct()
	fmt.Println(Pic(10, 5))
	fmt.Println(Pic2(10, 5))
}
