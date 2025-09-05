package main

import (
	"fmt"
	"sync"
)

// Singleton struct (unexported)
type singleton struct {
	data string
}

var instance *singleton
var once sync.Once

// GetInstance returns the singleton instance
func GetInstance() *singleton {
	once.Do(func() {
		instance = &singleton{data: "I am a Singleton"}
	})
	return instance
}

func CreateSingleton() {
	s1 := GetInstance()
	s2 := GetInstance()

	fmt.Println(s1.data)
	fmt.Println(s2.data)

	// Checking if both instances are the same
	fmt.Println("Are both instances the same?", s1 == s2)
}
