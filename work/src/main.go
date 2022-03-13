package main

import (
	"fmt"
	"bytes"
)

func main() {
	var buffer bytes.Buffer
	buffer.Write([]byte("Hello world!\n"))
	fmt.Println(buffer.String())
}
