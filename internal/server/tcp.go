package server

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func StartServer() {
	listener, err := net.Listen("tcp", ":6379")

	if err != nil {
		log.Fatal("Error listening: ", err)
	}

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting conn:", err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	buffer := make([]byte, 1024)

	for {
		n, err := reader.Read(buffer)
		if err != nil {
			log.Printf("Error reading: %s", err.Error())
			conn.Close()
			return
		}

		fmt.Println(n)
	}
}
