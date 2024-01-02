package main

import "fmt"

const LIMIT = 10

func main() {
	var a int8 = LIMIT * 4
	var a2 int32 = LIMIT * 8
	var a3 float32 = LIMIT * 8
	var a4 complex64 = LIMIT * 8
	var a5 = LIMIT * 2.38

	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", a2, a2)
	fmt.Printf("%v, %T\n", a3, a3)
	fmt.Printf("%v, %T\n", a4, a4)
	fmt.Printf("%v, %T\n", a5, a5)
}
