package main

import (
	"fmt"
	"strings"
)

func http200OK(body string) string {
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

	return b.String()
}

func http404NotFound() string {
	var b strings.Builder
	b.WriteString("HTTP/1.1 404 Not Found")
	b.WriteString(CRLF)
	// b.WriteString("Content-Type: text/plain")
	// b.WriteString(CRLF)
	b.WriteString(CRLF)
	b.WriteString("Not Found")
	return b.String()
}
