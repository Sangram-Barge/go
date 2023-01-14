package main

import (
	"fmt"
)

func main() {
	a := 1
	b := 2
	c := 4

	/*
		condition can be written wo brackets 
	*/

	if a < b {
		fmt.Printf("%v is less than %v \n", a, b)
	} else {
		fmt.Printf("%v is less than %v \n", b, a)
	}

	/*
		assignment can be done in condition itself saperated by semicolon (;) (d will be visible inside if block only)
	*/

	if d := a * b; d < 10 {
		fmt.Printf("%v is less than %v \n", d, 10)
	}
	/*
		just normal switch case (it seems break is not needed)
	*/

	switch c {
	case 3: fmt.Println("case is correct!!")
	case 2: fmt.Println("no no no no!!!!!!(insert michel scott)")
	default: fmt.Println("how the hell did I get here 🤨")
	}
}