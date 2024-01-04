package main

import "fmt"

func main() {
	a := 1
	b := 1

	for ; a <= 9; a++ {
		var found bool
		if b, found = find30(a); found {
			break
		}
	}
	fmt.Printf("%d * %d = %d\n", a, b, a*b)
}

func find30(a int) (int, bool) {
	for b := 1; b <= 9; b++ {
		if a*b == 30 {
			return b, true
		}
	}
	return 0, false
}
