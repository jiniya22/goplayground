package main

import "fmt"

func main() {
	var ch rune
	ch = '가'
	fmt.Printf("%%c : %c, %%d: %d, 데이터 타입: %T", ch, ch, ch)
}
