package main

type Stack struct {
	stack *LinkedList
}

func NewStack() *Stack {
	return &Stack{NewList()}
}

/**
 * push element to Stack
 */
func (q *Stack) push(e string) {
	q.stack.push(e)
	return
}

/**
 * get first element enterd Stack
 */
func (q *Stack) pop() string {
	top := q.top()
	if top == "" {
		return ""
	}
	q.stack.pop()
	return top
}

// /**
// * return size of Stack
//  */
func (q *Stack) size() int {
	return q.stack.length
}

// /**
// * check LinkedList length
// * return bool
//  */
func (q *Stack) isEmpty() bool {
	return q.size() == 0
}

// /**
// * the same value like deStack without removeing the front element
//  */
func (q *Stack) top() string {
	if !q.isEmpty() {
		return q.stack.end.value
	}
	return ""
}
