// conversions means take value of one type and convert it to other type

package main

import "fmt"

// a is of type int
var a int 

// type hello is a  int
type hello int 
var b hello

func main() {
        fmt.Printf("%T\n",a)
        fmt.Printf("%T\n",b)
	//converting b of type hello to type int
	a = int(b)
	fmt.Printf("%T\n",a)
}




