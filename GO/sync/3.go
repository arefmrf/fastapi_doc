package main

import (
	"fmt"
	"sync"
)

type Singleton struct {
	data string
}

var instance *Singleton
var once sync.Once

func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{data: "some data"}
	})
	return instance
}

func main() {
	s1 := GetInstance()
	s2 := GetInstance()
	if s1 == s2 {
		fmt.Println("Same instance")
	} else {
		fmt.Println("Different instances")
	}
}
