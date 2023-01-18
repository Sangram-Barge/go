package main

import "fmt"

var pr = fmt.Printf

func main() {
	worker()	
}

func worker() {
	val, err := aquire("A")
	if (err != nil) {
		pr("error\n")
	} 
	defer relese(val)

	val2, err2 := aquire("B")
	if (err2 != nil) {
		pr("error \n")
	}
	defer relese(val2)

	pr("worker\n")
}

func aquire(val string) (string, error) {
	return val, nil
}

func relese(val string)	{
	pr("relesing %s \n", val)
}
