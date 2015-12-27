package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	ServiceAddr := ":3000"
	Address, err := net.ResolveTCPAddr("tcp", ServiceAddr)
	checkError(err)
	Listener, err := net.ListenTCP("tcp", Address)
	for {
		conn, err := Listener.Accept()
		if(err != nil) {
			continue //Who cares? We're here to serve more clients.
		}
		conn.Write([]byte(time.Now().String()))
		conn.Close()
	}
}
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}