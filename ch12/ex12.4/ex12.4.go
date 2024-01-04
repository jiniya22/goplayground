package main

import "fmt"

func main() {
	arr := [...]int{11, 12, 13}
	for i := 0; i < len(arr); i++ {
		fmt.Printf("인덱스: %d, 값: %d, 주소값: %d\n", i, arr[i], &arr[i])
	}
}
