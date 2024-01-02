package main

import "fmt"

func main() {
	var x int8 = 19
	fmt.Printf("%v = %08b\n", x, x)
	fmt.Printf("%v << 2 = %08b = %d\n", x, x<<2, x<<2)
	fmt.Printf("%v << 3 = %08b = %d\n", x, x<<3, x<<3)
	fmt.Printf("%v << 4 = %08b = %d\n", x, x<<4, x<<4)
	fmt.Printf("%v >> 2 = %08b = %d\n", x, x>>2, x>>2)
}
