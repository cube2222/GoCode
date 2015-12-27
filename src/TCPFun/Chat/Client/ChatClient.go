package main
import (
	"net"
	"fmt"
	"os"
	"bufio"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:3000")
	checkError(err)
	for {
		data, err := bufio.NewReader(os.Stdin).ReadString('\n')
		checkError(err)
		conn.Write([]byte(data))
		message, err := bufio.NewReader(conn).ReadString('\n')
		fmt.Println(message)
	}
}

func checkError(err error) {
	if (err != nil) {
		fmt.Println(err)
	}
}