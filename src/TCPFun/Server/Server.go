package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	ServiceAddr := ":3000"
	RealAddress, err := net.ResolveTCPAddr("tcp", ServiceAddr)
	checkError(err)
	serv, err := net.ListenTCP("tcp", RealAddress)
	checkError(err)
	for {
		conn, err := serv.Accept()
		if err != nil {
			continue
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