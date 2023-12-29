package main

import "fmt"

func main() {
	input2()
}

func input1() {
	var a string
	var b string

	n, err := fmt.Scanf("%3s%2s", &a, &b)
	if err != nil {
		fmt.Println(n, err)
	} else {
		fmt.Println(n, a, b)
	}
}

func input2() {
	var a int
	var b int

	n, err := fmt.Scanln(&a, &b)
	if err != nil {
		fmt.Println(n, err)
	} else {
		fmt.Println(n, a, b)
	}
}
