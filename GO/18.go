package main

import (
	"fmt"
	"os"
	"strings"
)

func getType(i interface{}) string {
	switch i.(type) {
	case int:
		return "int"
	case float64:
		return "float64"
	case string:
		return "string"
	default:
		return fmt.Sprintf("%T", i)
	}
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Error:::", err)
		}
	}()
	bs, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	fmt.Printf("%T\n", bs)
	data := string(bs)
	fmt.Println(data)
	fmt.Println("=================")
	words := strings.Fields(data)
	for _, word := range words {
		fmt.Println(string(word))
		fmt.Println("--")
	}

	var i any = "s" // shortcut for empty interface
	fmt.Println(i.(string))
	defer func() {
		fmt.Println(recover())
	}()
	//fmt.Println(i.(int))
	fmt.Println(getType(i))
}
