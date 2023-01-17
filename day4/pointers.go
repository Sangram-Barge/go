package main

import "fmt"

var pr = fmt.Printf

func main() {
	arr := []int{1, 2}
	doubleArr(arr)
	pr("arry value : %v\n", arr)

	val := 2
	doubleVal(val)
	pr("val (pass by value) : %v \n", val)

	vaptr := &val
	doublePtr(vaptr)
	pr("val (pass by referance) : %v \n pointer value %v\n", val, vaptr)
}

func doubleArr(arr []int) {
	for i := range arr {
		arr[i] *= 2
	}
}

func doubleVal(val int) {
	val *= 2
}

func doublePtr(val *int) {
	*val *= 2
} 
