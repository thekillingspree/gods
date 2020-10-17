package queue

import "fmt"

type Item interface{}

//Queue represents a queue
type Queue struct {
	items []Item
}

//Enqueue adds an item to the queue
func (q *Queue) Enqueue(item Item) {
	q.items = append(q.items, item)
}

//Dequeue removes the first item from the queue and returns it.
func (q *Queue) Dequeue() (Item, error) {
	n := len(q.items)
	if n == 0 {
		return 0, fmt.Errorf("Queue is Empty")
	}

	result := q.items[0]
	q.items = q.items[1:]
	return result, nil
}

//Peek returns the first time in the queue.
func (q *Queue) Peek() (Item, error) {
	n := len(q.items)
	if n == 0 {
		return 0, fmt.Errorf("Queue is Empty")
	}

	return q.items[0], nil
}

// Len returns the number of items in the queue.
func (q Queue) Len() int {
	return len(q.items)
}
