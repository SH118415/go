package main

import "fmt"

func main() {
	// A map is an unordered collection of key-value pairs
	m := map[string]int{"Bella": 25, "Edward": 32, "Jack": 33}
	fmt.Println(m["Bella"])

	// checks whether a value exists in a map
	// it assigns a value zero if it does not exist
	// ok is false then the value does not exist
	v, ok := m["Alice"]
	fmt.Println(v)
	fmt.Println(ok)
	
	// adding Alice to m
	m["Alice"] = 27
	for k,v :=range m {
		fmt.Println(k,v)
	}
}

