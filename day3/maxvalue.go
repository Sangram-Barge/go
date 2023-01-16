package main

import "fmt"

var pr = fmt.Printf

func main() {
	nums := []int{1, 2, 43, 23 , 57, 23, 54}

	var max = 0

	for _, num := range nums {
		if (num > max) {
			max = num
		}
	}

	pr("max number is %v \n", max)
}