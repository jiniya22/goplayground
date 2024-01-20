package main

import (
	"container/ring"
	"fmt"
)

func main() {
	r := ring.New(5)
	n := r.Len()

	for i := 0; i < n; i++ {
		r.Value = i
		r = r.Next()
	}

	for i := 0; i < n; i++ {
		fmt.Print(r.Value, ", ")
		r = r.Next()
	}
	fmt.Println("\n--------------")
	for i := 0; i < n; i++ {
		fmt.Print(r.Value, ", ")
		r = r.Prev()
	}
}
