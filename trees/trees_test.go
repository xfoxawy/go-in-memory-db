package trees

import "testing"

var tree *Tree

func init() {
	tree = newTree()
}

func TestAddMainRoot(t *testing.T) {
	tree = tree.addMainRoot(&Node{"alaa", nil, nil})
	if tree.Size != 1 {
		t.Error("expected", 1, "got", tree.Size)
	}
}

func TestInsert(t *testing.T) {
	inserted_tree := tree.insert(&Node{"alaa", tree.Root, nil})
	if inserted_tree.Size != 2 {
		t.Error("expected", 2, "got", inserted_tree.Size)
	}
}
