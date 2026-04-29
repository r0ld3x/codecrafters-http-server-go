package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

const CRLF = "\r\n"

// Ensures gofmt doesn't remove the "net" and "os" imports above (feel free to remove this!)
var _ = net.Listen
var _ = os.Exit

func main() {

	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}

	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}

	buf := make([]byte, 1024)
	_, err = conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading from connection: ", err.Error())
		os.Exit(1)
	}

	read := string(buf)

	lines := strings.Split(read, CRLF)
	path := strings.Split(lines[0], " ")[1]
	var response string
	if path == "/" {
		response = "HTTP/1.1 200 OK" + CRLF + CRLF + "Hello, World!"
	} else {
		response = "HTTP/1.1 404 Not Found" + CRLF + CRLF + "Not Found"
	}
	conn.Write([]byte(response))

}
