
package main

import "fmt"

func main() {
	sum(1, 2, 3)
	x := []int{1, 2, 3, 4, 5, 6, 7}
	sum(x...)
}

// using ... before the type name of the last parameter indicates that it takes zero or more of those parameters
func sum(x ...int) {
	total := 0
	for _, v := range x {
		total += v
	}
	fmt.Println("sum=", total)
}

