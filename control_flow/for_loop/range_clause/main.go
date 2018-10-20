package main

import "fmt"

func main() {
	a := "Hello world"
	for m := range a {
		fmt.Println(m)
	}

	//on a map, range returns the key
	capitals := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo"}
	for key := range capitals {
		fmt.Println("Map item: Capital of", key, "is", capitals[key])
	}

	//range can also return two items, the index/key and the corresponding value
	for key, val := range capitals {
		fmt.Println(" Item: Capital of", key, "is", val)
	}
}

