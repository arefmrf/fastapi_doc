package main

import (
	"bufio"
	"fmt"
	"os"
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
	fmt.Println("==============================1")
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
	fmt.Println("---------------------------2")
	changeElement()
	fmt.Println("---------------------------3")
	changeValue()
	fmt.Println("---------------------------4")
	analyzeStruct()
	fmt.Println("---------------------------5")
	//analyzeMethod()
	fmt.Println("---------------------------6")
	tagging()
}

type RPerson struct {
	Name string
	Age  int
}

type Person4 struct {
	Name    string
	Age     int
	Address string
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

func changeValue() {
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

func analyzeStruct() {
	p := Person4{Name: "John", Age: 30, Address: "123 Main St."}

	v := reflect.ValueOf(p)
	if v.Kind() == reflect.Ptr {
		// wont come in because v.kind is not reflect.Ptr
		v = v.Elem()
	}

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fmt.Printf("Field %d: %s = %v\n", i, v.Type().Field(i).Name, field.Interface())
	}
}

// #######################################################
// #######################################################

type NativeCommandEngine struct{}

func (nse NativeCommandEngine) Method1() {
	fmt.Println("INFO: Method1 executed!")
}
func (nse NativeCommandEngine) Method2() {
	fmt.Println("INFO: Method2 executed!")
}
func (nse NativeCommandEngine) callMethodByName(methodName string) {
	method := reflect.ValueOf(nse).MethodByName(methodName)
	if !method.IsValid() {
		fmt.Println("ERROR: \"" + methodName + "\" is not implemented")
		return
	}
	method.Call(nil)
}
func (nse NativeCommandEngine) ShowCommands() {
	val := reflect.TypeOf(nse)
	for i := 0; i < val.NumMethod(); i++ {
		fmt.Println(val.Method(i).Name)
	}
}
func analyzeMethod() {
	nse := NativeCommandEngine{}
	fmt.Println("A simple Shell v1.0.0")
	fmt.Println("Supported commands:")
	nse.ShowCommands()
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("$ ")
	for scanner.Scan() {
		nse.callMethodByName(scanner.Text())
		fmt.Print("$ ")
	}
}

// #######################################################
// #######################################################

type tagPerson struct {
	Name string `customtag:"myname"`
	Age  int    `customtag:"myage"`
}

func tagging() {
	p := tagPerson{"John", 30}

	t := reflect.TypeOf(p)
	v := reflect.ValueOf(&p)
	/*
		ValueOf provides access to runtime values stored in the struct.
		Value is useful for retrieving actual data, it does not provide field tags.
	*/
	if v.Kind() == reflect.Ptr {
		v = v.Elem() // Dereference to access struct fields
	}

	p.Name = "sss"
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)

		tag := field.Tag.Get("customtag")

		fmt.Printf("Field: %s, Value: %v, Tag: %s\n", field.Name, value.Interface(), tag)
	}
}
