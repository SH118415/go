package main

import "fmt"

func main() {
	a := "Hello world"
	for m := range a {
		fmt.Println(m)
	}
}
