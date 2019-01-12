package skiplist

import (
	"math/rand"
	"testing"
	"time"
)

var list *SkipList

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func TestInsert(t *testing.T) {
	list := New()

	list.Insert(time.Now().Unix())
	list.Insert(time.Now().Unix())
	list.Insert(time.Now().Unix())

	if list.Length() != 3 {
		t.Fatalf("Invalid Length after inseration, Length %v", list.Length())
	}

	t.Logf("List Length is %v", list.Length())
}

func TestGet(t *testing.T) {
	list := New()

	point := time.Now().Unix()
	list.Insert(point)
	p := list.Get(point)
	t.Logf("point %v", p)
	// list.Insert(point + 1)
	// list.Insert(point + 2)

	// if list.Length() != 3 {
	// t.Fatalf("Invalid Length after inseration, Length %v", list.Length())
	// }
}
