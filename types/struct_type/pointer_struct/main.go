package main

import (
	"fmt"
)

type Employee struct {
	firstName, lastName string
	age, salary         int
}

func main() {

	// the language gives us the option to use emp8.firstName
	// instead of the explicit dereference (*emp8).firstName to access the firstName field
	emp8 := &Employee{"Sam", "Anderson", 55, 6000}
	fmt.Println("First Name:", emp8.firstName)
	fmt.Println("Age:", emp8.age)
}

