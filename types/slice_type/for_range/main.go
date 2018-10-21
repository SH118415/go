package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5}
	// i is the index and v is the value
	for i, v := range x {
		fmt.Println(i,v)
	}
}

