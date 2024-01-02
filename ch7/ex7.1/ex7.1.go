package main

import "fmt"

func Divide(a, b int) (int, bool) {
	if b == 0 {
		return 0, false
	}
	return a / b, true
}

func main() {
	result, success := Divide(5, 3)
	fmt.Println(result, success) // 1 true
	result2, success := Divide(5, 0)
	fmt.Println(result2, success) // 0 false
}

func Multiple(a, b int) int {
	return a * b
}
