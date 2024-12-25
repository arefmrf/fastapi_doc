package main

import (
	"fmt"
)

type Vertex2 struct {
	X, Y float64
}

func (v *Vertex2) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func Scale2(v *Vertex2, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func typeCheck() {
	i := interface{}("hello")

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64) // panic
	fmt.Println(f)
}

func typeCheck2(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	case struct{ name string }:
		fmt.Printf("%q type is %T\n", v, v)
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	v := Vertex2{3, 4}
	v.Scale(10)
	Scale2(&v, 10)
	fmt.Println("--> ", v.X, v.Y)

	//typeCheck()
	typeCheck2(1)
	typeCheck2("s")
	typeCheck2(struct {
		name string
	}{"aref"})
}
