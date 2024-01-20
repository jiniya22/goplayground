package main

import "fmt"

func sum(nums ...int) int {
	result := 0
	for _, num := range nums {
		result += num
	}
	return result
}

func main() {
	fmt.Println(sum(2))
	fmt.Println(sum(5, 6, 10))
	fmt.Println(sum(22, 99))
}
