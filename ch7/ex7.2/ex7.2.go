package main

import "fmt"

func main() {
	fmt.Println(F(9))
}

func F(n int) int {
	if n < 2 {
		return n
	}
	return F(n-2) + F(n-1)
}
