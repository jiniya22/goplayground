package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func printAlphabets() {
	alphabets := []rune{'A', 'B', 'C', 'D', 'E', 'F'}
	for _, alphabet := range alphabets {
		time.Sleep(200 * time.Millisecond)
		fmt.Printf("%c ", alphabet)
	}
	wg.Done()
}

func printNumbers() {
	for i := 0; i < 7; i++ {
		time.Sleep(300 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
	wg.Done()
}

func main() {
	wg.Add(2)
	go printAlphabets()
	go printNumbers()
	wg.Wait()
	fmt.Println("main 2")
}
