package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

func main() {
	if(len(os.Args) != 2) {
		fmt.Fprintln(os.Stderr, "Usage: %s host:port", os.Args[0])
		os.Exit(1)
	}
	target := os.Args[1]
	address, err := net.ResolveTCPAddr("tcp", target)
	checkError(err)
	conn, err := net.DialTCP("tcp", nil, address)
	checkError(err)
	//_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n")) //Uncomment to send http request.
	checkError(err)
	data, err := ioutil.ReadAll(conn)
	checkError(err)
	fmt.Println(string(data))
}
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}