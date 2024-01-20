package main

import (
	"container/list"
	"fmt"
)

func main() {
	a := list.New()
	e1 := a.PushBack(1)
	e2 := a.PushFront(2)

	a.InsertBefore(10, e1)
	a.InsertAfter(20, e2)

	for e := a.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, ", ")
	}

	fmt.Println("\n----------")
	for e := a.Back(); e != nil; e = e.Prev() {
		fmt.Print(e.Value, ", ")
	}
}
