package main

import "fmt"

func main() {
	m := map[string]int{"Bella": 25, "Edward": 32, "Jack": 33}
	fmt.Println(m)

	// delete a Jack from map
	delete(m, "Jack")
	fmt.Println(m)

}

