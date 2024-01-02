package main

import "fmt"

const PI = 3.14

func main() {
	var b float32 = PI * 2
	var b2 = PI * 1.2
	var b3 = PI + 8

	var c float32 = 8.12
	var b4 = PI * c

	var d = 9.87
	var b5 = PI * d

	//var a = 8
	//var b5 = PI * a		// Error

	fmt.Printf("%v, %T\n", b, b)
	fmt.Printf("%v, %T\n", b2, b2)
	fmt.Printf("%v, %T\n", b3, b3)
	fmt.Printf("%v, %T\n", b4, b4)
	fmt.Printf("%v, %T\n", b5, b5)
}
