package hashMaps

import (
	"errors"
)

// Comprehensive Hash Map Exercises

// 1. Implement a Multimap (map with multiple values per key).
//    - Write a generic multimap structure where each key maps to a slice of values.

// Generic multimap structure where each key maps to a slice of values
type Multimap[K comparable, A any] map[K][]A

//    - Support insert, remove, and get operations.

func (m Multimap[K, A]) Insert(key K, values []A) {
	m[key] = append(m[key], values...)
}

func (m Multimap[K, A]) Remove(key K) error {
	if _, ok := m[key]; !ok {
		return errors.New("not found key")
	}
	delete(m, key)
	return nil
}

func (m Multimap[K, A]) Get(key K) ([]A, error) {
	if _, ok := m[key]; !ok {
		var zero []A
		return zero, errors.New("not found key")
	}
	return m[key], nil
}

// 2. Design a Word Frequency Counter with Top-K Query.

// Struct for efficient top k query with words

// Define the node of the heap to implement an efficient algorithm
type WordFreq struct {
	Word string
	Freq int
}

// Max heap frecuency based structure to efficiently top k query
type TopKQuery struct {
	data        []WordFreq
	indexMap    map[string]int // word -> index in the heap
	currentFreq map[string]int // word -> frecuency
}

// Constructs a new TopKQuery
func NewTopKQuery() *TopKQuery {
	return &TopKQuery{
		data:        []WordFreq{},
		indexMap:    make(map[string]int),
		currentFreq: make(map[string]int),
	}
}

func (t TopKQuery) isEmpty() bool {
	return len(t.data) == 0
}

func (t TopKQuery) rootValue() WordFreq {
	return t.data[0]
}

func (t TopKQuery) leftChildIndex(index int) int {
	if index*2+1 >= len(t.data) {
		return -1
	}
	return index*2 + 1
}

func (t TopKQuery) rightChildIndex(index int) int {
	if index*2+2 >= len(t.data) {
		return -1
	}
	return index*2 + 2
}

func (t TopKQuery) parentIndex(index int) int {
	if (index-1)/2 < 0 {
		return -1
	}
	return (index - 1) / 2
}

func (t *TopKQuery) swapHeapNodes(i, j int) {
	t.data[i], t.data[j] = t.data[j], t.data[i]
	t.indexMap[t.data[i].Word] = i
	t.indexMap[t.data[j].Word] = j
}

func (t *TopKQuery) heapifyUp(index int) {
	for index > 0 {
		parentIdx := t.parentIndex(index)
		if parentIdx == -1 || t.data[index].Freq <= t.data[parentIdx].Freq {
			break
		}
		t.swapHeapNodes(index, parentIdx)
		index = parentIdx
	}
}

func (t *TopKQuery) heapifyDown(index int) {
	for {
		leftIdx := t.leftChildIndex(index)
		rightIdx := t.rightChildIndex(index)
		largest := index

		if leftIdx != -1 && t.data[leftIdx].Freq > t.data[largest].Freq {
			largest = leftIdx
		}
		if rightIdx != -1 && t.data[rightIdx].Freq > t.data[largest].Freq {
			largest = rightIdx
		}
		if largest == index {
			break
		}
		t.swapHeapNodes(index, largest)
		index = largest
	}
}

func (t *TopKQuery) RemoveRoot() {
	lastIndex := len(t.data) - 1
	rootIndex := 0
	t.swapHeapNodes(rootIndex, lastIndex)
	t.data = t.data[:lastIndex]
	t.heapifyDown(rootIndex)
}

//    - Given a stream of words, efficiently support:
//      * Adding a word (increment its count)

func (t *TopKQuery) addWord(word string) {
	t.currentFreq[word]++
	// If is an existing word in the slice add the word update the old frecuency and heapify up and down
	if index, exists := t.indexMap[word]; exists {
		t.data[index].Freq = t.currentFreq[word]
		t.heapifyUp(index)
		t.heapifyDown(index)
		return
	}
	// If instead is a new word just add it at the end and heapify up
	t.data = append(t.data, WordFreq{Word: word, Freq: t.currentFreq[word]})
	index := len(t.data) - 1
	t.indexMap[word] = index
	t.heapifyUp(index)
}

//      * Querying the top K most frequent words at any time

func (t TopKQuery) mostKFrecuentWords(k int) []string {
	// Deep copy the data slice
	dataCopy := make([]WordFreq, len(t.data))
	copy(dataCopy, t.data)
	// Deep copy the indexMap
	indexMapCopy := make(map[string]int, len(t.indexMap))
	for k, v := range t.indexMap {
		indexMapCopy[k] = v
	}
	tCopy := TopKQuery{
		data:        dataCopy,
		indexMap:    indexMapCopy,
		currentFreq: t.currentFreq, // frequencies are not mutated
	}
	var result []string
	for i := 0; i < k; i++ {
		result = append(result, tCopy.rootValue().Word)
		tCopy.RemoveRoot()
	}
	return result
}

// 3. Implement a Trie (Prefix Tree) Using Hash Maps.
//    - Build a trie for string storage and prefix search, using hash maps for children at each node.

// 4. Find All Subarrays With Equal Number of 0s and 1s.
//    - Given a binary array, use a hash map to find the number of subarrays with equal numbers of 0s and 1s.

// 5. Group Shifted Strings.
//    - Given a list of strings, group all strings that belong to the same shifting sequence (e.g., "abc", "bcd", "xyz" are in the same group).

// 6. Implement a Least Frequently Used (LFU) Cache.
//    - Design a cache that evicts the least frequently used item.
//    - Use hash maps for O(1) access and frequency tracking.

// 7. Find the Longest Substring with At Most K Distinct Characters.
//    - Given a string and an integer K, find the length of the longest substring with at most K distinct characters using a hash map.

// 8. Detect and Return All Duplicate Subtrees in a Binary Tree.
//    - Serialize each subtree using a hash map to detect duplicates.

// 9. Implement a Two Sum Data Structure.
//    - Support add(number) and find(value) to check if any two numbers sum to a value, using a hash map.

// 10. Find All Pairs of Anagrams in a List.
//     - Given a list of strings, return all pairs of indices where the strings are anagrams of each other.
