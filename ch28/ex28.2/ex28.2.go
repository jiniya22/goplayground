package main

import "fmt"

func fibonacci1(n int) int {
	if n < 0 {
		return 0
	} else if n < 2 {
		return n
	}
	return fibonacci1(n-2) + fibonacci1(n-1)
}

func fibonacci2(n int) int {
	if n < 0 {
		return 0
	} else if n < 2 {
		return n
	}
	first := 0
	second := 1
	result := 0

	for i := 2; i <= n; i++ {
		result = first + second
		first = second
		second = result
	}
	return second
}

func main() {
	fmt.Println(fibonacci1(10))
	fmt.Println(fibonacci2(10))
}
