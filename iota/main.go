package main

import (
	"fmt"
)

func main() {
	// iota identifier is used in const declarations to simplify definitions of incrementing numbers
	// The values of iota is reset to 0 whenever the reserved word const appears in the source (i.e. each const block)
	// and increments by one after each ConstSpec e.g. each Line.
	const (
		a = iota
		b
		c
	)
	const (
		d = iota
		e = iota
		f = iota
	)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf(" %T\n %T\n %T\n", a, b, c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Printf(" %T\n %T\n %T\n", d, e, f)

}

