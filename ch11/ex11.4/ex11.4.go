package main

import "fmt"

func main() {
	ex1()
	ex2()
	ex3()
	ex4()
}

func ex1() {
	for i := 10; i > 0; i-- {
		fmt.Print(i, " ")
	}
	fmt.Println()
	printBr()
}

func ex2() {
	for i := 2; i < 10; i++ {
		if i >= 3 && i <= 6 {
			continue
		}
		for j := 1; j < 10; j++ {
			fmt.Println(i, "*", j, "=", i*j)
		}
		fmt.Println()
	}
	printBr()
}

func ex3() {
	for i := 1; i < 10; i += 2 {
		fmt.Println(i, "*", i, "=", i*i)
	}
	printBr()
}

func ex4() {
	for i := 5; i > 0; i-- {
		for j := i; j > 0; j-- {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func printBr() {
	fmt.Println("-------")
}
