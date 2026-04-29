package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
	"strings"
)

var acceptedEncodings = []string{"gzip"}

func handleConnection(conn net.Conn, cfg *Config) {
	defer conn.Close()

	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)

		if err != nil {
			fmt.Println("Failed to bind:", err)
			return
		}
		read := string(buf)
		req, err := parseRequest(read[:n])
		if err != nil {
			if err != io.EOF {
				conn.Write([]byte("HTTP/1.1 400 Bad Request\r\nConnection: close\r\n\r\n"))
			}
			fmt.Println("Error parsing request: ", err.Error())
			return
		}

		path := req.Path
		encoding := getValidEncoding(req.Headers["accept-encoding"])
		closeAfter := strings.EqualFold(req.Headers["Connection"], "close")
		switch {
		case path == "/":
			if closeAfter {
				header := map[string]string{
					"Connection": "close",
				}
				http200OK(conn, "Hello, World!", header, encoding)
				return
			}
			http200OK(conn, "Hello, World!", nil, encoding)
		case strings.HasPrefix(path, "/echo/"):
			if closeAfter {
				header := map[string]string{
					"Connection": "close",
				}
				echo := strings.TrimPrefix(path, "/echo/")
				http200OK(conn, echo, header, encoding)
				return
			}
			echo := strings.TrimPrefix(path, "/echo/")
			http200OK(conn, echo, nil, encoding)
		case path == "/user-agent":
			if closeAfter {
				header := map[string]string{
					"Connection": "close",
				}
				ua := req.Headers["user-agent"]
				http200OK(conn, ua, header, encoding)
				return
			}
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
			if closeAfter {
				return
			}
		default:
			http404NotFound(conn)
			if closeAfter {
				return
			}
		}
	}
}
