//package main
//
//import (
//	"fmt"
//	"log"
//	"net"
//	"os"
//	"strings"
//)
//
//
//func main() {
//	conn, err := net.ListenUDP("udp", &net.UDPAddr{
//		Port: 3000,
//		IP:   net.ParseIP("0.0.0.0"),
//	})
//	if err != nil {
//		panic(err)
//	}
//
//	defer conn.Close()
//	fmt.Printf("server listening %s\n", conn.LocalAddr().String())
//
//	for {
//		message := make([]byte, 20)
//		readLen, remote, err := conn.ReadFromUDP(message[:])
//		if err != nil {
//			panic(err)
//		}
//
//		data := strings.TrimSpace(string(message[:readLen]))
//		fmt.Printf("received: %s from %s\n", data, remote)
//	}
//}

package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	// Create a log file to record errors
	logFile, err := os.OpenFile("server.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("Error opening log file: %v\n", err)
		return
	}
	defer logFile.Close()
	log.SetOutput(logFile)

	// Specify the address and port to listen for UDP packets
	addr := &net.UDPAddr{
		Port: 3000,
		IP:   net.ParseIP("0.0.0.0"),
	}

	// Establish a UDP connection
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatalf("Error setting up UDP listener: %v\n", err)
		return
	}
	defer conn.Close()
	fmt.Printf("Server listening on %s\n", conn.LocalAddr().String())

	for {
		// Read data from the UDP connection
		buf := make([]byte, 1024)
		_, remoteAddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			log.Printf("Error reading from UDP connection: %v\n", err)
			continue
		}

		// Convert received data to a string and trim any extra spaces
		data := strings.TrimSpace(string(buf))
		fmt.Printf("Received: %s from %s\n", data, remoteAddr)
	}
}
