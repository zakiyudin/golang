package main

import (
	"fmt"
)

func main() {
	var student string = "John Doe"
	var student2 = "jane Doe"
	x := 100 // this is teh inferred that means the type od the variable based on the value of variable

	fmt.Println("Hello, world!")
	fmt.Println(student)
	fmt.Println(student2)
	fmt.Println(x)

	var a string
	var b int
	var c bool

	a = "string a"
	b = 100
	c = true

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	/*
		! different between "var" and ":="
		* if "var" can be used in inside and outside of a function declaration
		* if ":=" can only be used in inside of a function declaration
	*/

	var A, k, l, m int = 1, 2, 3, 4

	fmt.Println(A, k, l, m)

	// ! Variable Constant

	const PI = 3.14

	fmt.Println(PI)
	fmt.Printf("PI has value : %v and type : %T", PI, PI, "\n")

	var qw, wq = "coba", "saja"
	fmt.Print(qw, "\n")
	fmt.Print(wq, "\n")
}
