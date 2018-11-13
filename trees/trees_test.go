package trees

import "testing"

var tree *Tree

func init() {
	tree = newTree()
}

var parent1 *Tree

func TestAddMainRoot(t *testing.T) {
	tree = tree.addMainRoot(&Node{"alaa1", nil, nil})
	parent1 = tree
	if tree.Size != 1 {
		t.Error("expected", 1, "got", tree.Size)
	}
}

func TestInsert(t *testing.T) {
	tree = tree.insert(&Node{"alaa2", tree.Root, nil})
	tree = tree.insert(&Node{"alaa2", nil, nil})
	if tree.Size != 3 {
		t.Error("expected", 3, "got", tree.Size)
	}
}

func TestRetrive(t *testing.T) {
	tree = tree.insert(&Node{"alaa3", parent1.Root, nil})

	if tree.Size != tree.retrive().Size {
		t.Error("expected", tree.Size, "got", tree.retrive().Size)
	}
}
