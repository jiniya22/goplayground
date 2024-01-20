package main

import (
	"container/list"
	"fmt"
)

type Queue struct {
	v *list.List
}

func (q *Queue) Push(val any) {
	q.v.PushBack(val)
}

func (q *Queue) Pop() any {
	ele := q.v.Front()
	if ele != nil {
		return q.v.Remove(ele)
	}
	return nil
}

func NewQueue() *Queue {
	return &Queue{list.New()}
}

func main() {
	queue := NewQueue()
	for i := 0; i < 5; i++ {
		queue.Push(i)
	}
	for ele := queue.Pop(); ele != nil; ele = queue.Pop() {
		fmt.Print(ele, ", ")
	}
}
