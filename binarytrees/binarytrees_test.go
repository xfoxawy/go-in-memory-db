package binarytrees

import (
	"testing"

	"github.com/go-in-memory-db/linkedlist"
)

func TestAddRoot(t *testing.T) {
	BinaryTree := addRoot(&Node{0, "alaa1", nil, nil, nil})
	if BinaryTree.node.id != 0 {
		t.Error("expected", 0, "got", BinaryTree.node.id)
	}
}

var lastBinaryTree *BinaryTree
var lastDeepInLeft *Node
var lastDeepInRight *Node

// Test will handle root with deepest element in right and in left only
func TestInsert(t *testing.T) {
	BinaryTree := addRoot(&Node{5, "node", nil, nil, nil}) // root
	node := BinaryTree.node
	node.insert(&Node{9, "node1", nil, nil, nil})                // first right
	node.insert(&Node{12, "node2", nil, nil, nil})               // right for node1
	node.insert(&Node{3, "node3", nil, nil, nil})                // left for node
	node.insert(&Node{7, "node4", nil, nil, nil})                // left for node1
	deepInRight := node.insert(&Node{8, "node5", nil, nil, nil}) // right for node4
	node.insert(&Node{3, "node6", nil, nil, nil})                // left for root
	deepInLeft := node.insert(&Node{4, "node7", nil, nil, nil})  // right for node6

	if node.right.left.right.name != deepInRight.name {
		t.Error("expected", deepInRight.name, "got", node.right.left.right.name)
	}

	if node.left.right.name != deepInLeft.name {
		t.Error("expected", deepInLeft.name, "got", node.left.right.name)
	}
	lastBinaryTree = BinaryTree
	lastDeepInLeft = deepInLeft
	lastDeepInRight = deepInRight
}

func TestRetrive(t *testing.T) {
	BinaryTree := lastBinaryTree.retrive()
	if BinaryTree.node.right.left.right.name != lastDeepInRight.name {
		t.Error("expected", lastDeepInRight.name, "got", BinaryTree.node.right.left.right.name)
	}

	if BinaryTree.node.left.right.name != lastDeepInLeft.name {
		t.Error("expected", lastDeepInLeft.name, "got", BinaryTree.node.left.right.name)
	}
}

func TestFind(t *testing.T) {
	// root node
	if node, _ := lastBinaryTree.node.find(5); node.name != "node" {
		t.Error("expected", "node", "got", node.name)
	}
	// test not found id
	if _, err := lastBinaryTree.node.find(50); err == nil {
		t.Errorf("expect error here")
	}
	// test deep element in right
	if node, _ := lastBinaryTree.node.find(lastDeepInRight.id); node.name != lastDeepInRight.name {
		t.Error("expected", lastDeepInRight.name, "got", node.name)
	}
	// test deep element in left
	if node, _ := lastBinaryTree.node.find(lastDeepInLeft.id); node.name != lastDeepInLeft.name {
		t.Error("expected", lastDeepInLeft.name, "got", node.name)
	}

}

func TestIsLeaf(t *testing.T) {
	testLeafTree := addRoot(&Node{5, "node", nil, nil, nil}) // root
	testLeafNode := testLeafTree.node
	if testLeafNode.isLeaf() != true {
		t.Error("expected", true, "got", false)
	}
	testLeafNode.insert(&Node{9, "node1", nil, nil, nil}) // first right
	if testLeafNode.isLeaf() != false {
		t.Error("expected", false, "got", true)
	}
}

func TestHasLeft(t *testing.T) {
	testHasLeftTree := addRoot(&Node{5, "node", nil, nil, nil}) // root
	testhasLeftNode := testHasLeftTree.node
	if testhasLeftNode.hasLeft() != false {
		t.Error("expected", false, "got", true)
	}
	testhasLeftNode.insert(&Node{3, "node6", nil, nil, nil}) // left for root
	if testhasLeftNode.hasLeft() != true {
		t.Error("expected", true, "got", false)
	}
}

func TestHasRight(t *testing.T) {
	testHasRightTree := addRoot(&Node{5, "node", nil, nil, nil}) // root
	testhasRightNode := testHasRightTree.node
	if testhasRightNode.hasRight() != false {
		t.Error("expected", false, "got", true)
	}
	testhasRightNode.insert(&Node{9, "node1", nil, nil, nil}) // first right
	if testhasRightNode.hasRight() != true {
		t.Error("expected", true, "got", false)
	}
}

func TestMin(t *testing.T) {
	lastBinaryTree.node.insert(&Node{1, "min node", nil, nil, nil})
	if lastBinaryTree.node.min().id != 1 {
		t.Error("expected", 3, "got", lastBinaryTree.node.min().id)
	}
}

func TestMax(t *testing.T) {
	lastBinaryTree.node.insert(&Node{55, "max node", nil, nil, nil})
	if lastBinaryTree.node.max().id != 55 {
		t.Error("expected", 55, "got", lastBinaryTree.node.max().id)
	}
}

func TestGetNodePayload(t *testing.T) {
	list := linkedlist.NewList()
	list.Push("string here")
	list.Push("string here1")
	list.Push("string here2")
	lastBinaryTree.node.insert(&Node{56, "max node1", list, nil, nil})
	node, _ := lastBinaryTree.node.find(56)
	payload := node.getNodePayload()
	if payload != nil && payload.Length != 3 {
		t.Error("expected", 3, "got", payload.Length)
	}
}
