package main

import (
	"os"
	"log"
)

func main() {
	// create a new file myfile
	file,err := os.Create("myfile")
	if err!=nil {
		log.Fatal(err)
	}
	 data:=[]byte("Hello world\n")

	// data as a byte array	
	file.Write(data)

	// write data as a string 
	file.WriteString("This is a String\n") 
	file.Close()
} 
