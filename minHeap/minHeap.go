package heap

import (
	"errors"
	"fmt"
)

type heap struct {
	store    []int
	capacity int
	size     int
}

//creates a heap from an int array
func buildMinHeap(values []int) (*heap, error) {
	if len(values) == 0 {
		return nil, errors.New("Array is empty")
	}
	minHeap := heap{capacity: (len(values) * 2), size: 0}
	h := &minHeap
	for _, value := range values {
		h.insert(value)
		fmt.Println(h.store)
	}
	return h, nil
}

// checks if the array is not full
func (h *heap) checkCapacity() {
	if h.capacity <= h.size {
		h.capacity = h.capacity * 2
		newStore := make([]int, h.capacity)
		for _, value := range h.store {
			newStore = append(newStore, value)
		}
		h.store = newStore
	}
}

// return the minimum value and pop it off the heap
func (h *heap) extractMin() (int, error) {
	if h.size == 0 {
		return 0, errors.New("Heap is empty")
	}
	min := h.store[0]
	h.store[0] = h.store[h.size-1]
	h.size--
	h.heapifyDown(0)
	return min, nil
}

// insert key to the heap
func (h *heap) insert(key int) {
	h.checkCapacity()
	h.size++
	h.store = append(h.store, key)
	h.heapifyUp(h.size - 1)
}

func (h *heap) heapifyUp(i int) {
	iParent, root := h.parent(i)
	if root {
		return
	}
	if h.store[iParent] > h.store[i] {
		h.store[i], h.store[iParent] = h.store[iParent], h.store[i]
		h.heapifyUp(iParent)
	}
}

func (h *heap) parent(i int) (int, bool) {
	if i == 0 {
		return 0, true
	}
	return (i - 1) / 2, false
}

func (h *heap) heapifyDown(i int) {
	min := i
	minChild := (2 * i) + 1 // set minChild to the left child
	if minChild < h.size {  // test for left child
		if minChild+1 < h.size { // test for right child
			if h.store[minChild] > h.store[minChild+1] { //if right child is smaller than the left child
				minChild++
			}
		}
		if h.store[i] > h.store[minChild] { //if one of the children is smaller than parent
			min = minChild
		}
	}
	if min != i {
		h.store[i], h.store[min] = h.store[min], h.store[i]
		h.heapifyDown(min)
	}
}
