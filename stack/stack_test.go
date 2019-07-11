package stack

import (
	"testing"
)

var stack *Stack

func init() {
	stack = NewStack()
}

func TestPush(t *testing.T) {
	stack.Push("string here")
	if stack.Size() == 0 {
		t.Errorf("expected size larger than one")
	}
}

func TestPop(t *testing.T) {
	stack := NewStack()
	if stack.Pop() != "" {
		t.Error("expected", "empty string", "got", stack.Pop())
	}
	stack.Push("string here")
	stack.Push("string here2")
	stack.Push("string here3")
	sizeBefore := stack.Size()
	here3 := stack.Pop()
	sizeAfter := stack.Size()
	if sizeAfter != sizeBefore-1 {
		t.Error("expected", sizeBefore-1, "got", sizeAfter)
	}

	if here3 != "string here3" {
		t.Error("expected", "string here3", "got", here3)
	}
}

func TestSize(t *testing.T) {
	stack.Push("string here")
	if stack.Size() == 0 {
		t.Errorf("expected size larger than one")
	}
}

func testIsEmpty(t *testing.T) {
	stack = NewStack()
	if stack.IsEmpty() {
		t.Errorf("expected stack is empty")
	}

	stack.Push("string here")
	if stack.IsEmpty() {
		t.Errorf("expected stack is not empty")
	}

}
