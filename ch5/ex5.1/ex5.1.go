package main

import "fmt"

func main() {
	var n = 12345.6789
	fmt.Print("1)")
	fmt.Println("2) n: ", n)
	fmt.Printf("3) n: %f => %.2f\n", n, n)
	fmt.Println("finish!!")

	fmt.Printf("%f, %f", 123456789.01234, 12341231234343123.56)
}
