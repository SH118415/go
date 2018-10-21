package main

import (
	"fmt"
)

type Waiter struct {
	firstName, lastName string
	age, salary int
}

func main() {
	//zero valued structure
	var w1 Waiter 
	fmt.Println(w1)
}
