package server

import (
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/ritayyshh/RedisServer/handler"
	"github.com/ritayyshh/RedisServer/resp"
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

	for {
		respReader := resp.NewResp(conn)
		value, err := respReader.Read()
		if err != nil {
			log.Printf("Error reading: %s", err.Error())
			conn.Close()
			return
		}

		if value.Typ != "array" {
			fmt.Println("Invalid request, expected array")
			continue
		}

		if len(value.Array) == 0 {
			fmt.Println("Invalid request, expected array length > 0")
			continue
		}

		command := strings.ToUpper(value.Array[0].Bulk)
		args := value.Array[1:]

		writer := resp.NewWriter(conn)

		handler, ok := handler.Handlers[command]
		if !ok {
			fmt.Println("Invalid command: ", command)
			writer.Write(resp.Value{Typ: "string", Str: ""})
			continue
		}

		result := handler(args)
		writer.Write(result)
	}
}
