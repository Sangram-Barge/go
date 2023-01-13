package main

import "fmt"

func main() {
	var a int = 1
	var b float64 = 1.0
	var c string = "hello"
	d := 3.0
	
	e := b + d

	fmt.Printf("value of a %v type %T\n", a, a)
	fmt.Printf("value of b %v type %T\n", b, b)
	fmt.Printf("value of c %v type %T\n", c, c)
	fmt.Printf("value od d %v type %T\n", d, d)
	fmt.Printf("value od e %v type %T\n", e, e)
}