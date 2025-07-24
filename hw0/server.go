package main

import (
	"log"
	"net"
)

const (
	network = "tcp"
	port    = ":8080"
)

func main() {
	tcpAddr, _ := net.ResolveTCPAddr(network, port)

	listener, err := net.ListenTCP(network, tcpAddr)
	if err != nil {
		log.Fatalln("Starting server error:", err)
	}
	defer listener.Close()

	log.Printf("Listening port %s\n", port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Connection error:", err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	log.Printf("Connected client %s\n", conn.RemoteAddr().String())

	_, err := conn.Write([]byte("OK\n"))

	if err != nil {
		log.Println("Data sending error:", err)
	}
}
