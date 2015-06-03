package main

import (
	"fmt"
	"crypto/md5"
)

func main() {
	data := []byte("Hello World")
	fmt.Printf("%x", md5.Sum(data))
}
