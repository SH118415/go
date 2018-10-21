
package main

import "fmt"

// A struct is a type which contains named fields
type passenger struct {
	Name         string
	Age          int
	Gender       string
	TicketNumber int
}

type secretAgent struct {
	passenger
	wears string
}

func main() {

	p1 := secretAgent{
		passenger: passenger{
			Name:         "Tim",
			Age:          28,
			Gender:       "Male",
			TicketNumber: 12345},
		wears: "suit"}
	fmt.Println(p1)
	fmt.Println(p1.passenger.TicketNumber)

	p2 := passenger{Name: "Alan", Age: 34, Gender: "Male", TicketNumber: 12344}
	fmt.Println(p2)
	fmt.Println(p2.Age)

}

