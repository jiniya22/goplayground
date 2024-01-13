package main

import "fmt"

type account struct {
	balance int
}

func withdrawFunc(a *account, amount int) {
	a.balance -= amount
}

func (a *account) withdrawMethod(amount int) {
	a.balance -= amount
}

func main() {
	a := &account{1000}
	withdrawFunc(a, 30)
	fmt.Printf("잔액: %d\n", a.balance)

	a.withdrawMethod(50)
	fmt.Printf("잔액: %d\n", a.balance)
}
