package main

import "fmt"

func main() {
	s := "AB12가나"
	f1(s)
	fmt.Println()
	f2(s)
	fmt.Println()
	f3(s)
}

func f1(s string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%%d: %d, %%c: %c (%T)\n", s[i], s[i], s[i])
	}
}

func f2(s string) {
	arr := []rune(s)
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%%d: %d, %%c: %c (%T)\n", arr[i], arr[i], arr[i])
	}
}

func f3(s string) {
	for _, v := range s {
		fmt.Printf("%%d: %d, %%c: %c (%T)\n", v, v, v)
	}
}
