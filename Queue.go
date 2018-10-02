package main

type Queue struct {
	queue *LinkedList
}

func NewQueue() *Queue {
	return &Queue{NewList()}
}

/**
* push element to queue
 */
func (q *Queue) enqueue(e string) {
	q.queue.push(e)
	return
}

/**
* get first element enterd queue
 */
func (q *Queue) dequeue() string {
	front := q.front()
	if front == "" {
		return ""
	}
	q.queue.unshift()
	return front
}

// /**
// * return size of queue
//  */
func (q *Queue) size() int {
	return q.queue.length
}

// /**
// * check LinkedList length
// * return bool
//  */
func (q *Queue) isEmpty() bool {
	return q.size() == 0
}

// /**
// * the same value like dequeue without removeing the front element
//  */
func (q *Queue) front() string {
	if !q.isEmpty() {
		return q.queue.start.value
	}
	return ""
}
