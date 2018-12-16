package binarytrees

import (
	"errors"

	"github.com/go-in-memory-db/linkedlist"
)

// BinaryTree root
type BinaryTree struct {
	node *Node
}

// Node struct
type Node struct {
	id      int
	name    string
	payload *linkedlist.LinkedList
	left    *Node
	right   *Node
}

func newBinaryTree() *BinaryTree {
	return &BinaryTree{nil}
}

func createNode(id int, name string, payload *linkedlist.LinkedList) *Node {
	return &Node{
		id,
		name,
		payload,
		nil,
		nil,
	}
}

func addRoot(node *Node) *BinaryTree {
	newBinaryTree := &BinaryTree{
		&Node{
			node.id,
			node.name,
			linkedlist.NewList(),
			nil,
			nil,
		},
	}
	return newBinaryTree
}

func (n *Node) insert(node *Node) *Node {
	var newNode *Node
	if node == nil {
		newNode = addRoot(node).node
	}
	if node.id > n.id {
		if n.right == nil {
			n.right = &Node{node.id, node.name, node.payload, nil, nil}
			newNode = n.right
		} else {
			newNode = n.right.insert(node)
		}
	}
	if node.id < n.id {
		if n.left == nil {
			n.left = &Node{node.id, node.name, node.payload, nil, nil}
			newNode = n.left
		} else {
			newNode = n.left.insert(node)
		}
	}
	return newNode
}

func (t *BinaryTree) retrive() *BinaryTree {
	return t
}

func (n *Node) find(id int) (*Node, error) {
	if id == n.id {
		return n, nil
	}
	if id > n.id && n.hasRight() {
		return n.right.find(id)
	}
	if id < n.id && n.hasLeft() {
		return n.left.find(id)
	}
	return n, errors.New("this id not found")
}

func (n *Node) isLeaf() bool {
	return !n.hasLeft() && !n.hasRight()
}

func (n *Node) hasLeft() bool {
	return n.left != nil
}

func (n *Node) hasRight() bool {
	return n.right != nil
}

func (n *Node) min() *Node {
	if !n.hasLeft() {
		return n
	}
	return n.left.min()
}

func (n *Node) max() *Node {
	if !n.hasRight() {
		return n
	}
	return n.right.max()
}

func (n *Node) getNodePayload() *linkedlist.LinkedList {
	return n.payload
}
