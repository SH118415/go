
package main

import "fmt"

// A struct is a type which contains named fields
type Passenger struct {
	name         string
	age          int
	gender       string
	ticketNumber int
}

func main() {

	person1 := Passenger{name: "Tim", age: 28, gender: "Male", ticketNumber: 12345}
	fmt.Println(person1)
	fmt.Println(person1.ticketNumber)

	person2 := Passenger{name: "Alan", age: 34, gender: "Male", ticketNumber: 12344}
	fmt.Println(person2)
	fmt.Println(person2.age)

}

