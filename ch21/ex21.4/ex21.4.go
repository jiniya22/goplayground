package main

import "fmt"

func main() {
	num := 1
	f := func(a int) {
		num *= a
	}
	num += 100
	f(2)
	fmt.Println(num)
}
