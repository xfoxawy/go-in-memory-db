package linkedlist

import "testing"

var list *LinkedList

func TestPush(t *testing.T) {
	list = NewList()
	list.Push("string here")
	if list.Length == 0 {
		t.Error("expected", "length > 0", "got", "0")
	}
}

func TestPop(t *testing.T) {
	list = NewList()
	_, err := list.Pop()
	if err == nil {
		t.Errorf("expected error linked list is empty")
	}
	list.Push("string1")
	list.Push("string2")
	list.Push("string3")
	poped, err := list.Pop()
	if poped.Value != "string3" {
		t.Error("expected", "string3", "got", poped.Value)
	}
}

func TestShift(t *testing.T) {
	list = NewList()
	list.Shift("string1")
	list.Shift("string2")
	list.Shift("string3")
	if list.Start.Value != "string3" {
		t.Error("expected", "string3", "got", list.Start.Value)
	}
}

func TestUnshift(t *testing.T) {
	list = NewList()
	_, err := list.Unshift()
	if err == nil {
		t.Errorf("expected error linked list is empty")
	}
	list.Shift("string1")
	list.Shift("string2")
	list.Shift("string3")
	shifted, _ := list.Unshift()
	if shifted.Value != "string3" || list.Start.Value != "string2" {
		t.Error("expected", "string3", "got", shifted.Value)
	}
}

func TestRemove(t *testing.T) {
	list = NewList()
	err := list.Remove("string here")
	if err == nil {
		t.Errorf("expected error linked list is empty")
	}
	list.Push("string1")
	list.Push("string2")
	list.Push("string3")
	res := list.Remove("string1")
	if res != nil {
		t.Errorf("expect error will be equal null")
	}

}

func TestUnlink(t *testing.T) {
	list = NewList()
	err := list.Unlink(1)
	if err == nil {
		t.Errorf("expected error linked list is empty or step not exist")
	}
	list.Push("string1")
	list.Push("string2")
	list.Push("string3")
	res := list.Unlink(1)
	if res != nil {
		t.Errorf("expect error will be equal null")
	}

}

func TestSeek(t *testing.T) {
	list = NewList()
	_, err := list.Seek(1)
	if err == nil {
		t.Errorf("expected error linked list is empty or step not exist")
	}
	list.Push("string1")
	list.Push("string2")
	list.Push("string3")
	res, _ := list.Seek(1)
	if res != "string1" {
		t.Error("expected", "string1", "Got", res)
	}
}
