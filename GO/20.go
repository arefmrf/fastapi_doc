package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	//go start()
	fmt.Println("Started")
	for i := 0; i < 20000000; i++ {
		go execute(i)
	}
	//time.Sleep(6 * time.Microsecond)
	fmt.Println(runtime.NumGoroutine())
	time.Sleep(4 * time.Second)
	fmt.Println("Finished")
	fmt.Println(runtime.NumCPU())
}

//func start() {
//	fmt.Println("In Goroutine")
//}

func execute(id int) {
	time.Sleep(3 * time.Second)
	//fmt.Printf("id: %d\n", id)
}
