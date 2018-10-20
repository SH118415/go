package main

import "fmt"

func main() {

	v := "John"
	switch v {
	case "Jonny":
		fmt.Println("Jonny")
		fallthrough
	case "John":
		fmt.Println("John")
		fallthrough
	case "Betty":
		fmt.Println("Betty")
		fallthrough
	default:
		fmt.Println("default")
	}
}

