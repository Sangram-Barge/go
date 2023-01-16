package main

import "fmt"

var pr = fmt.Printf

func main( ) {
	
	countries := map[string]int{
		"india": 1,
		"russia": 2,
		"japan": 3,
	}

	pr("%v \n", len(countries))

	pr("rank of india %v \n", countries["india"])

	pr("rank of china %v \n", countries["china"]) // as china is not in map this will print 0

	value, isPresent := countries["china"]

	if isPresent {
		pr("%v is present \n", value)
	} else {
		pr("value is not present instead gave %v \n", value)
	}

	countries["china"] = 4
	val1, prs := countries["china"]
	if prs {
		pr("%v is present \n", val1)
	} else {
		pr("value is not present instead gave %v \n", val1)
	}

	delete(countries, "china")

	for i := range countries {
		pr("%v value, %T type \n", i, i)
	}

}