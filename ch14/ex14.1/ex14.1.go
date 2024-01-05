package main

import "fmt"

func main() {
	var a = 10
	var p *int
	p = &a
	fmt.Printf("a의 값: %d\n", a)
	fmt.Printf("p의 값: %p\n", p)
	fmt.Printf("p가 가리키고있는 메모리에 들어있는 값: %d\n", *p)

	fmt.Println("---------")
	*p = 100
	fmt.Printf("p가 가리키고있는 메모리에 들어있는 값: %d\n", *p)
	fmt.Printf("a의 값: %d", a)
}
