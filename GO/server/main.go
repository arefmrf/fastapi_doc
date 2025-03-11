//package main
//
//import (
//	"fmt"
//	"log"
//	"net"
//)
//
//type Message struct {
//	from    string
//	payload []byte
//}
//
//type Server struct {
//	listenAddr string
//	// listener
//	ln    net.Listener
//	msgch chan Message
//}
//
//func newServer(listenAddr string) *Server {
//	return &Server{
//		listenAddr: listenAddr,
//		msgch:      make(chan Message, 10),
//	}
//}
//
//func (s *Server) start() error {
//	ln, err := net.Listen("tcp", s.listenAddr)
//	if err != nil {
//		return err
//	}
//	s.ln = ln
//	go s.acceptLoop()
//
//	return nil
//}
//
//func (s *Server) stop() {
//	if s.ln != nil {
//		err := s.ln.Close()
//		if err != nil {
//			log.Fatal(err)
//			return
//		}
//	}
//}
//
//func (s *Server) acceptLoop() {
//	for {
//		conn, err := s.ln.Accept()
//		if err != nil {
//			fmt.Println("accept error:", err)
//			continue
//		}
//		go s.readLoop(conn)
//	}
//}
//
//func (s *Server) readLoop(conn net.Conn) {
//	defer func(conn net.Conn) {
//		err := conn.Close()
//		if err != nil {
//			log.Fatal(err)
//		}
//	}(conn)
//	buf := make([]byte, 2048)
//
//	for {
//		n, err := conn.Read(buf)
//		if err != nil {
//			fmt.Println("read error:", err)
//			continue
//		}
//
//		s.msgch <- Message{
//			from:    conn.RemoteAddr().String(),
//			payload: buf[:n],
//		}
//
//		_, err = conn.Write([]byte("your message recived!\n"))
//		if err != nil {
//			return
//		}
//	}
//}
//
//func main() {
//	server := newServer(":3000")
//
//	//start the server
//	if err := server.start(); err != nil {
//		log.Fatal(err)
//	}
//
//	go func() {
//		for msg := range server.msgch {
//			fmt.Printf("recived new from connection(%s): %s\n", msg.from, msg.payload)
//		}
//	}()
//
//	select {}
//
//}

package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type Message struct {
	from    string
	payload string
}

type Server struct {
	listenAddr string
	ln         net.Listener
	msgch      chan Message
	wg         sync.WaitGroup
	ctx        context.Context
	cancel     context.CancelFunc
}

func newServer(listenAddr string) *Server {
	ctx, cancel := context.WithCancel(context.Background())
	return &Server{
		listenAddr: listenAddr,
		msgch:      make(chan Message, 100),
		ctx:        ctx,
		cancel:     cancel,
	}
}

func (s *Server) start() error {
	ln, err := net.Listen("tcp", s.listenAddr)
	if err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}
	s.ln = ln

	log.Printf("Server started on %s\n", s.listenAddr)

	s.wg.Add(1)
	go s.acceptLoop()

	return nil
}

func (s *Server) stop() {
	log.Println("Shutting down server...")

	// Stop accepting new connections
	if s.ln != nil {
		_ = s.ln.Close()
	}

	// Cancel all goroutines
	s.cancel()

	// Wait for all goroutines to finish
	s.wg.Wait()

	close(s.msgch)

	log.Println("Server shutdown complete.")
}

func (s *Server) acceptLoop() {
	defer s.wg.Done()

	for {
		conn, err := s.ln.Accept()
		if err != nil {
			select {
			case <-s.ctx.Done():
				return // Server is shutting down
			default:
				log.Printf("Accept error: %v\n", err)
			}
			continue
		}

		s.wg.Add(1)
		go s.handleConnection(conn)
	}
}

func (s *Server) handleConnection(conn net.Conn) {
	defer s.wg.Done()
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			log.Printf(err.Error())
		}
	}(conn)

	log.Printf("New connection from %s\n", conn.RemoteAddr())

	reader := bufio.NewReader(conn)
	for {
		// Read message line by line
		msg, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("Read error from %s: %v\n", conn.RemoteAddr(), err)
			return
		}

		s.msgch <- Message{
			from:    conn.RemoteAddr().String(),
			payload: msg,
		}

		_, _ = conn.Write([]byte("Your message received!\n"))
	}
}

func main() {
	server := newServer(":3000")

	// Start the server
	if err := server.start(); err != nil {
		log.Fatal(err)
	}

	// Start message handler
	go func() {
		for msg := range server.msgch {
			log.Printf("Received from %s: %s", msg.from, msg.payload)
		}
	}()

	// Handle graceful shutdown
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	<-sigCh
	server.stop()
	//go func() {
	//	<-sigCh
	//	server.stop()
	//}()
}
