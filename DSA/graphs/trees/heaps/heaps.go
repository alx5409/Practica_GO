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
	parent -> (i - 2) / 2
*/
type MinHeap[T bt.Number] struct {
	data []T
}

// 1. Implement a Min-Heap with insert and extract-min operations.

func (mh MinHeap[T]) isEmpty() bool {
	return len(mh.data) == 0
}

// Returns the root value
func (mh MinHeap[T]) rootValue() T {
	return mh.data[0]
}

// Returns the left child value
func (mh MinHeap[T]) leftChildValue(index int) (T, error) {
	leftChildIndex := 2*index + 1
	if leftChildIndex >= len(mh.data) {
		var zero T
		return zero, errors.New("index out of bounds")
	}
	return mh.data[leftChildIndex], nil
}

// Returns the right child value
func (mh MinHeap[T]) rightChildValue(index int) (T, error) {
	rightChildIndex := 2 * index
	if rightChildIndex >= len(mh.data) {
		var zero T
		return zero, errors.New("index out of bounds")
	}
	return mh.data[rightChildIndex], nil
}

// Returns the parent value
func (mh MinHeap[T]) parentValue(index int) (T, error) {
	parentIndex := (index - 2) / 2
	if parentIndex < 0 {
		var zero T
		return zero, errors.New("index out of bounds")
	}
	return mh.data[parentIndex], nil
}

// Helper function to keep the min heap structure
func(mh MinHeap[T]) heapifyMin[T bt.Number]() {
	lastIndex := len(mh.data)
	lastValue := data[lastIndex]
	if lastValue < mh.parent(lastIndex) {
		mh.data[lastIndex], mh.data[]
	}
}

func (mh *MinHeap[T]) insert(value T) {
	// If is empty insert the value at root
	if mh.isEmpty() {
		mh.data = append(mh.data, value)
	}
	// Add the new vale at last node and heapify
	mh.data = append(mh.data, value)
	heapifyMin(mh.data)
}

func (mh *MinHeap[T]) extractMin() T {
	var result T
	if mh.isEmpty() {
		return result
	}
	return mh.root()
}

// 2. Implement a Max-Heap with insert and extract-max operations.
type MaxHeap[T bt.Number] bt.BinaryTree[T]

func (Mh *MaxHeap[T]) insert(value T) {

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
