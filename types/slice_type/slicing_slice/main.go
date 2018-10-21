package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8}
	// x[1:3]low is the index of where to start the slice and high is the index where to end it (but not including the index itself)
	fmt.Println(x[1:3])
	fmt.Println(x[:])
	fmt.Println(x[1:])
	fmt.Println(x[:5])
}

