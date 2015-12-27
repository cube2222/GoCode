package main
import (
	"net"
	"fmt"
	"bufio"
)

func main() {
	server, err := net.Listen("tcp", ":3000")
	checkError(err)
	conn, err := server.Accept()
	checkError(err)
	for {
		data, err := bufio.NewReader(conn).ReadString('\n')
		checkError(err)
		fmt.Println(string(data))
		conn.Write([]byte("Message Received Printed by Write" + "\n"))
	}
}

func checkError(err error) {
	if (err != nil) {
		fmt.Println(err)
	}
}
