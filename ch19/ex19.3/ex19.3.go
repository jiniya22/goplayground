package main

import "fmt"

type account struct {
	balance int
}

func (a *account) withdrawPointer(amount int) {
	a.balance -= amount
}

func (a account) withdrawRetureValue(amount int) account {
	a.balance -= amount
	return a
}

func main() {
	a := account{1000}

	(&a).withdrawPointer(50)
	fmt.Printf("a 잔액: %d\n", a.balance)

	fmt.Println("------------")

	a2 := a.withdrawRetureValue(30)
	fmt.Printf("a 잔액: %d\n", a.balance)
	fmt.Printf("a2 잔액: %d\n", a2.balance)
}
