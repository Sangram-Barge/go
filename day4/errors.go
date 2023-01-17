package main

import (
	"fmt"
	"math"
)

var pr = fmt.Printf

func main() {
	n := -2.0
	root, err := getroot(n)	
	if err != nil {
		pr("ERROR: %v\n", err)
	} else {
		pr("sqrt of %v is %v\n", n, root)
	}
}

func getroot(val float64) (float64, error) {
	if val < 0.0 {
		return 0.0, fmt.Errorf("%v is less than zero \n", val)
	}
	return math.Sqrt(val), nil
}
