package trees

type Tree struct {
	node *Node
}

type Node struct {
	id    int
	name  string
	left  *Node
	right *Node
}

func newTree() *Tree {
	return &Tree{nil}
}

func addRoot(node *Node) *Tree {
	newTree := &Tree{
		&Node{
			node.id,
			node.name,
			nil,
			nil,
		},
	}
	return newTree
}

func (n *Node) insert(node *Node) *Node {
	var newNode *Node
	if node == nil {
		newNode = addRoot(node).node
	}
	if node.id > n.id {
		if n.right == nil {
			n.right = &Node{node.id, node.name, nil, nil}
			newNode = n.right
		} else {
			newNode = n.right.insert(node)
		}
	}
	if node.id < n.id {
		if n.left == nil {
			n.left = &Node{node.id, node.name, nil, nil}
			newNode = n.left
		} else {
			newNode = n.left.insert(node)
		}
	}
	return newNode
}

func (t *Tree) retrive() *Tree {
	return t
}
