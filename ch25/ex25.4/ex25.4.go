package main

import (
	"fmt"
	"sync"
	"time"
)

func power(wg *sync.WaitGroup, ch chan int) {
	tick := time.Tick(time.Second)
	terminate := time.After(10 * time.Second)

	for {
		t := time.Now().Format(time.RFC3339)
		select {
		case <-tick:
			fmt.Printf("%s | Tick!\n", t)
		case <-terminate:
			fmt.Printf("%s | power function terminate!", t)
			wg.Done()
			return
		case num := <-ch:
			fmt.Printf("%s | %d의 거듭제곱은 %d\n", t, num, num*num)
			time.Sleep(time.Second)
		}
	}
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)
	go power(&wg, ch)
	for i := 0; i < 10; i++ {
		ch <- i
	}
	wg.Wait()
}
