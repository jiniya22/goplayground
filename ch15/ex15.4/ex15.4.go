package main

import "fmt"

func main() {
	s1 := "abC4"
	s2 := "가나다"

	r1 := []rune(s1)
	r2 := []rune(s2)

	fmt.Println(len(r1))
	fmt.Println(len(r2))
}
