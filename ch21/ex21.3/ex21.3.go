package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func multiply(a, b int) int {
	return a * b
}

type opFunc func(int, int) int

func getOperator(op string) opFunc {
	if op == "+" {
		return add
	} else if op == "-" {
		return func(a int, b int) int {
			return a - b
		}
	} else if op == "*" {
		return multiply
	}
	return nil
}

func main() {
	operator := getOperator("+")
	fmt.Println(operator(3, 10))
	fmt.Println(getOperator("*")(5, 2))
}
