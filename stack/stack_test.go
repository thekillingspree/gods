package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPush(t *testing.T) {
	assert := assert.New(t)
	stack := Stack{}
	stack.Push(1)
	assert.NotEmpty(stack.items)
}

func TestPop(t *testing.T) {
	assert := assert.New(t)
	stack := Stack{}

	// Test 1 - Empty queue
	_, err1 := stack.Pop()
	assert.NotNil(err1, "there should be an error")

	// Test 2 - Queue with items
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	val, err2 := stack.Pop()
	assert.Nil(err2, "there should not be any error")
	assert.Equal(2, stack.Len(), "the length should be 2")
	assert.Equal(3, val, "they should be equal")
}

func TestPeek(t *testing.T) {
	assert := assert.New(t)
	stack := Stack{}

	// Test 1 - Empty queue
	_, err1 := stack.Peek()
	assert.NotNil(err1, "there should be an error")

	// Test 2 - Queue with items
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	val, err2 := stack.Peek()
	assert.Nil(err2, "there should not be any error")
	assert.Equal(3, val, "they should be equal")
}
