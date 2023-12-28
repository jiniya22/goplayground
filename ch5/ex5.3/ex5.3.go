package main

import "fmt"

func main() {
	var n1 = 1234.56789

	fmt.Printf("%%v(기본): %v\n", n1)
	fmt.Printf("%%T(데이터타입): %T\n", n1)
	fmt.Printf("%%f(실수): %f\n", n1)
	fmt.Printf("%%e(지수표기법): %e\n", n1)
	fmt.Printf("%%E(지수표기법): %E\n", n1)
	fmt.Printf("%%f(실수 소수점 3자리까지): %.3f\n", n1)
}
