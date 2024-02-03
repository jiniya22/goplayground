package main

import (
	"fmt"
	"sync"
	"time"
)

var mutex sync.Mutex

type Account struct {
	Balance int
}

func depositAndWithdraw(account *Account, n int) {
	mutex.Lock()
	defer mutex.Unlock()
	if account.Balance < 0 {
		panic(fmt.Sprintf("계좌 잔고는 0 이하일 수 없습니다: 잔고 %d원\n", account.Balance))
	}
	amount := n * 1000
	account.Balance += amount
	time.Sleep(time.Millisecond)
	account.Balance -= amount
}

func main() {
	var wg sync.WaitGroup
	cnt := 10

	account := &Account{0}
	wg.Add(cnt)
	for i := 0; i < cnt; i++ {
		a := i
		go func() {
			for {
				depositAndWithdraw(account, a)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
