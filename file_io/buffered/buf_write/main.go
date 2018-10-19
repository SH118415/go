package main

import (
	"os"
	"bufio"
)

func main() {
	file,err:=os.Create("newfile")
	if err!=nil {
		panic(err)
	}
	// create a buffer for the os file
	buf := bufio.NewWriter(file) 

	// write data to buffer
	buf.WriteString("Writing data to buffer\n")
	
	// clears the buffer data and writes to the file
	buf.Flush()
	file.Close()
}

