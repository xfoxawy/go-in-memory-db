package skiplist

import (
	"math/rand"
	"sync"
	"time"
)

const DefaultMaxHeight = 32

type Node struct {
	key  int64
	next []*Node
}

func (n *Node) Height() int {
	return len(n.next) - 1
}

func (n *Node) Key() int64 {
	return n.key
}

func newNode(key int64, height int) *Node {
	return &Node{
		key:  key,
		next: make([]*Node, height+1),
	}
}

type SkipList struct {
	sentinel *Node
	length   int
	height   int
	stack    []*Node
	mutex    sync.RWMutex
}

func New() *SkipList {
	return &SkipList{
		sentinel: newNode(0, DefaultMaxHeight),
		length:   0,
		height:   0,
		stack:    make([]*Node, DefaultMaxHeight),
	}
}

// Find the node<T> u that precedes the value x in the skiplist.

func (l *SkipList) findPredNode(key int64) *Node {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	current := l.sentinel
	pointer := l.height

	for pointer >= 0 {
		for current.next[pointer] != nil && current.next[pointer].key < key {
			current = current.next[pointer] // go right
		}
		pointer-- // go down
	}
	return current
}

func (l *SkipList) Get(key int64) *Node {
	prevnode := l.findPredNode(key)
	return prevnode.next[0]
}

func (l *SkipList) findGT(key int64) *Node {
	return nil
}

func (l *SkipList) findLT(key int64) *Node {
	return nil
}

func (l *SkipList) Insert(key int64) {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	current := l.sentinel
	pointer := l.height

	for pointer >= 0 {
		for current.next[pointer] != nil && current.next[pointer].key < key {
			// move right
			current = current.next[pointer]
		}
		if current.next[pointer] != nil && current.next[pointer].key == key {
			return
		}
		pointer--
		if pointer > -1 {
			l.stack[pointer] = current // go down, store current
		}
	}

	level := l.generateLevel(DefaultMaxHeight)
	newNode := newNode(key, level)

	for l.height < newNode.Height() {
		l.height++
		l.stack[l.height] = l.sentinel // height incresed, inserting sential in the beginning og each node
	}

	for i := 0; i < newNode.Height(); i++ {
		newNode.next[i] = l.stack[i].next[i]
		l.stack[i].next[i] = newNode
	}
	l.length++
	return
}

func (l *SkipList) generateLevel(maxLevel int) int {

	rand.Seed(time.Now().UnixNano())
	r := float64(rand.Int63()) / (1 << 63)

	level := 1

	for level <= maxLevel && r > 0.5 {
		level++
		r = float64(rand.Int63()) / (1 << 63)
	}

	return level
}

func (l *SkipList) Length() int {
	return l.length
}
