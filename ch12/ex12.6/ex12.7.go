package main

import "fmt"

func main() {
	a := [2][5]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}}

	fmt.Println(a)
	fmt.Println(a[0])
	fmt.Println(a[1])

	b := [...]int{1, 2, 3, 4, 5}
	fmt.Println(b)
}
