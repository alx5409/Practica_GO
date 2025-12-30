package hashMaps

import (
	utils "Practica_GO/DSA/utils"
	"errors"
	"fmt"
	"sort"
	"strings"
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

// TODO: the complexity of the problems yields this exercise to a new project
// 3. Implement a Trie (Prefix Tree) Using Hash Maps.

//    - Build a trie for string storage and prefix search, using hash maps for children at each node.

// Data about english alphabet
const alphabetSize int = 26
const alphabet = "abcdefghijklmnopqrstuvwxyz"

var alphabetRunes = []rune("abcdefghijklmnopqrstuvwxyz")

type PrefixNode struct {
	character   rune
	children    map[rune]*PrefixNode
	isEndOfWord bool
}

type Trie struct {
	Root *PrefixNode
}

func NewTrie() *Trie {
	return &Trie{
		Root: &PrefixNode{
			character:   '0',
			isEndOfWord: false,
			children:    make(map[rune]*PrefixNode),
		},
	}
}

// Search for all possible strings after introducing a string
func (t Trie) search(key string) bool {
	node := t.Root
	for _, char := range key {
		next, ok := node.children[char]
		if !ok {
			return false
		}
		node = next
	}
	return node.isEndOfWord
}

func (t *Trie) insert(character rune) {

}

func (t *Trie) delete(character rune) {

}

// 4. Find All Subarrays With Equal Number of 0s and 1s.
//   - Given a binary array, use a hash map to find the number of subarrays with equal numbers of 0s and 1s.
func isBinarySlice(slice []int) bool {
	zero := 0
	one := 1
	for _, value := range slice {
		if value != zero && value != one {
			return false
		}
	}
	return true
}

func efficientNumberOfSubslicesWithEqual0sAnd1s(binarySlice []int) int {
	if !isBinarySlice(binarySlice) {
		return -1
	}
	counter := 0
	prefixSum := 0
	prefixSumMap := make(map[int]int)
	prefixSumMap[0] = 1 // to count subarrays starting from index 0

	for _, value := range binarySlice {
		if value == 1 {
			prefixSum += value
		}
		// consider the values 0 as -1 to see if the sublice has same number of 0s and 1s which will have zero sum
		if value == 0 {
			prefixSum -= 1
		}
		// if the same prefix sum has occured befora at some previous index, the subslice has the same number of 0s and 1s
		counter += prefixSumMap[prefixSum]
		prefixSumMap[prefixSum]++
	}
	return counter
}

// 5. Group Shifted Strings.
//   - Given a list of strings, group all strings that belong to the same shifting sequence (e.g., "abc", "bcd", "xyz" are in the same group).
func largestStringLength(stringSlice []string) int {
	max := 0
	for _, s := range stringSlice {
		if max < len(s) {
			max = len(s)
		}
	}
	return max
}

func separationMap(stringSlice []string) map[string][]int {
	hmap := make(map[string][]int)
	for _, s := range stringSlice {
		intStringRunes := []rune(s)
		firstRune := intStringRunes[0]
		for _, r := range intStringRunes {
			hmap[s] = append(hmap[s], int(firstRune)-int(r))
		}
	}
	return hmap
}

func separationKey(sMap map[string][]int) map[string]string {
	sepKeyMap := make(map[string]string)
	for key, value := range sMap {
		// transform the integer slice into a string by just concatenating
		var parts []string
		for _, num := range value {
			sNum := fmt.Sprintf("%d", num)
			parts = append(parts, sNum)
		}
		sValue := strings.Join(parts, "")
		sepKeyMap[key] = sValue
	}
	return sepKeyMap
}

func groupShiftingStrings(stringList []string) [][]string {
	sepMap := separationMap(stringList)
	sepKeys := separationKey(sepMap)
	stringGroups := make(map[string][]string)
	for _, s := range stringList {
		key := sepKeys[s]
		stringGroups[key] = append(stringGroups[key], s)
	}
	var result [][]string
	for _, group := range stringGroups {
		result = append(result, group)
	}
	return result
}

// TODO: need to see double linked list first
// 6. Implement a Least Frequently Used (LFU) Cache.
//    - Design a cache that evicts the least frequently used item.
//    - Use hash maps for O(1) access and frequency tracking.

// 7. Find the Longest Substring with At Most K Distinct Characters.
//    - Given a string and an integer K, find the length of the longest substring with at most K distinct characters using a hash map.

func longestSubstringKDistinct(s string, k int) int {
	maxLength := 0
	if k == 0 || len(s) == 0 {
		return maxLength
	}
	// hash map to track the count of each character in the current sliding window
	charMap := make(map[byte]int)
	left := 0
	// sliding window to find the substring
	for right := 0; right < len(s); right++ {
		charMap[s[right]]++
		for len(charMap) > k {
			charMap[s[left]]--
			if charMap[s[left]] == 0 {
				delete(charMap, s[left])
			}
			left++
		}
		if right-left > maxLength {
			maxLength = right - left + 1
		}
	}
	return maxLength
}

// 8. Detect and Return All Duplicate Subtrees in a Binary Tree.
//    - Serialize each subtree using a hash map to detect duplicates.

type Node[T comparable] struct {
	value T
	left  *Node[T]
	right *Node[T]
}

type BinaryTree[T comparable] struct {
	Root *Node[T]
}

// function to make a unique signature to serialize a subtree in the following way: "value","left_serialization","right_serialization"
//
// For example, for a node with value 1, left child 2, and right child 3, the serialization is:
// "1,2,#,#,3,#,#"
func serializeSubtree[T comparable](node *Node[T]) string {
	if node == nil {
		return "#"
	}
	left := serializeSubtree(node.left)
	right := serializeSubtree(node.right)
	return fmt.Sprintf("%v,%s,%s", node.value, left, right)
}

// helper function to serialize the entire tree
func duplicatesHelper[T comparable](node *Node[T], seen map[string]int, result *[]*Node[T]) string {
	if node == nil {
		return "#"
	}
	serial := serializeSubtree(node)
	seen[serial]++
	if seen[serial] == 2 {
		*result = append(*result, node)
	}
	duplicatesHelper(node.left, seen, result)
	duplicatesHelper(node.right, seen, result)
	return serial
}

// Returns duplicates by traversing the tree and marking the seen nodes
func (b BinaryTree[T]) duplicates() []*Node[T] {
	seen := make(map[string]int)
	var result []*Node[T]
	duplicatesHelper(b.Root, seen, &result)
	return result
}

// 9. Implement a Two Sum Data Structure.

// Two sum problem for context: find two numbers a and b that sums an objective sum

// Function to find two numbers in a slice that sums up to an objective
func twoSums(intSlice []int, sum int) (int, int) {
	hmap := make(map[int]bool)
	for _, value := range intSlice {
		complement := sum - value
		if hmap[complement] {
			return value, complement
		}
		hmap[value] = true
	}
	return -1, -1
}

type TwoSum[N utils.Number] struct {
	data []N
}

// - Support add(number) and find(value) to check if any two numbers sum to a value, using a hash map.
func (t *TwoSum[N]) add(number N) {
	t.data = append(t.data, number)
}

func (t *TwoSum[N]) find(value N) bool {
	hmap := make(map[N]bool)
	for _, val := range t.data {
		complement := value - val
		if hmap[complement] {
			return true
		}
		hmap[val] = true
	}
	return false
}

// 10. Find All Pairs of Anagrams in a List.
//     - Given a list of strings, return all indices where the strings are anagrams of each other.

// Sorts the string and returns that as a signature
func signature(s string) string {
	runes := []rune(s)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}

func allPaisOfAnagrams(list []string) [][]int {
	indicesGroups := make(map[string][]int)
	// for each string in the list builds the signature and appends the index according its signature
	for i, s := range list {
		sgn := signature(s)
		indicesGroups[sgn] = append(indicesGroups[sgn], i)
	}

	var result [][]int
	for _, indices := range indicesGroups {
		if len(indices) > 1 {
			result = append(result, indices)
		}
	}
	return result
}
