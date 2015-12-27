package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: " + os.Args[0] + " host:port")
		return
	}
	addr, err := net.ResolveTCPAddr("tcp", os.Args[1])
	checkError(err)
	conn, err := net.DialTCP("tcp", nil, addr)
	checkError(err)
	data, err := ioutil.ReadAll(conn)
	checkError(err)
	fmt.Println(string(data))
	conn.Close()
}
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}