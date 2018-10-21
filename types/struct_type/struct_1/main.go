package main

import "fmt"

// A struct is a type which contains named fields
type passenger struct {
	Name         string
	Age          int
	Gender       string
	TicketNumber int
}

func main() {

	person1 := passenger{Name: "Tim", Age: 28, Gender: "Male", TicketNumber: 12345}
	fmt.Println(person1)
	fmt.Println(person1.TicketNumber)

	person2 := passenger{Name: "Alan", Age: 34, Gender: "Male", TicketNumber: 12344}
	fmt.Println(person2)
	fmt.Println(person2.Age)

}

