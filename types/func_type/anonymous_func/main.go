
package main

import (
	"fmt"
)

func main() {

	func() {
		fmt.Println("Anonymous func")
	}()

	func(x int) {
		fmt.Println("The value of x=", x)
	}(8)

	msg := []string{"Red", "Black", "Green"}

	func(m []string) {
		fmt.Println(m)
	}(msg)
	
	f := func (s string) {
	fmt.Printf("%T\n",s)
	} 
	f("hello")

}

