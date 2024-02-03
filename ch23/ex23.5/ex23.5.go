package main

import "fmt"

func divide(a, b int) {
	if b == 0 {
		panic("분모 b는 0일 수 없습니다")
	}
	fmt.Printf("%d / %d = %d\n", a, b, a/b)
}

func main() {
	divide(3, 2)
	divide(3, 0)
}
