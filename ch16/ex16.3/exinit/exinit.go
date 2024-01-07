package exinit

import "fmt"

var (
	a = c + b
	b = f()
	c = f()
	d = 3
)

func init() {
	d++
	fmt.Println("init 함수 실행! d: ", d)
}

func f() int {
	d++
	fmt.Println("f() 함수 실행! d: ", d)
	return d
}

func PrintD() {
	fmt.Println("d: ", d)
}
