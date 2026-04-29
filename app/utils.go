package main

import (
	"slices"
	"strings"
)

func parseEncodings(header string) []string {
	var parsedEncodings []string
	encodings := strings.SplitSeq(header, ",")
	for encoding := range encodings {
		encoding = strings.TrimSpace(encoding)
		parsedEncodings = append(parsedEncodings, encoding)
	}
	return parsedEncodings
}

func getValidEncoding(header string) string {
	data := parseEncodings(header)
	for _, encoding := range data {
		if slices.Contains(acceptedEncodings, encoding) {
			return encoding
		}
	}
	return ""
}
