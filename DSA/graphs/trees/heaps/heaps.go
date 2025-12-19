package heaps

import (
	bt "Practica_GO/DSA/graphs/trees/binaryTrees"
)

// 1. Implement a Min-Heap with insert and extract-min operations.
type MinHeap[T bt.Number] bt.BinaryTree[T]

func (mh *MinHeap[T]) insert(value T) {
	if mh.Root == nil {
		mh.Root = &bt.Node[T]{Value: value}
	}

}

func (mh *MinHeap[T]) extractMin() T {
	var result T
	if mh.Root == nil {
		return result
	}
	return mh.Root.Value
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
