package stack

import "fmt"

//Item represents an item in the stack
type Item interface{}

//Stack represents a stack
type Stack struct {
	items []Item
}

//Push item into the stack
func (s *Stack) Push(item Item) {
	s.items = append(s.items, item)
}

//Pop and return the last item added
func (s *Stack) Pop() (Item, error) {
	n := len(s.items)
	if n == 0 {
		return 0, fmt.Errorf("Stack is empty")
	}
	item := s.items[n-1]
	s.items = s.items[:n-1]

	return item, nil
}

//Peek returns the top of stack
func (s Stack) Peek() (Item, error) {
	if len(s.items) == 0 {
		return 0, fmt.Errorf("Stack is empty")
	}
	return s.items[len(s.items)-1], nil
}

//Len returns the size of the stack
func (s Stack) Len() int {
	return len(s.items)
}
