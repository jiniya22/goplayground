package main

import (
	"fmt"
	"unsafe"
)

func main() {
	s1 := "test"
	fmt.Printf("최초의 s1의 주소: %p\n", unsafe.StringData(s1))

	s1 += "123"
	fmt.Printf("문자열 합산 후 s1의 주소: %p\n", unsafe.StringData(s1))
}
