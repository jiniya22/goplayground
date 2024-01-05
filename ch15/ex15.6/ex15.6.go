package main

import (
	"fmt"
)

func main() {
	s1 := "hi"
	s2 := "wow"
	s3 := s1 + " " + s2

	s1 += " " + s2

	fmt.Println(s3)
	fmt.Println(s1)
	fmt.Println(&s1)
}
