package main

import "fmt"

func main() {
	var x int8 = 18
	var y int8 = 20

	fmt.Printf("%v & %v = %08b & %08b = %v\n", x, y, x, y, x&y)
	fmt.Printf("%v | %v = %08b | %08b = %v\n", x, y, x, y, x|y)
	fmt.Printf("%v ^ %v = %08b ^ %08b = %v\n", x, y, x, y, x^y)
}
