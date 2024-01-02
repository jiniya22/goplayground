package main

import "fmt"

func main() {
	var x int8 = 43
	var y int8 = 1

	fmt.Printf("%v &^ %v = %08b &^ %08b = %08b = %v\n", x, y, x, y, x&^y, x&^y)

	y = 2
	fmt.Printf("%v &^ %v = %08b &^ %08b = %08b = %v\n", x, y, x, y, x&^y, x&^y)

	y = 33
	fmt.Printf("%v &^ %v = %08b &^ %08b = %08b = %v\n", x, y, x, y, x&^y, x&^y)
}
