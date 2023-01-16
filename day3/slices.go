package main

import "fmt" 
var pr = fmt.Printf

func main() {
	names := []string{"sangram", "gouri", "anushree"}
	pr("%v type %T\n", names, names)	

	pr("%v \n", names[0])

	pr("%v\n", len(names))

	pr("%v \n", len(names[1]))

	for i := range names {
		pr("%v ", i)
	}
	pr("\n")

	for i, name := range names {
		pr("%v is at %v \n", name, i )
	}

	for _, name := range names {
		pr("%v \n", name)
	}

	names = append(names, "mangesh")
	pr("%v \n", names)
}