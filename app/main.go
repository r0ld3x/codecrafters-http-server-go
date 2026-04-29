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
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading from connection: ", err.Error())
		os.Exit(1)
	}

	read := string(buf)

	data, err := parseRequest(read[:n])
	if err != nil {
		fmt.Println("Error parsing request: ", err.Error())
		conn.Close()
		return
	}

	path := data.Path
	var response string
	if path == "/" {
		response = http200OK("Hello, World!")
	} else if after, ok := strings.CutPrefix(path, "/echo/"); ok {
		response = http200OK(after)
	} else if path == "/user-agent" {
		response = http200OK(data.Headers["user-agent"])
	} else {
		response = http404NotFound()
	}
	conn.Write([]byte(response))

}
