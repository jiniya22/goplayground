package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	go printTickEverySecond(ctx)

	time.Sleep(10 * time.Second)
	cancel()
	wg.Wait()
}

func printTickEverySecond(ctx context.Context) {
	tick := time.Tick(time.Second)
	for {
		select {
		case <-ctx.Done():
			wg.Done()
			return
		case <-tick:
			fmt.Println("Tick!")
		}
	}
}
