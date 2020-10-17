package linkedlist

import (
	"fmt"
	"strconv"
)

type node struct {
	data int
	next *node
}

//LinkedList struct represents a linked list
type LinkedList struct {
	head   *node
	length int
}

//Prepend adds a node at the first position of linked list
func (l *LinkedList) Prepend(data int) {
	n := &node{data: data}
	temp := l.head
	l.head = n
	l.head.next = temp
	l.length++
}

func (l LinkedList) String() string {
	str := "LinkedList (" + strconv.Itoa(l.length) + "): "
	t := l.head
	for t != nil {
		str += strconv.Itoa(t.data)
		if t.next != nil {
			str += "->"
		}
		t = t.next
	}
	return str
}

//DeleteValue deletes an item from linkedlist
func (l *LinkedList) DeleteValue(data int, displayAfterDeleting bool) error {
	if l.length == 0 {
		return fmt.Errorf("Cannot delete. LinkedList is empty")
	}

	if l.length == 1 && l.head.data != data {
		return fmt.Errorf("Value not found in LinkedList")
	}

	if l.head.data == data {
		l.head = l.head.next
		l.length--
		if displayAfterDeleting {
			fmt.Println(l)
		}
		return nil
	}

	toDelete := l.head
	for toDelete.next.data != data {
		if toDelete.next.next == nil {
			return fmt.Errorf("Value not found in LinkedList")
		}
		toDelete = toDelete.next
	}

	toDelete.next = toDelete.next.next
	l.length--
	if displayAfterDeleting {
		fmt.Println(l)
	}
	return nil
}

//Len returns the length of the linked list
func (l LinkedList) Len() int {
	return l.length
}
