package trees

import (
	"testing"
)

func TestAddRoot(t *testing.T) {
	tree := addRoot(&Node{0, "alaa1", nil, nil})
	if tree.node.id != 0 {
		t.Error("expected", 0, "got", tree.node.id)
	}
}

var lastTree *Tree
var lastDeepInLeft *Node
var lastDeepInRight *Node

// Test will handle root with deepest element in right and in left only
func TestInsert(t *testing.T) {
	tree := addRoot(&Node{5, "node", nil, nil}) // root
	node := tree.node
	node.insert(&Node{9, "node1", nil, nil})                // first right
	node.insert(&Node{12, "node2", nil, nil})               // right for node1
	node.insert(&Node{3, "node3", nil, nil})                // left for node
	node.insert(&Node{7, "node4", nil, nil})                // left for node1
	deepInRight := node.insert(&Node{8, "node5", nil, nil}) // right for node4
	node.insert(&Node{3, "node6", nil, nil})                // left for root
	deepInLeft := node.insert(&Node{4, "node7", nil, nil})  // right for node6

	if node.right.left.right.name != deepInRight.name {
		t.Error("expected", deepInRight.name, "got", node.right.left.right.name)
	}

	if node.left.right.name != deepInLeft.name {
		t.Error("expected", deepInLeft.name, "got", node.left.right.name)
	}
	lastTree = tree
	lastDeepInLeft = deepInLeft
	lastDeepInRight = deepInRight
}

func TestRetrive(t *testing.T) {
	tree := lastTree.retrive()
	if tree.node.right.left.right.name != lastDeepInRight.name {
		t.Error("expected", lastDeepInRight.name, "got", tree.node.right.left.right.name)
	}

	if tree.node.left.right.name != lastDeepInLeft.name {
		t.Error("expected", lastDeepInLeft.name, "got", tree.node.left.right.name)
	}
}
