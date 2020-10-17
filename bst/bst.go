package bst

import (
	"fmt"
)

// BST represents a binary search tree
type BST struct {
	root *Node
}

// Node represents a node in the tree
type Node struct {
	data  int
	left  *Node
	right *Node
}

//Insert will insert a node to the tree
func (bst *BST) Insert(data int) {
	if bst.root == nil {
		bst.root = &Node{data: data}
		return
	}
	bst.root.insert(data)
}

func (node *Node) insert(data int) {
	if node.data < data {
		if node.right == nil {
			node.right = &Node{data: data}
		} else {
			node.right.insert(data)
		}
	} else {
		if node.left == nil {
			node.left = &Node{data: data}
		} else {
			node.left.insert(data)
		}
	}
}

//InOrder will return the string representation of the in-order
//traversal
func (bst BST) InOrder() []int {
	return bst.root.inOrder()
}

func (node *Node) inOrder() []int {
	if node != nil {
		sl := node.left.inOrder()
		order := append(sl, node.data)
		sr := node.right.inOrder()
		return append(order, sr...)
	}
	return []int{}
}

//Search will search for data in the tree
func (bst BST) Search(data int) (Node, error) {
	return bst.root.search(data)
}

func (node *Node) search(data int) (Node, error) {
	if node == nil {
		return Node{}, fmt.Errorf("Not found")
	}

	if node.data > data {
		return node.left.search(data)
	} else if node.data < data {
		return node.right.search(data)
	} else {
		return *node, nil
	}
}

// func main() {
// 	fmt.Println("Binary Search")
// 	root := &Node{data: 10}
// 	root.insert(20)
// 	root.insert(15)
// 	root.insert(1)
// 	root.insert(4)
// 	root.insert(24)
// 	fmt.Println(root)
// 	val, err := root.search(4)
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println("Searched: ", val)
// 	}
// }
