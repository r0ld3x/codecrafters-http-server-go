package main

import (
	"errors"
	"strings"
)

type Request struct {
	Method  string
	Path    string
	Proto   string
	Headers map[string]string
	Body    string
}

func parseRequest(raw string) (*Request, error) {
	// 1. Split headers and body
	parts := strings.SplitN(raw, CRLF+CRLF, 2)
	if len(parts) < 1 {
		return nil, errors.New("invalid request")
	}

	headerPart := parts[0]
	body := ""
	if len(parts) == 2 {
		body = parts[1]
	}

	// 2. Split header lines
	lines := strings.Split(headerPart, CRLF)
	if len(lines) == 0 {
		return nil, errors.New("empty request")
	}

	// 3. Parse request line
	reqLine := strings.Split(lines[0], " ")
	if len(reqLine) != 3 {
		return nil, errors.New("invalid request line")
	}

	req := &Request{
		Method:  reqLine[0],
		Path:    reqLine[1],
		Proto:   reqLine[2],
		Headers: make(map[string]string),
		Body:    body,
	}

	// 4. Parse headers
	for _, line := range lines[1:] {
		if line == "" {
			continue
		}

		kv := strings.SplitN(line, ":", 2)
		if len(kv) != 2 {
			continue
		}

		key := strings.ToLower(strings.TrimSpace(kv[0]))
		val := strings.TrimSpace(kv[1])
		req.Headers[key] = val
	}

	return req, nil
}
