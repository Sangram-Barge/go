package main

import "fmt"

var pr = fmt.Printf

func main() {
	a := add(1, 3)
	pr("%v\n", a)

	b, c := addSub(1, 3)
	pr("%v %v \n", b, c)

	e, f, g := addSubMul(2, 3)
	pr("%v %v %v", e, f, g)
}

func add(a int, b int) int {
	return a + b
}

func addSub(a int, b int) (int, int) {
	return a+b, a-b
}

func addSubMul(a int, b int) (int, int, int) {
	return a+b, a-b, a*b
}