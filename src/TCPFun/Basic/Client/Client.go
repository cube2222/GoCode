package main

import (
	"fmt"
	"net"
	"os"
	"bufio"
)

func main() {
	if len(os.Args) != 2 {
		os.Exit(1)
	}
	conn, err := net.Dial("tcp", "127.0.0.1:3000")
	checkError(err)
	data, err := bufio.NewReader(conn).ReadString('\n')
	checkError(err)
	fmt.Println(data)
}
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}