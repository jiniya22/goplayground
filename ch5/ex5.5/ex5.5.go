package main

import "fmt"

func main() {

	var v1 = true
	fmt.Printf("%%T(데이터타입): %T\n", v1)
	fmt.Printf("%%t(boolean): %t\n", v1)

	fmt.Println("------------")

	var v2 = "hello\njiniworld\t!"
	fmt.Printf("%%s(문자열): %s\n", v2)
	fmt.Printf("%%q(문자열, 특수문자 기능 동작 무시): %q\n", v2)

	fmt.Println("------------")

	fmt.Printf("%%p(메모리 주소값) %p", v2)
}
