package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	server, err := net.Listen("tcp", ":3000")
	checkError(err)
	for {
		conn, err := server.Accept()
		if err != nil {
			continue
		}
		conn.Write(append([]byte(time.Now().String()), '\n'))
		conn.Close()
	}
}
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}