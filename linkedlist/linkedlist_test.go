package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrepend(t *testing.T) {
	assert := assert.New(t)
	ll := LinkedList{}
	ll.Prepend(1)
	assert.NotNil(ll.head, "head should not be nil")
	assert.Equal(1, ll.head.data, "they should be equal")
	assert.Equal(1, ll.length, "they should be equal")
}

func TestDeleteValue(t *testing.T) {
	assert := assert.New(t)
	ll := LinkedList{}
	//Empty List
	err := ll.DeleteValue(1, false)
	assert.NotNil(err, "there should be error since list is empty")

	//Value not in list
	ll.Prepend(110)
	err = ll.DeleteValue(1, false)
	assert.NotNil(err, "there should be error since value not in list")
	ll.Prepend(120)
	err = ll.DeleteValue(121, true)
	assert.NotNil(err, "there should be error since value not in list")
	ll.Prepend(130)
	//Value in list
	err = ll.DeleteValue(130, true)
	assert.Nil(err, "there should be no error")
	err = ll.DeleteValue(110, true)
	assert.Nil(err, "there should be no error")

}

func TestLen(t *testing.T) {
	ll := LinkedList{}
	assert.Equal(t, 0, ll.Len(), "should be 0 initially.")
	ll.Prepend(1)
	ll.Prepend(2)
	ll.Prepend(3)
	assert.Equal(t, 3, ll.Len(), "should be equal")
	ll.DeleteValue(1, false)
	assert.Equal(t, 2, ll.Len(), "should be equal")
}
