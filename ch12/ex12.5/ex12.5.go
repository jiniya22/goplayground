package main

import "fmt"

func main() {
	a := [3]int{10, 20, 30}
	b := [3]int{1, 2, 3}
	fmt.Println(a, &a[0])
	fmt.Println(b, &b[0])

	fmt.Println("-----------")

	a = b
	fmt.Println(a, &a[0])
	fmt.Println(b, &b[0])
}
