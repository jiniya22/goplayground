package main

import "fmt"

func power(num int) int {
	return num * num
}

func main() {
	n := 5
	result := power(5)
	fmt.Printf("%d의 거듭제곱은 %d\n", n, result)
}
