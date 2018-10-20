package main

import "fmt"

func main() {
	i := 50
	switch i {
	case 10:
		fmt.Println(10)
	case 20:
		fmt.Println(20)
	default:
		fmt.Println(i)
	}
}
