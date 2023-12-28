package main

import "fmt"

func main() {
	var n1 = 43
	fmt.Printf("%%v(기본): %v\n", n1)
	fmt.Printf("%%T(데이터타입): %T\n", n1)
	fmt.Printf("%%d(10진수 정수): %d\n", n1)
	fmt.Printf("%%b(2진수 정수): %b\n", n1)
	fmt.Printf("%%o(8진수 정수): %o\n", n1)
	fmt.Printf("%%O(8진수 정수, prefix로 0o 붙임): %O\n", n1)
	fmt.Printf("%%x(16진수 정수): %x\n", n1)
	fmt.Printf("%%X(16진수 정수): %X\n", n1)
	fmt.Printf("%%c(유니코드 문자): %b\n", n1)

	fmt.Print()

	fmt.Printf("%v = %c(유니코드), %v = %c(유니코드)\n", 65, 65, 97, 97)
	var n2 = 1.2
	fmt.Printf("%d, %o, %c, %s, %t\n", n2, n2, n2, n2, n2)
	fmt.Printf("%x, %b\n", n2, n2)

}
