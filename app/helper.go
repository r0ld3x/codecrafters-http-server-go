package main

import (
	"fmt"
	"net"
	"strings"
)

func http200OK(conn net.Conn, body string) {
	fmt.Printf("http200OK body: %s", body)
	var b strings.Builder

	b.WriteString("HTTP/1.1 200 OK")
	b.WriteString(CRLF)

	b.WriteString("Content-Type: text/plain")
	b.WriteString(CRLF)

	b.WriteString(fmt.Sprintf("Content-Length: %d", len(body)))
	b.WriteString(CRLF)
	b.WriteString(CRLF)

	b.WriteString(body)

	conn.Write([]byte(b.String()))
}

func http404NotFound(conn net.Conn) {
	var b strings.Builder
	b.WriteString("HTTP/1.1 404 Not Found")
	b.WriteString(CRLF)
	// b.WriteString("Content-Type: text/plain")
	// b.WriteString(CRLF)
	b.WriteString(CRLF)
	b.WriteString("Not Found")
	conn.Write([]byte(b.String()))
}
