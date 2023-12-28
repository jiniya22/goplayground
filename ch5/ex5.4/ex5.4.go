package main

import "fmt"

func main() {
	var n1 = 900
	for i := n1; i > 0; i-- {
		fmt.Printf("%v(%c), ", i, i)
		if i%10 == 0 {
			fmt.Println()
		}
	}

}
