package main

import "fmt"

func main() {
	var n1 = 1234.56789
	fmt.Println(n1)
	fmt.Println(n1 * 1.2)

	fmt.Println()

	// float32 타입은 소수부 7자리까지 표현가능
	var n2 float32 = 1234.56789
	fmt.Println(n2)       // 맨마지막 수는 반올림처리 됨 1234.5679
	fmt.Println(n2 * 1.2) // 1234.5678 * 1.2 = 1481.4814
	fmt.Printf("%f", n2*1.2)
}
