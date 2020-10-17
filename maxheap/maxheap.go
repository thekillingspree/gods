package maxheap

import "fmt"

// MaxHeap is a struct representing a Max-Heap
type MaxHeap struct {
	contents []int
}

// Insert will insert an element to the heap
func (h *MaxHeap) Insert(value int) {
	h.contents = append(h.contents, value)
	h.heapifyUp(len(h.contents) - 1)
}

//GetLargest will return the largest element in the heap
func (h MaxHeap) GetLargest() (int, error) {
	if len(h.contents) == 0 {
		return 0, fmt.Errorf("heap is empty")
	}
	return h.contents[0], nil
}

// ExtractLargest will extract the largest element in the heap
func (h *MaxHeap) ExtractLargest() (int, error) {

	if len(h.contents) == 0 {
		return 0, fmt.Errorf("heap is empty")
	}

	extracted := h.contents[0]
	lastIdx := len(h.contents) - 1
	h.contents[0] = h.contents[lastIdx]
	h.contents = h.contents[:lastIdx]
	h.heapifyDown(0)
	return extracted, nil
}

//Len will return the number of elements in the heap
func (h MaxHeap) Len() int {
	return len(h.contents)
}

func (h *MaxHeap) heapifyUp(pos int) {
	for h.contents[parentIndex(pos)] < h.contents[pos] {
		h.swap(parentIndex(pos), pos)
		pos = parentIndex(pos)
	}
}

func (h *MaxHeap) heapifyDown(pos int) {
	current := pos
	lastIdx := len(h.contents) - 1
	l, r := leftChild(current), rightChild(current)
	childToCompare := 0
	for l <= lastIdx {
		if l == lastIdx || h.contents[l] > h.contents[r] {
			childToCompare = l
		} else if h.contents[r] > h.contents[current] {
			childToCompare = r
		}

		if h.contents[childToCompare] > h.contents[current] {
			h.swap(childToCompare, current)
			current = childToCompare
			l, r = leftChild(current), rightChild(current)
		} else {
			return
		}
	}
}

func parentIndex(i int) int {
	return (i - 1) / 2
}

func leftChild(parent int) int {
	return parent*2 + 1
}

func rightChild(parent int) int {
	return parent*2 + 2
}

func (h *MaxHeap) swap(i1, i2 int) {
	h.contents[i1], h.contents[i2] = h.contents[i2], h.contents[i1]
}
