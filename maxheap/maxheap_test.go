package maxheap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) {
	assert := assert.New(t)
	mh := MaxHeap{}
	assert.Equal(0, len(mh.contents), "should be 0 initially")
	mh.Insert(100)
	assert.Equal(1, len(mh.contents), "should be equal")
}

func TestGetLargest(t *testing.T) {
	assert := assert.New(t)
	mh := MaxHeap{}
	val, err := mh.GetLargest()
	assert.NotNil(err, "should not be nil, since the heap is empty")
	mh.Insert(1)
	mh.Insert(100)
	mh.Insert(-100)
	val, err = mh.GetLargest()
	assert.Nil(err, "should be nil")
	assert.Equal(100, val, "should get the largest element in heap")
}

func TestExtractLargest(t *testing.T) {
	assert := assert.New(t)
	mh := MaxHeap{}
	val, err := mh.ExtractLargest()
	assert.NotNil(err, "should not be nil. Heap is empty.")
	mh.Insert(1)
	mh.Insert(200)
	mh.Insert(100)
	mh.Insert(-100)
	mh.Insert(10)
	val, err = mh.ExtractLargest()
	assert.Nil(err, "should be nil")
	assert.Equal(200, val, "should get the largest element in heap")
	val, err = mh.ExtractLargest()
	assert.Nil(err, "should be nil")
	assert.Equal(100, val, "should get the largest element in heap")
	val, err = mh.ExtractLargest()
	assert.Nil(err, "should be nil")
	assert.Equal(10, val, "should get the largest element in heap")
	assert.Equal(2, len(mh.contents), "should have decreased")
}

func TestLen(t *testing.T) {
	assert := assert.New(t)
	mh := MaxHeap{}
	assert.Equal(0, mh.Len(), "should be 0 initially")
	mh.Insert(10)
	mh.Insert(100)
	mh.Insert(-100)
	assert.Equal(3, mh.Len(), "should be 3")
	mh.ExtractLargest()
	assert.Equal(2, mh.Len(), "should've decreased by 1")
}
