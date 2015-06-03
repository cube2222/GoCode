package main

import (
	"net"
	"fmt"
	"time"
)

func main() {
	ln, err := net.Listen("tcp", ":3000")
	if err != nil {
		fmt.Println(err)
	}
	for {
		conn, err := ln.Accept()
		if (err != nil) {
			fmt.Println(err)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	var data byte
	conn.SetDeadline(time.Duration.Seconds(10))
	for a := 0; a < 10; a++ {
		conn.Read()
	}
}
