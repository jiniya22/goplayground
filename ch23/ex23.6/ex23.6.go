package main

import "fmt"

func f() {
	fmt.Println("f() start")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("panic 복구 - ", r)
		}
	}()
	g()
	fmt.Println("f() finish")
}

func g() {
	fmt.Printf("3 / 2 = %d\n", divide(3, 2))
	fmt.Printf("3 / 0 = %d\n", divide(3, 0))
}

func divide(a, b int) int {
	if b == 0 {
		panic("분모 b는 0일 수 없습니다")
	}
	return a / b
}

func main() {
	f()
	fmt.Println("main 끝부분")
}
