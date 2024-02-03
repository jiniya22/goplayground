package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func sumAtoB(i, a, b int) {
	sum := 0
	for i := a; i <= b; i++ {
		sum += i
	}
	fmt.Printf("%d) %d 부터 %d 까지의 총 합 = %d\n", i, a, b, sum)
	wg.Done()
}
func main() {
	cnt := 10
	wg.Add(cnt)
	for i := 0; i < cnt; i++ {
		go sumAtoB(i, 1, 1000000)
	}
	wg.Wait()
}
