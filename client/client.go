package main

import (
	"log"
	"net"
	"time"
)

func sendMessage(message []byte, conn net.Conn) {
    _, err := conn.Write(message)
    if err != nil {
        log.Fatal("Erro ao enviar mensagem:", err)
    }
}

func getResponse(buffer []byte, conn net.Conn) int {
    n, err := conn.Read(buffer)
    if err != nil {
        log.Fatal("Erro ao ler resposta:", err)
    }
    return n
}

func main() {

    for i := 0; i < 10000; i++ {
        conn, err := net.Dial("tcp", "localhost:9000")
        buffer := make([]byte, 1024)
        start := time.Now()
        if err != nil {
            log.Fatal("Erro ao conectar ao servidor:", err)
        }
        
        sendMessage([]byte("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa\n"), conn)
        
        n := getResponse(buffer, conn)
        response := string(buffer[:n])
        
        log.Printf("Recebido do servidor: %s\n", response)
        log.Println("Latência: ", time.Since(start))
        conn.Close()
    }
}