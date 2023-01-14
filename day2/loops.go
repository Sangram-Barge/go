package main 

import "fmt"

func main() {
	/* normal for loop */
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	fmt.Println("---")
	/* kind of like a while loop */
	a := 0
	for a < 3 {
		fmt.Println(a)
		a += 1
	}
	fmt.Println("---")
	/* like a while true loop */
	a = 0
	for {
		fmt.Println(a)
		if (a > 3) {
			break
		}
		a += 1
	}
}