package main

import "fmt"

func main() {
	x := 42
	if x > 42 {
		fmt.Println("x is greater than 42")
	} else if x == 42 {
		fmt.Println("x equal to 42")
	} else {
		fmt.Println("x less than 42")
	}
}

