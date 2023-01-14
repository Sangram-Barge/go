package main

import "fmt"

func main() {
	/*
		var -> variable
		a -> name of variable
		int/float64/string -> data type
		1/1.0/"hello" -> assignment
	*/
	var a int = 1
	var b float64 = 1.0
	var c string = "hello"

	/*
		type inference based on value
		use walrus operator 
		(similar to var a = "hello" from java perspective type of a inferred)
	*/

	d := 3.0
	
	e := b + d

	/*
		functions exported from package must start with capital case
		%v -> any var
		%T -> type of variable
	*/

	fmt.Printf("value of a %v type %T\n", a, a)
	fmt.Printf("value of b %v type %T\n", b, b)
	fmt.Printf("value of c %v type %T\n", c, c)
	fmt.Printf("value od d %v type %T\n", d, d)
	fmt.Printf("value od e %v type %T\n", e, e)
}