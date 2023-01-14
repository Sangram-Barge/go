package main

import "fmt"

func main() {
	printEvenEnded()
}

func printEvenEnded() {
	print := fmt.Printf
	count := 0
	for a := 1000; a < 9999; a++ {
		for b := a; b < 9999; b++ {
			numStr := fmt.Sprintf("%v", a * b)
			if numStr[0] == numStr[len(numStr) - 1] {
				count++
			}
		}
	}
	print("%v", count)
}