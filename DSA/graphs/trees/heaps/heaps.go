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

// 3. Build a min heap from an unsorted array (heapify).
func minHeapifySlice[T bt.Number](slice []T) MinHeap[T] {
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

func (mh *MinHeap[T]) heapifyMinDown() {
	index := rootIndex()
	lastIndex := len(mh.data) - 1

	for {
		leftIdx := leftChildIndex(index)
		rightIdx := rightChildIndex(index)
		smallest := index

		if leftIdx <= lastIndex && mh.data[leftIdx] < mh.data[smallest] {
			smallest = leftIdx
		}
		if rightIdx <= lastIndex && mh.data[rightIdx] < mh.data[smallest] {
			smallest = rightIdx
		}
		if smallest == index {
			break
		}
		mh.data[index], mh.data[smallest] = mh.data[smallest], mh.data[index]
		index = smallest
	}
}

func (mh *MinHeap[T]) removeRoot() {
	lastIndex := len(mh.data) - 1
	mh.data[rootIndex()] = mh.data[lastIndex]
	mh.data = mh.data[:lastIndex]
	mh.heapifyMinDown()
}

func heapSort[T bt.Number](slice []T) []T {
	var orderedSlice []T
	if len(slice) == 0 {
		return orderedSlice
	}
	minHeap := minHeapifySlice(slice)
	for !minHeap.isEmpty() {
		orderedSlice = append(orderedSlice, minHeap.rootValue())
		minHeap.removeRoot()
	}
	return orderedSlice
}

// 5. Find the kth smallest element in an array using a heap.
func kthSmallestElement[T bt.Number](slice []T, k int) (T, error) {
	var kthSmallest T
	if k <= 0 || k > len(slice) {
		return kthSmallest, errors.New("index out of bounds")
	}
	minHeap := minHeapifySlice(slice)
	for i := 1; i < k; i++ {
		minHeap.removeRoot()
	}
	return minHeap.rootValue(), nil
}

// 6. Find the kth largest element in an array using a heap.
func maxHeapifySlice[T bt.Number](slice []T) MaxHeap[T] {
	var result MaxHeap[T]
	if len(slice) == 0 {
		return result
	}
	for _, value := range slice {
		result.insert(value)
	}
	return result
}

func (mH *MaxHeap[T]) heapifyMaxDown() {
	index := rootIndex()
	lastIndex := len(mH.data) - 1

	for {
		leftIdx := leftChildIndex(index)
		rightIdx := rightChildIndex(index)
		largest := index

		if leftIdx <= lastIndex && mH.data[leftIdx] > mH.data[largest] {
			largest = leftIdx
		}
		if rightIdx <= lastIndex && mH.data[rightIdx] > mH.data[largest] {
			largest = rightIdx
		}
		if largest == index {
			break
		}
		mH.data[index], mH.data[largest] = mH.data[largest], mH.data[index]
		index = largest
	}
}

func (mh *MaxHeap[T]) removeRoot() {
	lastIndex := len(mh.data) - 1
	mh.data[rootIndex()] = mh.data[lastIndex]
	mh.data = mh.data[:lastIndex]
	mh.heapifyMaxDown()
}

func kthLargestElement[T bt.Number](slice []T, k int) (T, error) {
	var kthSmallest T
	if k <= 0 || k > len(slice) {
		return kthSmallest, errors.New("index out of bounds")
	}
	maxHeap := maxHeapifySlice(slice)
	for i := 1; i < k; i++ {
		maxHeap.removeRoot()
	}
	return maxHeap.rootValue(), nil
}

// 7. Merge k sorted arrays using a min-heap.

type heapNode[T bt.Number] struct {
	value    T
	arrayIdx int // from which array comes from
	elemIdx  int // index within that array
}

// MinHeap for heapNode[T]
type MinHeapNode[T bt.Number] struct {
	data []heapNode[T]
}

func (h *MinHeapNode[T]) isEmpty() bool {
	return len(h.data) == 0
}

func (h *MinHeapNode[T]) insert(node heapNode[T]) {
	h.data = append(h.data, node)
	h.heapifyUp(len(h.data) - 1)
}

func (h *MinHeapNode[T]) rootValue() heapNode[T] {
	return h.data[0]
}

func (h *MinHeapNode[T]) removeRoot() {
	lastIndex := len(h.data) - 1
	h.data[0] = h.data[lastIndex]
	h.data = h.data[:lastIndex]
	h.heapifyDown(0)
}

func (h *MinHeapNode[T]) heapifyUp(index int) {
	for index > 0 {
		parent := (index - 1) / 2
		if h.data[parent].value <= h.data[index].value {
			break
		}
		h.data[parent], h.data[index] = h.data[index], h.data[parent]
		index = parent
	}
}

func (h *MinHeapNode[T]) heapifyDown(index int) {
	n := len(h.data)
	for {
		left := 2*index + 1
		right := 2*index + 2
		smallest := index

		if left < n && h.data[left].value < h.data[smallest].value {
			smallest = left
		}
		if right < n && h.data[right].value < h.data[smallest].value {
			smallest = right
		}
		if smallest == index {
			break
		}
		h.data[index], h.data[smallest] = h.data[smallest], h.data[index]
		index = smallest
	}
}

func mergeSortedArrays[T bt.Number](slices [][]T) []T {
	var minHeap MinHeapNode[T]
	var orderedSlice []T

	// Insert the first element of each slice and keep track where the min comes from
	for i, slice := range slices {
		if len(slice) > 0 {
			minHeap.insert(heapNode[T]{value: slice[0], arrayIdx: i, elemIdx: 0})
		}
	}

	// Remove the root and add the next element in the heap where the removed element came from
	for !minHeap.isEmpty() {
		orderedSlice = append(orderedSlice, minHeap.rootValue().value)
		removedArrayIdx := minHeap.data[0].arrayIdx
		removedElemIdx := minHeap.data[0].elemIdx
		minHeap.removeRoot()
		if removedElemIdx+1 >= len(slices[removedArrayIdx]) {
			continue
		}
		minHeap.insert(heapNode[T]{
			value:    slices[removedArrayIdx][removedElemIdx+1],
			arrayIdx: removedArrayIdx,
			elemIdx:  removedElemIdx + 1,
		})
	}
	return orderedSlice
}

// 8. Check if a given array represents a valid min-heap.
func isSliceValidMinHeap[T bt.Number](slice []T) bool {
	size := len(slice)
	// Check that every posible node satisfies the heap condition
	for i, parentValue := range slice {
		leftChildIdx := leftChildIndex(i)
		rightChildIdx := rightChildIndex(i)
		if rightChildIdx < size && parentValue > slice[rightChildIdx] {
			return false
		}
		if leftChildIdx < size && parentValue > slice[leftChildIdx] {
			return false
		}
	}
	return true
}

// 9. Check if a given array represents a valid max-heap.
func isSliceValidMaxHeap[T bt.Number](slice []T) bool {
	size := len(slice)
	// Check that every posible node satisfies the heap condition
	for i, parentValue := range slice {
		leftChildIdx := leftChildIndex(i)
		rightChildIdx := rightChildIndex(i)
		if rightChildIdx < size && parentValue < slice[rightChildIdx] {
			return false
		}
		if leftChildIdx < size && parentValue < slice[leftChildIdx] {
			return false
		}
	}
	return true
}

// 10. Convert a min-heap to a max-heap (and vice versa).
func minHeapToMaxHeap[T bt.Number](minHeap MinHeap[T]) MaxHeap[T] {
	maxHeap := MaxHeap[T]{Heap: Heap[T]{data: minHeap.data}}
	maxHeap.heapifyMax()
	return maxHeap
}

func maxHeapToMinHeap[T bt.Number](maxHeap MaxHeap[T]) MinHeap[T] {
	minHeap := MinHeap[T]{Heap: Heap[T]{data: maxHeap.data}}
	minHeap.heapifyMin()
	return minHeap
}

// 11. Implement a priority queue using a heap.
type PairPriority[T bt.Number] struct {
	priority int
	value    T
}

type MaxHeapPriority[T bt.Number] struct {
	data []PairPriority[T]
}

func (h *MaxHeapPriority[T]) isEmpty() bool {
	return len(h.data) == 0
}

func (h *MaxHeapPriority[T]) insert(pair PairPriority[T]) {
	h.data = append(h.data, pair)
	h.heapifyUp(len(h.data) - 1)
}

func (h *MaxHeapPriority[T]) rootValue() PairPriority[T] {
	return h.data[0]
}

func (h *MaxHeapPriority[T]) removeRoot() {
	lastIndex := len(h.data) - 1
	h.data[0] = h.data[lastIndex]
	h.data = h.data[:lastIndex]
	h.heapifyDown(0)
}

func (h *MaxHeapPriority[T]) heapifyUp(index int) {
	for index > 0 {
		parent := (index - 1) / 2
		if h.data[parent].priority >= h.data[index].priority {
			break
		}
		h.data[parent], h.data[index] = h.data[index], h.data[parent]
		index = parent
	}
}

func (h *MaxHeapPriority[T]) heapifyDown(index int) {
	n := len(h.data)
	for {
		left := 2*index + 1
		right := 2*index + 2
		largest := index

		if left < n && h.data[left].priority > h.data[largest].priority {
			largest = left
		}
		if right < n && h.data[right].priority > h.data[largest].priority {
			largest = right
		}
		if largest == index {
			break
		}
		h.data[index], h.data[largest] = h.data[largest], h.data[index]
		index = largest
	}
}

type MaxPriorityQueue[T bt.Number] struct {
	MaxHeapPriority[T]
}

func (q *MaxPriorityQueue[T]) enqueue(value T, priority int) {
	q.insert(PairPriority[T]{value: value, priority: priority})
}

func (q *MaxPriorityQueue[T]) dequeue() error {
	if q.isEmpty() {
		return errors.New("empty queue")
	}
	q.removeRoot()
	return nil
}

func (q MaxHeapPriority[T]) peek() (T, error) {
	if q.isEmpty() {
		var zero T
		return zero, errors.New("empty queue")
	}
	return q.rootValue().value, nil
}

// 12. Increase or decrease the key value of a given element in a heap.
func (h *MinHeap[T]) bubbleUp(index int) {
	for index > 0 {
		parentIdx := parentIndex(index)
		if h.data[parentIdx] <= h.data[index] {
			break
		}
		h.data[parentIdx], h.data[index] = h.data[index], h.data[parentIdx]
		index = parentIdx
	}
}

func (h *MinHeap[T]) bubbleDown(index int) {
	size := len(h.data)
	for {
		left := leftChildIndex(index)
		right := rightChildIndex(index)
		smallest := index
		if left < size && h.data[left] < h.data[smallest] {
			smallest = left
		}
		if right < size && h.data[right] < h.data[smallest] {
			smallest = right
		}
		if smallest == index {
			break
		}
		h.data[index], h.data[smallest] = h.data[smallest], h.data[index]
		index = smallest
	}
}

func (h *MinHeap[T]) updateElement(value T, index int) {
	oldValue := h.data[index]
	if oldValue == value {
		return
	}
	h.data[index] = value
	// If is greater bubble up from that node
	if oldValue < value {
		h.bubbleUp(index)
	}
	// If is greater bubble down from that node
	if oldValue > value {
		h.bubbleDown(index)
	}
}

// 13. Remove an arbitrary element from a heap.
func (h *MinHeap[T]) delete(index int) error {
	if h.isEmpty() {
		return errors.New("empty heap")
	}
	if index < 0 || index >= len(h.data) {
		return errors.New("index out of bounds")
	}
	// If the element to be removed, there is no problem with the rest of the heap
	lastIndex := len(h.data) - 1
	if index == lastIndex {
		h.data = h.data[:lastIndex]
		return nil
	}
	// Replace the last value with the value at the index
	h.data[index], h.data[lastIndex] = h.data[lastIndex], h.data[index]
	h.data = h.data[:lastIndex]
	// Bubble up and down to mantain the heap structure
	h.bubbleUp(index)
	h.bubbleDown(index)
	return nil
}

// 14. Find the median of a stream of numbers using two heaps.
type MedianFinder[T bt.Number] struct {
	low  MaxHeap[T]
	high MinHeap[T]
}

func (mf *MedianFinder[T]) BalanceHeaps() {
	if len(mf.low.data) > len(mf.high.data)+1 {
		removedValue := mf.low.rootValue()
		mf.low.removeRoot()
		mf.high.insert(removedValue)
		return
	}
	if len(mf.high.data) > len(mf.low.data)+1 {
		removedValue := mf.high.rootValue()
		mf.high.removeRoot()
		mf.low.insert(removedValue)
		return
	}
}

func (mf *MedianFinder[T]) AddNumber(number T) {
	// Insert into max-heap if empty or num <= max of low
	if len(mf.low.data) == 0 || number <= mf.low.rootValue() {
		mf.low.insert(number)
	} else {
		mf.high.insert(number)
	}
	mf.BalanceHeaps()
}

func (mf *MedianFinder[T]) FindMedian() T {
	// When the max heap contains one more element than the min heap return the root value on the max heap
	if len(mf.low.data) == len(mf.high.data)+1 {
		return mf.low.rootValue()
	}
	if len(mf.high.data) == len(mf.low.data)+1 {
		return mf.high.rootValue()
	}
	// When min and max heap contains the same amount of elements return the minimum of those root values
	return min(mf.low.rootValue(), mf.high.rootValue())
}

// 15. Implement a d-ary heap (where each node has d children) and its operations.
type DaryMinHeap[T bt.Number] struct {
	data []T
	d    int // number of max children per node
}

/*
Since every node has at most d children now the index of the children are as follows:
	for node with index i:
		has children with indices d*i + 1 to d*i + d
	for node with index i:
		parent has index = (i-1)/d
*/

func (h DaryMinHeap[T]) isEmpty() bool {
	return len(h.data) == 0
}

func (h DaryMinHeap[T]) parentIndex(index int) int {
	if index == 0 {
		return -1
	}
	return (index - 1) / h.d
}

func (h DaryMinHeap[T]) childIndex(index int, k int) int {
	return h.d*index + k
}

func (h DaryMinHeap[T]) childrenIndices(index int) []int {
	var indices []int
	for i := 1; i <= h.d; i++ {
		indices = append(indices, h.childIndex(index, i))
	}
	return indices
}

func (h *DaryMinHeap[T]) rootValue() (T, error) {
	if h.isEmpty() {
		var zero T
		return zero, errors.New("empty heap")
	}
	return h.data[0], nil
}

func (h *DaryMinHeap[T]) bubbleUp(index int) {
	for index > 0 {
		parentIdx := h.parentIndex(index)
		if h.data[parentIdx] <= h.data[index] {
			break
		}
		// swap the parent and the children
		h.data[parentIdx], h.data[index] = h.data[index], h.data[parentIdx]
	}
}

func (h *DaryMinHeap[T]) bubbleDown(index int) {
	n := len(h.data)
	for {
		min := index
		for k := 1; k < h.d; k++ {
			childIdx := h.d*index + k
			if childIdx < n && h.data[childIdx] < h.data[min] {
				min = childIdx
			}
		}
		if min == index {
			break
		}
		h.data[min], h.data[index] = h.data[index], h.data[min]
	}
}

func (h *DaryMinHeap[T]) insert(value T) {
	// Add value at the end of the slice and then bubble up from there
	h.data = append(h.data, value)
	lastIndex := len(h.data) - 1
	h.bubbleUp(lastIndex)
}

func (h *DaryMinHeap[T]) removeRoot() {
	if h.isEmpty() {
		return
	}
	// Remove the root value by updating it with the last one and bubble down from the root
	rootIdx := rootIndex()
	lastIndex := len(h.data) - 1
	h.data[0] = h.data[lastIndex]
	h.data = h.data[:lastIndex]
	h.bubbleDown(rootIdx)
}

func buildDaryHeapFromSlice[T bt.Number](slice []T, d int) DaryMinHeap[T] {
	heap := DaryMinHeap[T]{data: append([]T{}, slice...), d: d}
	n := len(slice)
	if n == 0 {
		return heap
	}

	// Start from the last parent and bubble down to root
	for i := (n - 2) / d; i >= 0; i-- {
		heap.bubbleDown(i)
	}
	return heap
}
