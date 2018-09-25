package main

import (
	"errors"
	// "fmt"
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
		l.length--

		return popped, nil
	}
}

// should push a new element in front
// ex 1 -> 2 -> 3 .
// shift(8)
// 8 -> 1 -> 2 ->3
func (l *LinkedList) shift(v string) {

	start := l.start
	new_e := &Element{value: v , next: start}
	l.start = new_e
	l.length++
}

// 8 -> 1 -> 2 ->3
// unshift(8)
//  1 -> 2 -> 3
func (l *LinkedList) unshift() (*Element, *Element , error) {
	if l.length == 0 {
		return nil,nil, errors.New("LinkedList is empty")
	}
	old_start := l.start
	l.start = old_start.next
	l.length--
	return old_start , l.start ,nil
}

// removes an element
// 1->2->3
// remove(2)
// 1->3
func (l *LinkedList) remove(value string) {

}

// remove element by step
// 4 -> 5 ->6
// remove step 3
// 4 -> 5
func (l *LinkedList) unlink(step int) {

}

// test main function
// func main() {
// 	v := NewList()
// 	v.push("x")
// 	v.push("y")
// 	v.push("z")
// 	// x, err := v.pop()
// 	// if err != nil {
// 	// 	fmt.Println(err)
// 	// 	return
// 	// }
// 	// fmt.Println(x.value)
// 		fmt.Println(v)

// 	s,b,_ := v.unshift()
// 	fmt.Println(s.value,b.value)
// 	fmt.Println(v)
// 	// z, _ := v.pop()
// 	// fmt.Println(z.value)
// 	// fmt.Println(v)
// 	// v.push("m")
// 	// fmt.Println(v)
// 	// m, _ := v.pop()
// 	// fmt.Println(m.value)
// 	// fmt.Println(v)

// }
