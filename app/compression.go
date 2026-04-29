package main

import (
	"bytes"
	"compress/gzip"
)

func gzipCompress(data []byte) ([]byte, error) {
	var buf bytes.Buffer

	gz := gzip.NewWriter(&buf)

	_, err := gz.Write(data)
	if err != nil {
		return nil, err
	}

	// IMPORTANT: close to flush all data
	if err := gz.Close(); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
