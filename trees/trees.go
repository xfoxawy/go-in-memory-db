package trees

type Tree struct {
	Root *Node
	Size int
}

type Node struct {
	name   string
	parent *Node
	childs *Tree
}

func newTree() *Tree {
	return &Tree{&Node{}, 0}
}

func (t *Tree) addMainRoot(node *Node) *Tree {
	return &Tree{
		&Node{
			node.name,
			nil,
			nil,
		},
		t.Size + 1,
	}
}

func (t *Tree) insert(node *Node) *Tree {
	var tree *Tree
	if node.parent == nil {
		tree = t.addMainRoot(node)
	} else {
		tree = &Tree{
			&Node{
				node.name,
				node.parent,
				nil,
			},
			t.Size + 1,
		}
		t.Root.childs = tree
	}
	return tree
}

func (t *Tree) retrive() *Tree {
	return t
}
