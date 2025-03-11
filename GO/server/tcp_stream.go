package main

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"time"
)

type FileServer struct{}

func (fs *FileServer) start() {
	ln, err := net.Listen("tcp", "0.0.0.0:3000")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := ln.Accept()

		if err != nil {
			fmt.Println("Error", err)
			continue
		}

		fmt.Printf("new connection: %s\n", conn.RemoteAddr().String())

		// read data from accepted connections
		go fs.readLoop(conn)
	}
}

func (fs *FileServer) readLoop(conn net.Conn) {

	// make a new buffer
	buf := new(bytes.Buffer)

	for {

		var size int64

		// get the size from connection
		err := binary.Read(conn, binary.LittleEndian, &size)
		if err != nil {
			return
		}

		// copy from connection until the end of file
		n, err := io.CopyN(buf, conn, size)
		if err != nil {
			fmt.Println("Error", err)
			continue
		}

		//fmt.Println(buf.Bytes())
		fmt.Printf("received %d bytes over the network\n", n)
	}
}

func main() {

	go func() {

		time.Sleep(4 * time.Second)

		// set your file szie
		err := sendFile(2000000)
		if err != nil {
			return
		}
	}()

	s := &FileServer{}
	s.start()
}

// client example that send a large file to server!
func sendFile(size int) error {

	file := make([]byte, size)

	// make a random file from the size provided
	_, err := io.ReadFull(rand.Reader, file)
	if err != nil {
		return err
	}

	// dial with the tcp server (you can do this is an other file)
	conn, err := net.Dial("tcp", ":3000")
	if err != nil {
		return err
	}

	// send the size of file
	err = binary.Write(conn, binary.LittleEndian, int64(size))
	if err != nil {
		return err
	}

	// copy file over the network until the end of file
	n, err := io.CopyN(conn, bytes.NewReader(file), int64(size))
	if err != nil {
		return err
	}

	fmt.Printf("written %d byte over the network\n", n)
	return nil
}

//package main
//
//import (
//	"bytes"
//	"context"
//	"crypto/rand"
//	"encoding/binary"
//	"io"
//	"log"
//	"net"
//	"os"
//	"os/signal"
//	"syscall"
//	"time"
//)
//
//// FileServer struct
//type FileServer struct {
//	listener net.Listener
//}
//
//// Start listening for incoming connections
//func (fs *FileServer) Start(ctx context.Context) {
//	ln, err := net.Listen("tcp", "0.0.0.0:3000")
//	if err != nil {
//		log.Fatalf("Failed to start server: %v", err)
//	}
//	fs.listener = ln
//	log.Println("Server started on port 3000")
//
//	go func() {
//		<-ctx.Done() // Wait for termination signal
//		log.Println("Shutting down server...")
//		_ = ln.Close()
//	}()
//
//	for {
//		conn, err := ln.Accept()
//		if err != nil {
//			select {
//			case <-ctx.Done():
//				return // Server is shutting down
//			default:
//				log.Println("Accept error:", err)
//			}
//			continue
//		}
//
//		log.Printf("New connection: %s\n", conn.RemoteAddr().String())
//		go fs.readLoop(conn)
//	}
//}
//
//// Read incoming files from clients
//func (fs *FileServer) readLoop(conn net.Conn) {
//	defer conn.Close()
//
//	// Read file size first
//	var fileSize int64
//	err := binary.Read(conn, binary.LittleEndian, &fileSize)
//	if err != nil {
//		log.Println("Failed to read file size:", err)
//		return
//	}
//	log.Printf("Expecting file of size: %d bytes\n", fileSize)
//
//	// Create a temp file for writing chunks
//	tempFile, err := os.CreateTemp("", "received_*")
//	if err != nil {
//		log.Println("Failed to create temp file:", err)
//		return
//	}
//	defer tempFile.Close()
//
//	// Read and process data in chunks (streaming)
//	buf := make([]byte, 4096) // 4KB buffer
//	var totalReceived int64
//
//	for totalReceived < fileSize {
//		n, err := conn.Read(buf)
//		if err != nil {
//			if err == io.EOF {
//				break
//			}
//			log.Println("Read error:", err)
//			return
//		}
//
//		// Write to file immediately
//		_, writeErr := tempFile.Write(buf[:n])
//		if writeErr != nil {
//			log.Println("Write error:", writeErr)
//			return
//		}
//
//		totalReceived += int64(n)
//		log.Printf("Received %d/%d bytes\n", totalReceived, fileSize)
//	}
//
//	log.Println("File transfer complete")
//}
//
//// Client function to send a random file
//func sendFile(size int64) error {
//	// Generate random data
//	file := make([]byte, size)
//	_, err := io.ReadFull(rand.Reader, file)
//	if err != nil {
//		return err
//	}
//
//	// Dial server
//	conn, err := net.Dial("tcp", ":3000")
//	if err != nil {
//		return err
//	}
//	defer conn.Close()
//
//	// Send file size first
//	err = binary.Write(conn, binary.LittleEndian, size)
//	if err != nil {
//		return err
//	}
//
//	// Send file data
//	n, err := io.CopyN(conn, bytes.NewReader(file), size)
//	if err != nil {
//		return err
//	}
//
//	log.Printf("Sent %d bytes successfully\n", n)
//	return nil
//}
//
//func main() {
//	// Create a cancelable context for graceful shutdown
//	ctx, cancel := context.WithCancel(context.Background())
//	defer cancel()
//
//	// Handle OS signals (CTRL+C, SIGTERM)
//	sigCh := make(chan os.Signal, 1)
//	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
//
//	// Start server
//	fs := &FileServer{}
//	go fs.Start(ctx)
//
//	// Start client after a delay
//	go func() {
//		time.Sleep(4 * time.Second)
//		if err := sendFile(2_000_000); err != nil {
//			log.Println("Error sending file:", err)
//		}
//	}()
//
//	// Wait for termination signal
//	<-sigCh
//	log.Println("Received shutdown signal")
//	cancel()
//}
