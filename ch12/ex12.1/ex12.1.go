package main

import "fmt"

func main() {
	arr := [5]float32{3.14, 5.9, 9.0, 7, 12.3}

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
	fmt.Println(arr)

	var a float32 = 7.0
	fmt.Println(a)
}
