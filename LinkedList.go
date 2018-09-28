package main

import (
	"errors"
)

type Element struct {
	value string
	next  *Element
}

// note LinkedList starts from zero index
type LinkedList struct {
	start  *Element
	end    *Element
	length int
}

func NewList() *LinkedList {
	return &LinkedList{length: 0}
}

func (l *LinkedList) push(v string) {
	e := Element{value: v}
	l.append(&e)
}

func (l *LinkedList) append(e *Element) {
	if l.length == 0 {
		l.start = e
		l.end = l.start
	} else {
		end := l.end
		end.next = e
		l.end = e
	}
	l.length++
}

func (l *LinkedList) pop() (*Element, error) {
	if l.length == 0 {
		return nil, errors.New("LinkedList is empty")
	}

	if l.length == 1 {
		popped := l.start
		l.start = nil
		l.end = nil
		l.length--
		return popped, nil
	} else {
		counter := l.length
		pointer := l.start

		for counter != 2 {
			pointer = pointer.next
			counter--
		}

		l.end = pointer
		popped := pointer.next
		l.end.next = nil
		l.length--

		return popped, nil
	}
}

// should push a new element in front
// ex 1 -> 2 -> 3 .
// shift(8)
// 8 -> 1 -> 2 ->3
func (l *LinkedList) shift(v string) {
	if l.length == 0 {
		e := Element{value: v}
		l.append(&e)
	} else {
		start := l.start
		new_e := &Element{value: v, next: start}
		l.start = new_e
	}

	l.length++
}

// 8 -> 1 -> 2 ->3
// unshift(8)
//  1 -> 2 -> 3
func (l *LinkedList) unshift() (*Element, error) {
	if l.length == 0 {
		return nil, errors.New("LinkedList is empty")
	}
	old_start := l.start
	l.start = old_start.next
	l.length--
	return old_start, nil
}

// removes an element
// 1->2->3
// remove(2)
// 1->3
func (l *LinkedList) remove(value string) error {
	if l.length == 0 {
		return errors.New("LinkedList is empty")
	}
	if l.start.value == value {
		l.unshift()
		return nil
	}
	if l.end.value == value {
		l.pop()
		return nil
	}
	current := l.start
	last := l.start
	for current.next != nil {
		if current.value == value {
			last.next = current.next
			current.next = nil
			break
		}
		last = current
		current = current.next
	}
	l.length--
	return nil
}

// remove element by step
// 4 -> 5 ->6
// remove(2)
// 4 -> 5
func (l *LinkedList) unlink(step int) error {
	if l.length == 0 || l.length < step {
		return errors.New("LinkedList is empty OR Step Not Exist")
	}
	if step == 1 {
		l.unshift()
		return nil
	}
	if step == l.length {
		l.pop()
		return nil
	}
	current := l.start
	last := l.start
	i := 1
	for current.next != nil {
		if step == i {
			last.next = current.next
			current.next = nil
			break
		}
		last = current
		current = current.next
		i++
	}
	l.length--
	return nil
}

// 4->5->6
// seek(1)
// 5
// seek(0)
// 4
func (l *LinkedList) seek(step int) (string , error) {
	if l.length == 0 || l.length < step {
		return "", errors.New("LinkedList is empty OR Step Not Exist")
	}
	if step == 1 {
		return l.start.value, nil
	}
	if step == l.length {
		return l.end.value , nil
	}
	i := 1
	current := l.start
	for current.next != nil {
		if step == i {
			return current.value , nil
		}
		current = current.next
		i++
	}
	return "", errors.New("LinkedList is empty OR Step Not Exist")
}
