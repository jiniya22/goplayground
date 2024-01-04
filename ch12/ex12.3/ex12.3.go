package main

import "fmt"

func main() {
	arr := [...]string{"red", "blue", "green"}
	for i, v := range arr {
		fmt.Println(i, v)
	}
}
