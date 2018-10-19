package main

import "os"

func main() {
	// opens an exiting file with Append mode and Read/Write permission
	file,err := os.OpenFile("newfile",os.O_APPEND| os.O_WRONLY,0600)
	if err != nil {
		panic(err)
	}
	file.WriteString("Appending new data\n")
	file.Close()
}

