package main

import (
	"fmt"
	"strings"
)

func shout(ping <-chan string, pong chan<- string) {
	for {
		s, ok := <-ping
		if !ok {
			println("channel is closed")
		}
		pong <- fmt.Sprintf("%s!!!", strings.ToUpper(s))
	}
}

func main() {
	ping := make(chan string)
	pong := make(chan string)

	go shout(ping, pong)
	fmt.Println("Type something and press ENTER (send Q to quit)")

	for {
		fmt.Println("->")

		var userString string
		_, _ = fmt.Scanln(&userString)
		if strings.ToLower(userString) == "q" {
			break
		}
		ping <- userString
		response := <-pong
		fmt.Println("Response:", response)
	}
	fmt.Println("Closing channel")
	close(ping)
	close(pong)
}
