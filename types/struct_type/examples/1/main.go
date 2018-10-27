// create your own type “person” which will have an underlying type of “struct” so that it can store the following data:
// name
// age
// favorite colours
// print out the values, ranging over the elements in the slice which stores the favorite colours.

package main

import (
	"fmt"
)

type person struct {
	name   string
	age    int
	colour []string
}

func main() {
	p1 := person{name: "John", age: 12,
		colour: []string{"red", "green", "orange"},
	}
	fmt.Println(p1.name)
	fmt.Println(p1.age)
	for i, v := range p1.colour {
		fmt.Println(i, v)
	}
}

