package linkedlist

import (
	"errors"
)

type Element struct {
	Value string
	Next  *Element
}

// note LinkedList starts from zero index
type LinkedList struct {
	Start  *Element
	end    *Element
	Length int
}

func NewList() *LinkedList {
	return &LinkedList{Length: 0}
}

func (l *LinkedList) Push(v string) {
	e := Element{Value: v}
	l.append(&e)
}

func (l *LinkedList) append(e *Element) {
	if l.Length == 0 {
		l.Start = e
		l.end = l.Start
	} else {
		end := l.end
		end.Next = e
		l.end = e
	}
	l.Length++
}

func (l *LinkedList) Pop() (*Element, error) {
	if l.Length == 0 {
		return nil, errors.New("LinkedList is empty")
	}

	if l.Length == 1 {
		popped := l.Start
		l.Start = nil
		l.end = nil
		l.Length--
		return popped, nil
	} else {
		counter := l.Length
		pointer := l.Start

		for counter != 2 {
			pointer = pointer.Next
			counter--
		}

		l.end = pointer
		popped := pointer.Next
		l.end.Next = nil
		l.Length--

		return popped, nil
	}
}

// should push a new element in front
// ex 1 -> 2 -> 3 .
// shift(8)
// 8 -> 1 -> 2 ->3
func (l *LinkedList) Shift(v string) {
	if l.Length == 0 {
		e := Element{Value: v}
		l.append(&e)
	} else {
		start := l.Start
		new_e := &Element{Value: v, Next: start}
		l.Start = new_e
	}

	l.Length++
}

// 8 -> 1 -> 2 ->3
// unshift(8)
//  1 -> 2 -> 3
func (l *LinkedList) Unshift() (*Element, error) {
	if l.Length == 0 {
		return nil, errors.New("LinkedList is empty")
	}
	old_start := l.Start
	l.Start = old_start.Next
	l.Length--
	return old_start, nil
}

// removes an element
// 1->2->3
// remove(2)
// 1->3
func (l *LinkedList) Remove(value string) error {
	if l.Length == 0 {
		return errors.New("LinkedList is empty")
	}
	if l.Start.Value == value {
		l.Unshift()
		return nil
	}
	if l.end.Value == value {
		l.Pop()
		return nil
	}
	current := l.Start
	last := l.Start
	for current.Next != nil {
		if current.Value == value {
			last.Next = current.Next
			current.Next = nil
			break
		}
		last = current
		current = current.Next
	}
	l.Length--
	return nil
}

// remove element by step
// 4 -> 5 ->6
// remove(2)
// 4 -> 5
func (l *LinkedList) Unlink(step int) error {
	if l.Length == 0 || l.Length < step {
		return errors.New("LinkedList is empty OR Step Not Exist")
	}
	if step == 1 {
		l.Unshift()
		return nil
	}
	if step == l.Length {
		l.Pop()
		return nil
	}
	current := l.Start
	last := l.Start
	i := 1
	for current.Next != nil {
		if step == i {
			last.Next = current.Next
			current.Next = nil
			break
		}
		last = current
		current = current.Next
		i++
	}
	l.Length--
	return nil
}

// 4->5->6
// seek(1)
// 5
// seek(0)
// 4
func (l *LinkedList) Seek(step int) (string, error) {
	if l.Length == 0 || l.Length < step {
		return "", errors.New("LinkedList is empty OR Step Not Exist")
	}
	if step == 1 {
		return l.Start.Value, nil
	}
	if step == l.Length {
		return l.end.Value, nil
	}
	i := 1
	current := l.Start
	for current.Next != nil {
		if step == i {
			return current.Value, nil
		}
		current = current.Next
		i++
	}
	return "", errors.New("LinkedList is empty OR Step Not Exist")
}
