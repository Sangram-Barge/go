package main 

import "fmt"

func main() {
	fizzbuzz()
}

func fizzbuzz() {
	for num := 1; num <= 20; num++ {
		if num % 3 == 0 && num % 5 == 0 {
			fmt.Println("fizzbuzz")
		} else if num % 3 == 0 {
			fmt.Println("fizz")
		} else if num % 5 == 0 {
			fmt.Println(("buzz"))
		} else {
			fmt.Println(num)
		}
	}
}