package main

import (
	"bufio"
	"log"
	"net"
)

const (
	network = "tcp"
	port    = ":8080"
)

func main() {
	tcpAddr, _ := net.ResolveTCPAddr(network, port)

	conn, err := net.DialTCP(network, nil, tcpAddr)
	if err != nil {
		log.Fatalln("Connection to server error:", err)
	}
	defer conn.Close()

	message, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		log.Fatalln("Data reading error", err)
	}

	if message == "OK\n" {
		log.Printf("Success: %s", message)
	} else {
		log.Printf("Not success: %s", message)
	}
}
