package main

import (
	"fmt"
	"net"
	"strings"
)

func http200OK(conn net.Conn, body string, headers map[string]string, encoding string) {
	if headers == nil {
		headers = make(map[string]string)
	}
	if encoding == "gzip" {
		compressed, err := gzipCompress([]byte(body))
		if err == nil {
			body = string(compressed)
			headers["Content-Encoding"] = "gzip"
		}
	}
	if headers["Content-Length"] == "" {
		headers["Content-Length"] = fmt.Sprintf("%d", len(body))
	}
	if headers["Content-Type"] == "" {
		headers["Content-Type"] = "text/plain"
	}
	fmt.Printf("http200OK body: %s", body)
	var b strings.Builder

	b.WriteString("HTTP/1.1 200 OK")
	b.WriteString(CRLF)

	for key, value := range headers {
		b.WriteString(key)
		b.WriteString(": ")
		b.WriteString(value)
		b.WriteString(CRLF)
	}
	b.WriteString(CRLF)

	b.WriteString(body)
	b.WriteString(CRLF)

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

func http201Created(conn net.Conn) {
	var b strings.Builder
	b.WriteString("HTTP/1.1 201 Created")
	b.WriteString(CRLF)
	b.WriteString(CRLF)
	conn.Write([]byte(b.String()))
}

func http500InternalServerError(conn net.Conn) {
	var b strings.Builder
	b.WriteString("HTTP/1.1 500 Internal Server Error")
	b.WriteString(CRLF)
	b.WriteString(CRLF)
	b.WriteString("Internal Server Error")
	conn.Write([]byte(b.String()))
}
