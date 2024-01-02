package main

import "fmt"

const (
	APPLE = iota + 11
	BANANA
	GREEN
)

const (
	C1 = 10 << iota
	C2
	C3
	C4
)

func main() {
	fmt.Printf("APPLE: %v, BANANA: %v, GREEN: %v\n", APPLE, BANANA, GREEN)
	fmt.Printf("C1: %v, C2: %v, C3: %v, C4: %v\n", C1, C2, C3, C4)
}
