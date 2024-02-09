package main

import (
	"context"
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	ctx := context.WithValue(context.Background(), "num", 5)
	go power(ctx)
	wg.Wait()
}

func power(ctx context.Context) {
	if v := ctx.Value("num"); v != nil {
		num := v.(int)
		fmt.Printf("%d의 거듭제곱은 %d\n", num, num*num)
	}
	wg.Done()
}
