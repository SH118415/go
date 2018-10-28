package main

import "fmt"

func main() {
	// defer schedules a function call to be run after the function completes
	defer sum(1, 2, 3)
	defer sum(10,20,30)
	x := []int{1, 2, 3, 4, 5, 6, 7}
	sum(x...)
}

func sum(x ...int) {
	total := 0
	for _, v := range x {
		total += v
	}
	fmt.Println("sum=", total)
}

// output
// sum= 28
// sum= 60
// sum= 6


