package heaps

import (
	bt "Practica_GO/DSA/graphs/trees/binaryTrees"
	"errors"
)

/*
root -> index 0
for any node at index i:

	left child -> index 2 * i + 1
	right child -> index 2 * + 2
	parent -> (i - 1) / 2
*/
type Heap[T bt.Number] struct {
	data []T
}

func (h Heap[T]) isEmpty() bool {
	return len(h.data) == 0
}

// Returns the root index
func rootIndex() int {
	return 0
}

// Returns the root value
func (h Heap[T]) rootValue() T {
	return h.data[rootIndex()]
}

// Retunrs the left child index
func leftChildIndex(index int) int {
	return 2*index + 1
}

// Returns the left child value
func (h Heap[T]) leftChildValue(index int) (T, error) {
	childIndex := leftChildIndex(index)
	if childIndex >= len(h.data) {
		var zero T
		return zero, errors.New("index out of bounds")
	}
	return h.data[childIndex], nil
}

// Retunrs the right child index
func rightChildIndex(index int) int {
	return 2*index + 2
}

// Returns the right child value
func (h Heap[T]) rightChildValue(index int) (T, error) {
	childIndex := rightChildIndex(index)
	if childIndex >= len(h.data) {
		var zero T
		return zero, errors.New("index out of bounds")
	}
	return h.data[childIndex], nil
}

// Returns the parent index
func parentIndex(index int) int {
	return (index - 1) / 2
}

// Returns the parent value
func (h Heap[T]) parentValue(index int) (T, error) {
	parentInd := parentIndex(index)
	if parentInd < 0 {
		var zero T
		return zero, errors.New("index out of bounds")
	}
	return h.data[parentInd], nil
}

// 1. Implement a Min-Heap with insert and extract-min operations.
type MinHeap[T bt.Number] struct {
	Heap[T]
}

// Helper function to keep the min heap structure
func (h *Heap[T]) heapifyMin() {
	lastIndex := len(h.data) - 1
	lastValue := h.data[lastIndex]
	rootIndex := rootIndex()
	for lastIndex != rootIndex {
		parentValue, err := h.parentValue(lastIndex)
		if err != nil {
			return
		}
		if parentValue <= lastValue {
			break
		}
		h.data[lastIndex], h.data[parentIndex(lastIndex)] = h.data[parentIndex(lastIndex)], h.data[lastIndex]
		lastIndex, lastValue = parentIndex(lastIndex), parentValue
	}
}

func (hm *MinHeap[T]) insert(value T) {
	// If is empty insert the value at root
	if hm.isEmpty() {
		hm.data = append(hm.data, value)
		return
	}
	// Add the new vale at last node and heapify
	hm.data = append(hm.data, value)
	hm.heapifyMin()
}

func (hm *MinHeap[T]) extractMin() T {
	var result T
	if hm.isEmpty() {
		return result
	}
	return hm.rootValue()
}

// 2. Implement a Max-Heap with insert and extract-max operations.

type MaxHeap[T bt.Number] struct {
	Heap[T]
}

// Helper function to keep the max heap structure
func (h *MaxHeap[T]) heapifyMax() {
	lastIndex := len(h.data) - 1
	lastValue := h.data[lastIndex]
	rootIndex := rootIndex()
	for lastIndex != rootIndex {
		parentValue, err := h.parentValue(lastIndex)
		if err != nil {
			return
		}
		if parentValue >= lastValue {
			break
		}
		h.data[lastIndex], h.data[parentIndex(lastIndex)] = h.data[parentIndex(lastIndex)], h.data[lastIndex]
		lastIndex, lastValue = parentIndex(lastIndex), parentValue
	}
}

func (hM *MaxHeap[T]) insert(value T) {
	// If is empty insert the value at root
	if hM.isEmpty() {
		hM.data = append(hM.data, value)
		return
	}
	// Add the new vale at last node and heapify
	hM.data = append(hM.data, value)
	hM.heapifyMax()
}

func (hM *MaxHeap[T]) extractMax() T {
	if hM.isEmpty() {
		var zero T
		return zero
	}
	return hM.rootValue()
}

// 3. Build a heap from an unsorted array (heapify).
func heapifyslice[T bt.Number](slice []T) MinHeap[T] {
	var result MinHeap[T]
	if len(slice) == 0 {
		return result
	}
	for _, value := range slice {
		result.insert(value)
	}
	return result
}

// 4. Implement heap sort using a heap.

// 5. Find the kth smallest element in an array using a heap.

// 6. Find the kth largest element in an array using a heap.

// 7. Merge k sorted arrays using a min-heap.

// 8. Check if a given array represents a valid min-heap.

// 9. Check if a given array represents a valid max-heap.

// 10. Convert a min-heap to a max-heap (and vice versa).

// 11. Implement a priority queue using a heap.

// 12. Increase or decrease the key value of a given element in a heap.

// 13. Remove an arbitrary element from a heap.

// 14. Find the median of a stream of numbers using two heaps.

// 15. Implement a d-ary heap (where each node has d children) and its operations.
