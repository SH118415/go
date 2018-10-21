package main

import "fmt"

func main() {

	z := []int{22, 23, 24, 25, 6, 10}
	fmt.Println(z)

	// appending new values to z
	z = append(z, 1, 2, 3, 4, 5)
	fmt.Println(z)

	d := []int{100, 200, 300, 3, 6, 9}
	z = append(z, d...)
	fmt.Println(z)

	// deleting from a slice
	z = append(z[:3], z[7:]...)
	fmt.Println(z)

}

