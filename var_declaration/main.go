package main

import "fmt"

// declaring a variable 'x'
var x =20

// declaring a variable z of type int 
// assigns a zero value
var z int

func main() {

	// shorthand declaration of a variable 'b' and is assigned to a value 12
	b := 12
	fmt.Println(b)

	// once the b is declared we can assign another value to 'b' just using =
	b = 6
	fmt.Println(b)

	fmt.Println(x)
	fmt.Println(z)
}
	// note short declaration operator can be used only inside the function
