package main

import "fmt"

// a is of type int
var a int

// Type hello is a int
type hello int 
var b hello

func main() {
	fmt.Printf("type of a: %T\n",a)
	fmt.Printf("type of b: %T\n",b)
}


