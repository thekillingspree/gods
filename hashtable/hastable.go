package hashtable

import "fmt"

// ArraySize is the size of initial array of the hashtable
const ArraySize = 1

// HashTable represents Hash table
type HashTable struct {
	array [ArraySize]*Bucket
	size  int
}

// Insert will insert a key-value to hashtable
func (h *HashTable) Insert(key string, value interface{}) {
	index := hash(key)
	wasReplaced := h.array[index].insert(key, value)
	if !wasReplaced {
		h.size++
	}
}

// Search will search for a key in hashtable
func (h *HashTable) Search(key string) bool {
	index := hash(key)
	return h.array[index].search(key)
}

// Get will return the value for a key
func (h *HashTable) Get(key string) (interface{}, error) {
	index := hash(key)
	val, err := h.array[index].get(key)
	if err != nil {
		return nil, err
	}
	return val, nil
}

// Delete will remove a key from the hash table
func (h *HashTable) Delete(key string) error {
	index := hash(key)
	err := h.array[index].delete(key)
	if err == nil {
		h.size--
	}
	return err
}

// Len will return the number of items in the hash table
func (h HashTable) Len() int {
	return h.size
}

// hash function (for string)
func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}

	return sum % ArraySize
}

// Bucket is the linkedList
type Bucket struct {
	head   *bucketNode
	length int
}

type bucketNode struct {
	key   string
	value interface{}
	next  *bucketNode
}

//insert will insert a key-value pair into the bucket and returns true if the value was replaced
func (b *Bucket) insert(key string, value interface{}) bool {
	if !b.search(key) {
		newNode := &bucketNode{key: key, value: value}
		newNode.next = b.head
		b.head = newNode
		b.length++
		return false
	}
	b.delete(key)
	b.insert(key, value)
	return true
}

//search will search the bucket for a key
func (b *Bucket) search(key string) bool {
	t := b.head
	for t != nil {
		if t.key == key {
			return true
		}
		t = t.next
	}

	return false
}

func (b *Bucket) get(key string) (interface{}, error) {
	t := b.head
	for t != nil {
		if t.key == key {
			return t.value, nil
		}
		t = t.next
	}

	return nil, fmt.Errorf("Key not Found")
}

func (b *Bucket) delete(key string) error {

	if b.length == 0 {
		return fmt.Errorf("Key not found")
	}

	if b.head.key == key {
		b.head = b.head.next
		b.length--
		return nil
	}

	t := b.head
	for t.next != nil {
		if t.next.key == key {
			t.next = t.next.next
			b.length--
			return nil
		}
		t = t.next
	}

	return fmt.Errorf("Key not found")
}

// InitHashTable will initialize the hashtable
func InitHashTable() *HashTable {
	result := &HashTable{}
	for i := range result.array {
		result.array[i] = &Bucket{}
	}

	return result
}
