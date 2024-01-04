package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Print(i, ", ")
	}
	fmt.Println("\n----")
	i := 0
	for ; i < 10; i++ {
		fmt.Print(i, ", ")
	}
	fmt.Println("\n----")

	for i := 0; i < 10; {
		fmt.Print(i, ", ")
		i++
	}
	fmt.Println("\n----")

	i = 0
	for i < 10 {
		fmt.Print(i, ", ")
		i += 2
	}
}
