package main

import (
	"fmt"
)

func main() {
	// s is a variable of string type
	s := "Hello World"
	fmt.Printf("%T\n", s)

	// converting a string into a slice of byte
	// byte is a alias of uint8
	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)

}

