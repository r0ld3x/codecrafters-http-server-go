package main

import (
	"fmt"
	"net"
	"strings"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)

	if err != nil {
		fmt.Println("Failed to bind:", err)
		return
	}
	read := string(buf)
	req, err := parseRequest(read[:n])
	if err != nil {
		fmt.Println("Error parsing request: ", err.Error())
		return
	}

	path := req.Path

	if req.Method != "GET" {
		http404NotFound(conn)
		return
	}

	switch {
	case path == "/":
		http200OK(conn, "Hello, World!")
	case strings.HasPrefix(path, "/echo/"):
		echo := strings.TrimPrefix(path, "/echo/")
		http200OK(conn, echo)
	case path == "/user-agent":
		ua := req.Headers["user-agent"]
		http200OK(conn, ua)
	default:
		http404NotFound(conn)
	}

}
