package main

import "fmt"

func main() {
	// this creates a slice that is associated with an underlying int array of length 8
	// 10 represents the capacity of the underlying array
	z := make([]int, 8, 20)
	fmt.Println(z)
	fmt.Println(len(z))
	fmt.Println(cap(z))

	z[7] = 10
	fmt.Println(z)

	// if you want to assign a value at z[8] then you have to append the value to z
	// as the length specified is from 0-7
	z = append(z, 20, 30, 40)
	fmt.Println(z)
	fmt.Println(len(z))
	fmt.Println(cap(z))

}

