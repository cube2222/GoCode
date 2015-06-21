package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	data := []byte("Hello World")
	fmt.Printf("%x", md5.Sum(data))
}
