package main

import (
	"fmt"
	"time"
)

func printAlphabets() {
	alphabets := []rune{'A', 'B', 'C', 'D', 'E', 'F'}
	for _, alphabet := range alphabets {
		time.Sleep(200 * time.Millisecond)
		fmt.Printf("%c ", alphabet)
	}
}

func printNumbers() {
	for i := 0; i < 7; i++ {
		time.Sleep(300 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}

func main() {
	go printAlphabets()
	go printNumbers()

	fmt.Println("main 1")
	time.Sleep(600 * time.Millisecond)
	fmt.Println("main 2")
}
