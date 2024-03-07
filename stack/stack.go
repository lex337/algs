package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	current_value *StackElement
}
type StackElement struct {
	value      int
	next_value *StackElement
}

type IStack interface {
	CreateStack()
	Push()
	Pop()
	Top()
	IsStackEmpty()
}

func (s *Stack) Pop() (int, error) {
	if &s.current_value == nil {
		return 0, errors.New("empty stack")
	}
	v := s.current_value.value
	s.current_value = s.current_value.next_value
	return v, nil
}

func CreateStack(init_value int) *Stack {
	return &Stack{current_value: &StackElement{value: init_value, next_value: nil}}
}

func (s *Stack) Push(v int) {
	s.current_value = &StackElement{value: v, next_value: s.current_value}
}

func main() {
	stack := CreateStack(10)
	for i := 0; i < 10; i++ {
		stack.Push(i)
	}
	fmt.Println(stack)
	for i := 0; i < 10; i++ {
		fmt.Println(stack.Pop())
	}
}
