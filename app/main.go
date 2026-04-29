package main

import (
	"flag"
	"fmt"
	"net"
	"os"
)

const CRLF = "\r\n"

// Ensures gofmt doesn't remove the "net" and "os" imports above (feel free to remove this!)
var _ = net.Listen
var _ = os.Exit

type Config struct {
	Directory string
}

func main() {
	dir := flag.String("directory", ".", "the directory to serve files from")
	flag.Parse()
	cfg := &Config{
		Directory: *dir,
	}
	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}
		go handleConnection(conn, cfg)
	}

}
