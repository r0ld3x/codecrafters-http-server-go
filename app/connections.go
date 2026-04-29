package main

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

var acceptedEncodings = []string{"gzip"}

func handleConnection(conn net.Conn, cfg *Config) {
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
	encoding := req.Headers["accept-encoding"]
	if encoding != "" && !slices.Contains(acceptedEncodings, encoding) {
		encoding = ""
	}

	switch {
	case path == "/":
		http200OK(conn, "Hello, World!", nil, encoding)
	case strings.HasPrefix(path, "/echo/"):
		echo := strings.TrimPrefix(path, "/echo/")
		http200OK(conn, echo, nil, encoding)
	case path == "/user-agent":
		ua := req.Headers["user-agent"]
		http200OK(conn, ua, nil, encoding)
	case strings.HasPrefix(path, "/files/"):
		fullPath := filepath.Join(cfg.Directory, strings.TrimPrefix(path, "/files/"))
		fmt.Printf("Full path: %s", fullPath)
		if req.Method == "POST" {
			err := os.WriteFile(fullPath, []byte(req.Body), 0644)
			if err != nil {
				http500InternalServerError(conn)
				return
			}
			http201Created(conn)
			return
		}
		contents, err := os.ReadFile(fullPath)
		if err != nil {
			http404NotFound(conn)
			return
		}
		header := map[string]string{
			"Content-Type": "application/octet-stream",
		}
		http200OK(conn, string(contents), header, encoding)
	default:
		http404NotFound(conn)
	}

}
