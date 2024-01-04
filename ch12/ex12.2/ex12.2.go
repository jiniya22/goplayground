package main

import "fmt"

func main() {
	ex1()
	ex2()
	ex3()
}

func ex1() {
	arr := [5]int{5, 2}
	fmt.Println(arr)

	arr2 := [5]string{"apple", "banana"}
	fmt.Println(arr2)
}

func ex2() {
	arr := [5]int{2: 11, 4: 98}
	fmt.Println(arr)
}
func ex3() {
	arr := [...]string{"red", "blue", "green"}
	fmt.Println(arr, len(arr))
}
