package main

import "fmt"

func PrintValue(v interface{}) {
	switch t := v.(type) {
	case int:
		fmt.Printf("%d는 int 타입", t)
	case float64:
		fmt.Printf("%f는 float64 타입", t)
	case string:
		fmt.Printf("%s는 string 타입", t)
	default:
		fmt.Printf("%v는 int, float64, string 외의 타입", t)
	}
	fmt.Println()
}

func main() {
	PrintValue(32)
	PrintValue(3.14)
	PrintValue("지니")
	var a int8 = 9
	PrintValue(a)
}
