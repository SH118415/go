package main

import "fmt"

func main() {
	// A map is an unordered collection of key-value pairs
	m := map[string]int{"Bella": 25, "Edward": 32, "Jack": 33}
	fmt.Println(m["Bella"])

	// checks whether a value exists in a map
	// it assigns a value zero if it does not exist
	// if ok is false then the value does not exist
	v, ok := m["Alice"]
	fmt.Println(v)
	fmt.Println(ok)

	if v, ok := m["Edward"]; ok {
		fmt.Println(v)
	}
}

