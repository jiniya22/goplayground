package main

import (
	"fmt"
	"sync"
	"time"
)

func power(wg *sync.WaitGroup, nums chan int, quit chan bool) {
	for {
		select {
		case num := <-nums:
			fmt.Printf("%d의 거듭제곱은 %d\n", num, num*num)
			time.Sleep(time.Second)
		case <-quit:
			wg.Done()
			return
		}
	}
}

func main() {
	var wg sync.WaitGroup
	nums := make(chan int)
	quit := make(chan bool)
	wg.Add(1)
	go power(&wg, nums, quit)
	for i := 1; i <= 5; i++ {
		nums <- i
	}
	quit <- true
	wg.Wait()
}
