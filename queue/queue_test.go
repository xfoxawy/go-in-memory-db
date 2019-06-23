package queue

import "testing"

var queue *Queue

func init() {
	queue = NewQueue()
}

func TestEnqueue(t *testing.T) {
	queue.Enqueue("string here")
	if queue.Size() == 0 {
		t.Errorf("expected size larger than one")
	}
}

func TestDequeue(t *testing.T) {
	queue = NewQueue()
	if queue.Dequeue() != "" {
		t.Error("expected", "empty string", "got", queue.Dequeue())
	}
	queue.Enqueue("string here")
	sizeBefore := queue.Size()
	queue.Dequeue()
	sizeAfter := queue.Size()
	if sizeAfter != sizeBefore-1 {
		t.Error("expected", sizeBefore-1, "got", sizeAfter)
	}
}

func TestSizet(t *testing.T) {
	queue.Enqueue("string here")
	if queue.Size() == 0 {
		t.Errorf("expected size larger than one")
	}
}

func TestFront(t *testing.T) {
	queue = NewQueue()
	if queue.Front() != "" {
		t.Error("expected", "empty string", "got", queue.Front())
	}
	value1 := "custom string here1"
	queue.Enqueue(value1)
	if queue.Front() != value1 {
		t.Error("expected", value1, "got", queue.Front())
	}
}
