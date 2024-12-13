package main

import "fmt"

func qq() {
	//fmt.Printf("%.200f", 1.1-0.3)
	fmt.Printf("%.5f\n", 1.1-0.3)
	fmt.Println(float64(10) / float64(3))
	fmt.Println(float64(10) / float64(3))
	fmt.Println(10 / 3)
	fmt.Println(10 % 3)

	var i float64 = 1.1
	i++
	var j float64
	j += i
	fmt.Println(j)
}

func ww() {
	var i string = "xyz"
	if i == "x" {
		fmt.Println("x")
	} else if i == "y" {
		fmt.Println("y")
	} else {
		fmt.Println("z")
	}
}

func www() {
	var i int64 = 100
	switch i {
	case int64(100):
		fmt.Println("x")
	case int64(101), int64(102):
		fmt.Println("y")
	default:
		fmt.Println("z")
	}

	switch {
	case i <= 100:
		fmt.Println("i <= 100")
		fallthrough
	case i <= 200:
		fmt.Println("i <= 200")
		fallthrough
	default:
		fmt.Println("i default")
	}
}

func main() {
	qq()
	ww()
	www()
}
