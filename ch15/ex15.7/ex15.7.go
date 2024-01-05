package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	s1 := "apple"
	s2 := s1
	fmt.Printf("s1의 주소: %p, %d\n", unsafe.StringData(s1), unsafe.StringData(s1))
	fmt.Printf("s2의 주소: %p, %d\n", unsafe.StringData(s2), unsafe.StringData(s2))

	fmt.Println((*reflect.StringHeader)(unsafe.Pointer(&s1)))
	fmt.Println((*reflect.StringHeader)(unsafe.Pointer(&s2)))
}
