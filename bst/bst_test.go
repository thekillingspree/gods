package bst

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) {
	assert := assert.New(t)
	bst := BST{}
	bst.Insert(10)
	bst.Insert(2)
	bst.Insert(4)
	bst.Insert(15)
	bst.Insert(13)
	assert.Equal(10, bst.root.data, "should be equal")
	assert.Equal(2, bst.root.left.data, "should be equal")
	assert.Equal(15, bst.root.right.data, "should be equal")
}

func TestSearch(t *testing.T) {
	assert := assert.New(t)
	bst := BST{}
	node, err := bst.Search(10)
	assert.NotNil(err, "should not be nil")
	bst.Insert(10)
	//Only root
	node, err = bst.Search(10)
	assert.Nil(err, "should be nil")
	assert.NotNil(node, "should not be nil")
	assert.Equal(10, node.data, "should not be nil")
	bst.Insert(1)
	bst.Insert(14)
	//Left
	node, err = bst.Search(1)
	assert.Nil(err, "should be nil")
	assert.NotNil(node, "should not be nil")
	assert.Equal(1, node.data, "should not be nil")
	//Right
	node, err = bst.Search(14)
	assert.Nil(err, "should be nil")
	assert.NotNil(node, "should not be nil")
	assert.Equal(14, node.data, "should not be nil")
}

func TestInOrder(t *testing.T) {
	assert := assert.New(t)
	bst := BST{}
	bst.Insert(10)
	bst.Insert(1)
	bst.Insert(15)
	inOrder := bst.InOrder()
	assert.Equal([]int{1, 10, 15}, inOrder, "should be equal")

}
