package main

import (
	"container/list"
	"fmt"
)

type Stack struct {
	v *list.List
}

func (s *Stack) Push(value any) {
	s.v.PushBack(value)
}

func (s *Stack) Pop() any {
	ele := s.v.Back()
	if ele != nil {
		return s.v.Remove(ele)
	}
	return nil
}

func NewStack() *Stack {
	return &Stack{list.New()}
}

func main() {
	stack := NewStack()
	for i := 0; i < 5; i++ {
		stack.Push(i)
	}

	for ele := stack.Pop(); ele != nil; ele = stack.Pop() {
		fmt.Print(ele, ", ")
	}
}
