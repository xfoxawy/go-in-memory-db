package queue

import "github.com/alaaelgndy/go-in-memory-db/linkedlist"

type Queue struct {
	Queue *linkedlist.LinkedList
}

func NewQueue() *Queue {
	return &Queue{linkedlist.NewList()}
}

/**
* push element to queue
 */
func (q *Queue) Enqueue(e string) {
	q.Queue.Push(e)
	return
}

/**
* get first element enterd queue
 */
func (q *Queue) Dequeue() string {
	front := q.Front()
	if front == "" {
		return ""
	}
	q.Queue.Unshift()
	return front
}

// /**
// * return size of queue
//  */
func (q *Queue) Size() int {
	return q.Queue.Length
}

// /**
// * check LinkedList length
// * return bool
//  */
func (q *Queue) isEmpty() bool {
	return q.Size() == 0
}

// /**
// * the same value like dequeue without removeing the front element
//  */
func (q *Queue) Front() string {
	if !q.isEmpty() {
		return q.Queue.Start.Value
	}
	return ""
}
