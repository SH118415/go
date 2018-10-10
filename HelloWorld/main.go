// Basic go program for printing HelloWorld

// Package declaration
package main  

//The import keyword is how we include code from other packages to use with our program. The fmt package implements formatting for input and output.
import "fmt"

//main is the name of the function 
func main(){
	
//First we access another function inside of the fmt package called Println (that's the fmt.Println piece, Println means Print Line). Then we create a new string that contains Hello World.
	fmt.Println("Hello World") 
}
