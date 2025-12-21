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
type MinHeap[T bt.Number] struct {
	data []T
}

// 1. Implement a Min-Heap with insert and extract-min operations.

func (mh MinHeap[T]) isEmpty() bool {
	return len(mh.data) == 0
}

// Returns the root index
func rootIndex() int {
	return 0
}

// Returns the root value
func (mh MinHeap[T]) rootValue() T {
	return mh.data[rootIndex()]
}

// Retunrs the left child index
func leftChildIndex(index int) int {
	return 2*index + 1
}

// Returns the left child value
func (mh MinHeap[T]) leftChildValue(index int) (T, error) {
	childIndex := leftChildIndex(index)
	if childIndex >= len(mh.data) {
		var zero T
		return zero, errors.New("index out of bounds")
	}
	return mh.data[childIndex], nil
}

// Retunrs the right child index
func rightChildIndex(index int) int {
	return 2*index + 2
}

// Returns the right child value
func (mh MinHeap[T]) rightChildValue(index int) (T, error) {
	childIndex := rightChildIndex(index)
	if childIndex >= len(mh.data) {
		var zero T
		return zero, errors.New("index out of bounds")
	}
	return mh.data[childIndex], nil
}

// Returns the parent index
func parentIndex(index int) int {
	return (index - 1) / 2
}

// Returns the parent value
func (mh MinHeap[T]) parentValue(index int) (T, error) {
	parentInd := parentIndex(index)
	if parentInd < 0 {
		var zero T
		return zero, errors.New("index out of bounds")
	}
	return mh.data[parentInd], nil
}

// Helper function to keep the min heap structure
func (mh *MinHeap[T]) heapifyMin() {
	lastIndex := len(mh.data) - 1
	lastValue := mh.data[lastIndex]
	rootIndex := rootIndex()
	for lastIndex != rootIndex {
		parentValue, err := mh.parentValue(lastIndex)
		if err != nil {
			return
		}
		if parentValue <= lastValue {
			break
		}
		mh.data[lastIndex], mh.data[parentIndex(lastIndex)] = mh.data[parentIndex(lastIndex)], mh.data[lastIndex]
		lastIndex, lastValue = parentIndex(lastIndex), parentValue
	}
}

func (mh *MinHeap[T]) insert(value T) {
	// If is empty insert the value at root
	if mh.isEmpty() {
		mh.data = append(mh.data, value)
		return
	}
	// Add the new vale at last node and heapify
	mh.data = append(mh.data, value)
	mh.heapifyMin()
}

func (mh *MinHeap[T]) extractMin() T {
	var result T
	if mh.isEmpty() {
		return result
	}
	return mh.rootValue()
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
