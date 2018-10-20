package main

import (
	"fmt"
)

func main() {
	for i := 0; i <= 5; i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("The outer loop is %d\t The inner loop is %d\n", i, j)
		}
	}
}

