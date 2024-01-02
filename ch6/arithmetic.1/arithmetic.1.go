package main

import "fmt"

func main() {
	var x = 13
	var y = 4

	var f1 = 7.14
	var f2 float64 = 5

	fmt.Println("x + y = ", x+y)
	fmt.Println("x - y = ", x-y)
	fmt.Println("x * y = ", x*y)
	fmt.Println("x / y = ", x/y)
	fmt.Println("x % y = ", x%y)

	fmt.Printf("%v + %v = %v\n", f1, f2, f1+f2)
	fmt.Printf("%v - %v = %v\n", f1, f2, f1-f2)
	fmt.Printf("%v * %v = %v\n", f1, f2, f1*f2)
	fmt.Printf("%v / %v = %v\n", f1, f2, f1/f2)
}
