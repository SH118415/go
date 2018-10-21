
package main

import "fmt"

// declaring structures without declaring a new type are called anonymous structures

func main() {

	p1 := struct {
		Name         string
		Age          int
		Gender       string
		TicketNumber int
	}{
		Name:         "Tim",
		Age:          28,
		Gender:       "Male",
		TicketNumber: 12345,
	}
	fmt.Println(p1)
	fmt.Println(p1.TicketNumber)

}

