package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	// read from file
	file,_ := os.Open("newfile")

	// creating  a reader for buffer
	reader := bufio.NewReader(file)

	// from the buffer read 10 bytes
	data,_ := reader.Peek(10)
	fmt.Println(string(data))
	file.Close()
}

