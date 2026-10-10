package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

// Ensures gofmt doesn't remove the "net" and "os" imports above (feel free to remove this!)
var (
	_ = net.Listen
	_ = os.Exit
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	// TODO:

	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	defer l.Close()

	conn, err := l.Accept()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	buffer := make([]byte, 1024)

	n, err := conn.Read(buffer)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	request := buffer[:n]

	reqLine := strings.Split(string(request), "\r\n")

	parts := reqLine[0]

	path := strings.Split(parts, " ")

	var response string
	if path[1] == "/" {
		response = "HTTP/1.1 200 OK\r\n\r\n"
	} else {
		response = "HTTP:/1.1 404 Not Found \r\n\r\n"
	}

	_, err = conn.Write([]byte(response))
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	os.Exit(0)
}
