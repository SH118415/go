package main

import (
	"fmt"
)

func main() {

	x := 8
	fmt.Printf("%d\t %b\n", x, x)

	// left shift by 1 bit
	y := x << 1
	fmt.Printf("%d\t %b\t", y, y)

}

