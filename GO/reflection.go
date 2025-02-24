package main

import (
	"fmt"
	"reflect"
)

func main() {
	x := 10
	name := "Go Lang"
	type Book struct {
		name   string
		author string
	}
	sampleBook := Book{"Reflection in Go", "John"}
	w := reflect.TypeOf(x)
	fmt.Println(w)                          // int
	fmt.Println(reflect.TypeOf(w))          // *reflect.rtype
	fmt.Println(reflect.TypeOf(name))       // string
	fmt.Println(reflect.TypeOf(sampleBook)) // main.Book
	fmt.Println("==============================")
	var (
		str                    = "Hello, world!"
		num                    = 42
		flt                    = 3.14
		boo                    = true
		slice                  = []int{1, 2, 3}
		mymap                  = map[string]int{"foo": 1, "bar": 2}
		structure              = struct{ Name string }{Name: "John Doe"}
		interface1 interface{} = "hello"
		interface2 interface{} = &structure
	)

	fmt.Println(reflect.TypeOf(str).Kind(), reflect.TypeOf(str).Size(), reflect.ValueOf(str))     // string 16 Bytes
	fmt.Println(reflect.TypeOf(reflect.TypeOf(str).Size()), reflect.TypeOf(reflect.ValueOf(str))) // uintptr reflect.Value
	fmt.Println(reflect.TypeOf(num).Kind(), reflect.TypeOf(num).Size())                           // int
	fmt.Println(reflect.TypeOf(flt).Kind(), reflect.TypeOf(flt).Size())                           // float64
	fmt.Println(reflect.TypeOf(boo).Kind(), reflect.TypeOf(boo).Size())                           // bool
	fmt.Println(reflect.TypeOf(slice).Kind(), reflect.TypeOf(slice).Size())                       // slice
	fmt.Println(reflect.TypeOf(mymap).Kind(), reflect.TypeOf(mymap).Size())                       // map
	fmt.Println(reflect.TypeOf(structure).Kind(), reflect.TypeOf(structure).Size())               // struct
	fmt.Println(reflect.TypeOf(interface1).Kind(), reflect.TypeOf(interface1).Size())             // string
	fmt.Println(reflect.TypeOf(interface2).Kind(), reflect.TypeOf(interface2).Size())             // ptr
	fmt.Println("---------------------------")
	changeElement()
	changeValue()
}

type RPerson struct {
	Name string
	Age  int
}

func changeElement() {
	p := RPerson{Name: "John", Age: 30}
	fmt.Println("Before update:", p)

	v := reflect.ValueOf(&p)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	f := v.FieldByName("Name")
	if f.IsValid() && f.CanSet() {
		f.SetString("Jane")
	}

	fmt.Println("After update:", p)
}

type Person3 struct {
	Name string
	Age  int
}

func changeValue() {
	p := Person3{Name: "John", Age: 30}
	fmt.Println("Before update:", p)

	v := reflect.ValueOf(&p)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	f := v.FieldByName("Name")
	if f.IsValid() && f.CanSet() {
		f.SetString("Jane")
	}

	fmt.Println("After update:", p)
}
