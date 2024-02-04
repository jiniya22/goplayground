package main

import (
	"fmt"
	"sync"
	"time"
)

func power(wg *sync.WaitGroup, ch chan int) {
	for num := range ch {
		fmt.Printf("%d의 거듭제곱은 %d\n", num, num*num)
		time.Sleep(200 * time.Millisecond)
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(1)
	go power(&wg, ch)
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
	wg.Wait()
}
