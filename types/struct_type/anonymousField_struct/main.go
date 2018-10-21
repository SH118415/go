package main

import (
	"fmt"
)

type Passenger struct {
	string
	int
}

func main() {
	p := Passenger{"Naveen", 50}
	fmt.Println(p)
}

