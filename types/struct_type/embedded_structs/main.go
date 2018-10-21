
package main

import "fmt"

// A struct is a type which contains named fields
type Passenger struct {
	name         string
	age          int
	gender       string
	ticketNumber int
}

type secretAgent struct {
	Passenger
	wears string
}

func main() {

	p1 := secretAgent{
		Passenger: Passenger{
			name:         "Tim",
			age:          28,
			gender:       "Male",
			ticketNumber: 12345},
		wears: "suit"}
	fmt.Println(p1)
	fmt.Println(p1.Passenger.ticketNumber)

	p2 := Passenger{name: "Alan", age: 34, gender: "Male", ticketNumber: 12344}
	fmt.Println(p2)
	fmt.Println(p2.age)

}

