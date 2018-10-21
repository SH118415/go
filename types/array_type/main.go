package main

import "fmt"

func main() {

	// An array is a numbered sequence of elements of a single type with a fixed length
	var x [5]int
	x[4] = 100
	fmt.Println(x)
	fmt.Println(len(x))
}

