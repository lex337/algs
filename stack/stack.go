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
	if s.current_value == nil {
		return 0, errors.New("empty stack")
	}
	v := s.current_value.value
	s.current_value = s.current_value.next_value
	return v, nil
}

func (s *Stack) Top() (int, error) {
	if s.current_value == nil {
		return 0, errors.New("empty stack")
	}
	return s.current_value.value, nil
}

func (s *Stack) Push(v int) {
	s.current_value = &StackElement{value: v, next_value: s.current_value}
}

func CreateStack(init_value int) *Stack {
	return &Stack{current_value: &StackElement{value: init_value, next_value: nil}}
}

func (s *Stack) IsStackEmpty() bool {
	return s.current_value == nil
}

func StockStackSpan(quotes []int) []int {
	spans := make([]int, len(quotes))
	stack := CreateStack(0)
	for i := 1; i < len(quotes); i++ {
		top, _ := stack.Top()
		for !stack.IsStackEmpty() && quotes[top] <= quotes[i] {
			stack.Pop()
		}
		if stack.IsStackEmpty() {
			spans[i]++
		} else {
			top, _ := stack.Top()
			spans[i] = i - top
		}
		stack.Push(i)
	}
	return spans
}

func main() {
	a := []int{7, 11, 8, 6, 3, 8, 9}
	fmt.Println(StockStackSpan(a))
}
