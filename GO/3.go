package main

import (
	"fmt"
	"strconv"
)

func xx() {
	outterVar := 1
	{
		innerVar := 1
		outterVar = 2
		fmt.Println(innerVar)
	}
	fmt.Println(outterVar)

}

func getInput() {
	var name string
	var age int
	var isMale bool
	fmt.Print("Enter name: ")
	count, err := fmt.Scanf("%s", &name)
	fmt.Print("count, error: ", count, err, "\n")
	fmt.Print("Enter age and is male: ")
	count, err = fmt.Scanf("%d %t", &age, &isMale)
	fmt.Print("count, error: ", count, err, "\n")
	fmt.Println("Hello ", name, "age: ", age, "is male: ", isMale)
}

func castType() {
	var q int = 10
	var w = float64(q)
	var e = strconv.Itoa(q)
	r, err := strconv.Atoi(e)
	fmt.Printf("%v type is %T\n", q, q)
	fmt.Printf("%v type is %T\n", w, w)
	fmt.Printf("%q type is %T\n", e, e)
	fmt.Printf("%v type is %T\n", r, r)
	fmt.Printf("error: %v type is %T\n", err, err)

	fmt.Print("============================\n")
	r, err = strconv.Atoi("200abc")
	fmt.Printf("%v type is %T\n", r, r)
	fmt.Printf("error: %v type is %T\n", err, err)
}

func main() {
	//xx()
	//getInput()
	castType()
}
