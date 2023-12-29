package main

import "fmt"

func main() {
	var a, b int

	n, err := fmt.Scan(&a, &b)
	if err != nil {
		fmt.Println(n, err)
	} else {
		fmt.Println(n, a, b)
	}

	var c, d string
	n, err = fmt.Scan(&c, &d)
	if err != nil {
		fmt.Println(n, err)
	} else {
		fmt.Println(n, c, d)
	}
}
