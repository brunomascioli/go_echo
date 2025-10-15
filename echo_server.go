package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

var (
    countRequest int
)

func main() {
    listener, err := net.Listen("tcp", ":8080")
    if err != nil {
        log.Fatalf("Failed to listen on port 8080: %v", err)
    }
    defer listener.Close()
    log.Println("Echo server is listening on port 8080...")

    for {
        conn, err := listener.Accept()
        if err != nil {
            log.Printf("Failed to accept connection: %v", err)
            continue
        }

        go handleConnection(conn)
    }
}

func fibonacciRecursivo(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacciRecursivo(n-1) + fibonacciRecursivo(n-2)
}

func handleConnection(conn net.Conn) {
    log.Println(countRequest)
    countRequest++
    defer conn.Close()
    log.Printf("New connection from %s", conn.RemoteAddr())
    scanner := bufio.NewScanner(conn)
    for scanner.Scan() {
        text := scanner.Text()

        log.Printf("Received: %s", text)

        fibonacciRecursivo(30);

        _, err := fmt.Fprintf(conn, "Echo: %s\n", text)
        if err != nil {
            log.Printf("Failed to write to client: %v", err)
            return
        }
    }

    if err := scanner.Err(); err != nil {
        log.Printf("Connection error: %v", err)
    }

    log.Printf("Connection from %s closed.", conn.RemoteAddr())
}

