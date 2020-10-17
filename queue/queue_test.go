package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnqueue(t *testing.T) {
	assert := assert.New(t)
	q := Queue{}
	q.Enqueue(1)
	assert.NotEmpty(q.items)
}

func TestDequeue(t *testing.T) {
	assert := assert.New(t)
	q := Queue{}

	// Test 1 - Empty queue
	_, err1 := q.Dequeue()
	assert.NotNil(err1, "should not be nil")

	// Test 2 - Queue with items
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	val, err2 := q.Dequeue()
	assert.Nil(err2, "there should not be any error")
	assert.Equal(2, q.Len(), "the length should be 2")
	assert.Equal(1, val, "they should be equal")
}

func TestPeek(t *testing.T) {
	assert := assert.New(t)
	q := Queue{}

	// Test 1 - Empty queue
	_, err1 := q.Peek()
	assert.NotNil(err1, "should not be nil")

	// Test 2 - Queue with items
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	val, err2 := q.Peek()
	assert.Nil(err2, "there should not be any error")
	assert.Equal(1, val, "they should be equal")
}
