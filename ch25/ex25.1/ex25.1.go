package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan int, 2)
	wg.Add(1)
	go power(&wg, ch)
	fmt.Println("data insert before")
	ch <- 5
	fmt.Println("data insert after")
	wg.Wait()
}

func power(wg *sync.WaitGroup, ch chan int) {
	fmt.Println("power function start")

	time.Sleep(time.Second)
	//num := <-ch
	time.Sleep(time.Second)
	//fmt.Printf("%d의 거듭제곱은 %d\n", num, num*num)
	fmt.Println("power function end")
	wg.Done()
}
