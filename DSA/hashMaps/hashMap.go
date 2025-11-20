package hashMaps

import (
	"errors"
	"fmt"
	"math/rand"
	"sort"
	"strings"
)

// 1. Create a hash map to store integer keys and values.
var intmapintvalue map[int]int

// 2. Check if a key exists in the hash map.
func isKeyInMap(hmap map[string]int, key string) bool {
	_, ok := hmap[key]
	return ok
}

// 3. Insert a key-value pair into the hash map.
func instertKeyValueInMap(hmap map[string]int, key string, value int) {
	hmap[key] = value
}

// 4. Delete a key-value pair from the hash map.
func deleteKeyValue(hmap map[string]int, key string, value int) error {
	if !isKeyInMap(hmap, key) {
		return errors.New("The key is not in the hash map")
	}
	delete(hmap, key)
	return nil
}

// 5. Retrieve the value for a given key.(
func valueWithKey(hmap map[string]int, key string) int {
	return hmap[key]
}

// 6. Update the value for an existing key.
func updateValueWithKey(hmap map[string]int, key string, newValue int) error {
	if !isKeyInMap(hmap, key) {
		return errors.New("Key not found")
	}
	hmap[key] = newValue
	return nil
}

// 7. Find the number of key-value pairs in the hash map.
func size(hmap map[string]int) int {
	counter := 0
	for key, _ := range hmap {
		counter++
	}
	return counter
}

// 8. Iterate over all key-value pairs in the hash map.
func printKeyValue(hmap map[string]int) {
	for key, value := range hmap {
		fmt.Println("Key: ", key, " Value: ", value)
	}
}

// 9. Clear all entries in the hash map.
func clearEntries(hmap map[string]int) {
	for k := range hmap {
		delete(hmap, k)
	}
}

// 10. Find all keys in the hash map.
func findKeys(hmap map[string]int) {
	for key := range hmap {
		fmt.Println("Key = ", key)
	}
}

// 11. Find all values in the hash map.
func findValue(hmap map[string]int) {
	for _, value := range hmap {
		fmt.Println("Value = ", value)
	}
}

// 12. Count the frequency of each element in a slice using a hash map.
func countFrecuency[T comparable](slice []T) map[T]int {
	hashMap := make(map[T]int)
	for _, element := range slice {
		hashMap[element]++
	}
	return hashMap
}

// 13. Find the first non-repeating element in a slice using a hash map.
func firstNonRepeatingElement[T comparable](slice []T) (T, error) {
	frecuency := make(map[T]int)
	for _, element := range slice {
		frecuency[element]++
	}
	for _, element := range slice {
		if frecuency[element] == 1 {
			return element, nil
		}
	}
	err := errors.New("Element not found")
	var zero T
	return zero, err
}

// 14. Merge two hash maps into one. If a collision happens, second map overrides the first one
func mergeHashMaps[T comparable, U any](hmap1 map[T]U, hmap2 map[T]U) map[T]U {
	mergedMap := hmap1
	for key, value := range hmap2 {
		mergedMap[key] = value
	}
	return mergedMap
}

// 15. Find the intersection of two hash maps (keys present in both).
type Pair[U any] struct {
	Value1 U
	Value2 U
}

func intersectionOfMaps[T comparable, U any](hmap1 map[T]U, hmap2 map[T]U) map[T]Pair[U] {
	intersection := make(map[T]Pair[U])
	for key, value1 := range hmap1 {
		if value2, ok := hmap2[key]; ok {
			intersection[key] = Pair[U]{value1, value2}
		}
	}
	return intersection
}

// 16. Find the union of two hash maps (keys that are unique to each map).
func unionOfMaps[T comparable, U any](hmap1 map[T]U, hmap2 map[T]U) map[T]U {
	union := make(map[T]U)
	for key, value1 := range hmap1 {
		if _, ok := hmap2[key]; !ok {
			union[key] = value1
		}
	}
	for key, value2 := range hmap2 {
		if _, ok := hmap1[key]; !ok {
			union[key] = value2
		}
	}
	return union
}

// 17. Reverse the key-value pairs in a hash map (values become keys).
func reverseMap[T, U comparable](hmap map[T]U) map[U]T {
	reversedMap := make(map[U]T)
	for key, value := range hmap {
		reversedMap[value] = key
	}
	return reversedMap
}

// 18. Find the key with the maximum value in the hash map.
func maxValueInMap[T comparable](hmap map[T]float64) T {
	var keyMax T
	first := true
	var max float64
	for key, value := range hmap {
		if first {
			max = value
			keyMax = key
			first = false
		} else if value > max {
			max = value
			keyMax = key
		}
	}
	return keyMax
}

// 19. Find the key with the minimum value in the hash map.
func minValueInMap[T comparable](hmap map[T]float64) T {
	var keyMin T
	first := true
	var min float64
	for key, value := range hmap {
		if first {
			min = value
			keyMin = key
			first = false
		} else if value < min {
			min = value
			keyMin = key
		}
	}
	return keyMin
}

// 20. Remove all entries with a specific value from the hash map.
func removeEntriesWithValue[T, U comparable](hmap map[T]U, value U) {
	for key, val := range hmap {
		if val == value {
			delete(hmap, key)
		}
	}
}

// 21. Group strings by their length using a hash map.
func groupByLength(stringSlice []string) map[int][]string {
	groups := make(map[int][]string)
	for _, element := range stringSlice {
		groups[len(element)] = append(groups[len(element)], element)
	}
	return groups
}

// 22. Implement a simple cache using a hash map.
type SimpleCache[K comparable, V any] struct {
	data map[K]V
}

func NewSimpleCache[K comparable, V any]() *SimpleCache[K, V] {
	return &SimpleCache[K, V]{data: make(map[K]V)}
}

func (c SimpleCache[K, V]) Get(key K) (V, bool) {
	val, ok := c.data[key]
	return val, ok
}

func (c *SimpleCache[K, V]) Set(key K, value V) {
	c.data[key] = value
}

func (c *SimpleCache[K, V]) Delete(key K) error {
	_, ok := c.data[key]
	if !ok {
		return errors.New("Key not found")
	}
	delete(c.data, key)
	return nil
}

func (c *SimpleCache[K, V]) Clear() {
	for key := range c.data {
		delete(c.data, key)
	}
}

func (c SimpleCache[K, V]) Size() int {
	return len(c.data)
}

// 23. Check if two slices are anagrams using a hash map.
func areAnagrams[T comparable](slice1 []T, slice2 []T) bool {
	frecuency1 := make(map[T]int)
	frecuency2 := make(map[T]int)
	for _, val := range slice1 {
		frecuency1[val]++
	}
	for _, val := range slice2 {
		frecuency2[val]++
	}
	for key, value := range frecuency1 {
		if frecuency2[key] != value {
			return false
		}
	}
	for key, value := range frecuency2 {
		if frecuency1[key] != value {
			return false
		}
	}
	return true
}

// 24. Find duplicate values in a hash map.
func duplicateValuesInMap[T, U comparable](hmap map[T]U) []U {
	var result []U
	count := make(map[U]int)
	for _, value := range hmap {
		count[value]++
	}
	for value, c := range count {
		if c > 1 {
			result = append(result, value)
		}
	}
	return result
}

// 25. Implement a hash map with custom struct keys.
type CustomStruct struct {
	name string
	age  int
}

var personMap map[CustomStruct]int

// 26. Implement a hash map to count word occurrences in a string.
func countWordOcurrences(phrase string, word string) int {
	words := strings.Split(phrase, " ")
	count := make(map[string]int)
	for _, w := range words {
		count[w]++
	}
	return count[word]
}

// 27. Find all keys with a specific value.
func findAllKeysByValues[K, V comparable](hmap map[K]V, value V) []K {
	var keys []K
	for key, val := range hmap {
		if val == value {
			keys = append(keys, key)
		}
	}
	return keys
}

// 28. Swap keys and values in a hash map (assume values are unique).
func swapKeysToValues[K, V comparable](hmap map[K]V) map[V]K {
	swappedMap := make(map[V]K)
	for key, value := range hmap {
		swappedMap[value] = key
	}
	return swappedMap
}

// 29. Implement a hash map to store student grades (studentName to grade).
type studentName string
type grade float64

func storeStudentGrades() map[studentName]grade {
	studentGrades := make(map[studentName]grade)
	exitCommand := "exit"
	for {
		fmt.Println("Write exit and press enter to leave")
		fmt.Println("Give me the name of the student: ")
		var name studentName
		fmt.Scanln(&name)
		if string(name) == exitCommand {
			break
		}
		fmt.Println("Give me the grade of the student: ", name)
		var grade grade
		fmt.Scanln(&grade)
		studentGrades[name] = grade
	}
	return studentGrades
}

// 30. Find the most frequent value in a hash map.
func mostFrecuentValueInMap[K, V comparable](hmap map[K]V) V {
	counter := make(map[V]int)
	for _, value := range hmap {
		counter[value]++
	}
	max := 0
	var mostFrecuentValue V
	for key, value := range counter {
		if max < value {
			max = value
			mostFrecuentValue = key
		}
	}
	return mostFrecuentValue
}

// 31. Implement an LRU (Least Recently Used) cache using a hash map and a doubly linked list.
type DoubleLinkedNode[V comparable] struct {
	value V
	prev  *DoubleLinkedNode[V]
	next  *DoubleLinkedNode[V]
}

type DoubleLinkedList[V comparable] struct {
	head *DoubleLinkedNode[V]
	tail *DoubleLinkedNode[V]
}

type LRU[K, V comparable] struct {
	data     map[K]*DoubleLinkedNode[V]
	history  *DoubleLinkedList[V]
	capacity int
}

// Move a node to the front of the doubly linked list
func moveNodeToFront[V comparable](node *DoubleLinkedNode[V], list *DoubleLinkedList[V]) bool {
	if node == nil || list.head == node {
		return true
	}
	// Remove node from its current position
	if node.prev != nil {
		node.prev.next = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	}
	// If node was tail, update tail
	if list.tail == node {
		list.tail = node.prev
	}
	// Insert node at the front
	node.prev = nil
	node.next = list.head
	if list.head != nil {
		list.head.prev = node
	}
	list.head = node
	// If list was empty or had one node, update tail
	if list.tail == nil {
		list.tail = node
	}
	return true
}

func (l *DoubleLinkedList[V]) insertNodeAtFront(val V) {
	newNode := &DoubleLinkedNode[V]{value: val}
	newNode.next = l.head
	if l.head != nil {
		l.head.prev = newNode
	}
	l.head = newNode
	if l.tail == nil {
		l.tail = newNode
	}
}

func (l *DoubleLinkedList[V]) insertNodeAtBack(val V) {
	newNode := &DoubleLinkedNode[V]{value: val}
	newNode.prev = l.tail
	newNode.next = nil
	if l.tail != nil {
		l.tail.next = newNode
	}
	l.tail = newNode
	if l.tail != nil {
		l.tail = newNode
	}
}

func (l *DoubleLinkedList[V]) removeNode(node *DoubleLinkedNode[V]) error {
	if node == nil {
		return errors.New("Node is nil")
	}
	// Update previous node's next pointer
	if node.prev != nil {
		node.prev.next = node.next
	} else {
		// Node is head
		l.head = node.next
	}
	// Update next node's prev pointer
	if node.next != nil {
		node.next.prev = node.prev
	} else {
		// Node is tail
		l.tail = node.prev
	}
	node.prev = nil
	node.next = nil
	return nil
}

// Retrieve a value by key and mark it as recently used
func (l *LRU[K, V]) Get(key K) (V, bool) {
	node, ok := l.data[key]
	if !ok {
		var zero V
		return zero, false
	}
	moveNodeToFront(node, l.history)
	return node.value, true
}

// Insert or update a value and mark it as recently used
func (l *LRU[K, V]) Set(key K, value V) error {
	// If the maximum cache capacity has been reached delete the tail and put the new item there
	if len(l.data) >= l.capacity {
		l.history.insertNodeAtBack(value)
		l.data[key] = l.history.tail
	}
	return nil
}

// Remove a key-value pair form the cache
func (l *LRU[K, V]) Delete(key K) error {
	node, ok := l.data[key]
	if !ok {
		return errors.New("Key not found")
	}
	l.history.removeNode(node)
	delete(l.data, key)
	return nil
}

// Clear the cache
func (l *LRU[K, V]) Clear() {
	// Clear the data
	for data := range l.data {
		delete(l.data, data)
	}
	// Clear the history
	l.history.head = nil
	l.history.tail = nil
}

func (l *LRU[K, V]) Size() int {
	return len(l.data)
}

// 32. Given a slice of integers, find all pairs that sum to a target value using a hash map.

// Function that returns the first 2 numbers that sums a given cuantity
func twoSums(slice []int, sum int) (Pair[int], error) {
	hmap := make(map[int]int)
	for _, number := range slice {
		complement := sum - number
		if val, ok := hmap[complement]; ok {
			return Pair[int]{val, number}, nil
		}
		hmap[number] = number
	}
	return Pair[int]{}, errors.New("No pairs found")
}

// Function that solver the exercise 32
func allPairsWithSum(slice []int, sum int) ([]Pair[int], error) {
	var pairsWithSum []Pair[int]
	hmap := make(map[int]int)
	for _, number := range slice {
		complement := sum - number
		if val, ok := hmap[complement]; ok {
			pairsWithSum = append(pairsWithSum, Pair[int]{number, val})
		}
		hmap[number] = number
	}
	if len(pairsWithSum) == 0 {
		return pairsWithSum, errors.New("No pairs were found")
	}
	return pairsWithSum, nil
}

// 33. Find the longest consecutive sequence (eg: 1 , 2, 3) in a slice using a hash map.
func naiveConsecutiveSequence(slice []int) []int {
	maxLength := 0
	maxIndex := 0
	for i, _ := range slice {
		counter := 0
		for j := i; j < len(slice)-1; j++ {
			if slice[j+1] != 1+slice[j] {
				break
			}
			counter++
		}
		if counter > maxLength {
			maxLength = counter
			maxIndex = i
		}
	}
	// If the slice is not empty and has max consecutive secuence of 0
	if maxLength == 0 && len(slice) > 0 {
		return []int{slice[0]}
	}
	return slice[maxIndex : maxIndex+maxLength+1]
}

func maxValueInIntMap[T comparable](hmap map[T]int) T {
	var keyMax T
	first := true
	var max int
	for key, value := range hmap {
		if first {
			max = value
			keyMax = key
			first = false
		} else if value > max {
			max = value
			keyMax = key
		}
	}
	return keyMax
}

func longestConsecutiveSequence(slice []int) []int {
	if len(slice) == 0 {
		return []int{}
	}
	if len(slice) == 1 {
		return slice
	}
	numpMap := make(map[int]int)
	for position, _ := range slice {
		for i := position; i < len(slice)-1; i++ {
			if slice[i+1] != slice[i]+1 {
				break
			}
			numpMap[position]++
		}
	}
	startPosition := maxValueInIntMap(numpMap)
	maxLength := numpMap[startPosition]
	return slice[startPosition : startPosition+maxLength+1]
}

// 34. Given two slices, find the elements that appear in both more than once using hash maps.
func findMostCommon[T comparable](slice1 []T, slice2 []T) []T {
	hmap1 := make(map[T]int)
	hmap2 := make(map[T]int)
	var result []T
	for _, key := range slice1 {
		hmap1[key]++
	}
	for _, key := range slice2 {
		hmap2[key]++
	}
	for key := range hmap1 {
		if hmap1[key] > 1 && hmap2[key] > 1 {
			result = append(result, key)
		}
	}
	return result
}

// 35. Implement a hash map that supports constant time get, set, and delete, and can return a random key in constant time.
type FastMap[K comparable, V any] struct {
	data     map[K]V
	keys     []K
	keyIndex map[K]int
}

func (m FastMap[K, V]) Get(key K) V {
	return m.data[key]
}

func (m *FastMap[K, V]) Set(key K, value V) {
	if _, exists := m.data[key]; exists {
		m.data[key] = value
		return
	}
	m.data[key] = value
	m.keys = append(m.keys, key)
	m.keyIndex[key] = len(m.keys) - 1
}

func (m *FastMap[K, V]) Delete(key K) {
	index, exists := m.keyIndex[key]
	if !exists {
		return
	}
	lastIndex := len(m.keys) - 1
	lastKey := m.keys[lastIndex]
	// Swap the key to delete with the last key
	m.keys[index] = lastKey
	m.keyIndex[lastKey] = index
	// Remove the last key
	m.keys = m.keys[:lastIndex]
	delete(m.keyIndex, key)
	delete(m.data, key)
}

func (m FastMap[K, V]) GetRandomKey() K {
	randomIndex := rand.Intn(len(m.keys))
	return m.keys[randomIndex]
}

// 36. Given a slice of strings, group all anagrams together using a hash map.
func groupAnagrams(slice []string) map[string][]string {
	anagrams := make(map[string][]string)
	for _, word := range slice {
		// Sort the word's letters to use as a key
		letters := []rune(word)
		sort.Slice(letters, func(i, j int) bool { return letters[i] < letters[j] })
		key := string(letters)
		anagrams[key] = append(anagrams[key], word)
	}
	return anagrams
}

// 37. Implement a hash map to track the frequency of rolling window elements in a slice.
func frequencyRollingWindow[T comparable](slice []T, windowSize int) []map[T]int {
	var frequencies []map[T]int
	for i := 0; i < len(slice)-windowSize; i++ {
		frequency := make(map[T]int)
		for j := 0; j < windowSize; j++ {
			frequency[slice[j]]++
		}
		frequencies = append(frequencies, frequency)
	}
	return frequencies
}

// 38. Given a slice of integers, find the length of the smallest subarray with the same degree as the original slice using a hash map.
func maxFrecuency[T comparable](slice []T) int {
	hmap := countFrecuency(slice)
	return hmap[maxValueInIntMap(hmap)]

}

func findLengthSmallestSubarrayWithSameDegree[T comparable](slice []T) int {
	degree := maxFrecuency(slice)
	frequency := countFrecuency(slice)
	firstIndex := make(map[T]int)
	lastIndex := make(map[T]int)
	for i, num := range slice {
		if _, ok := firstIndex[num]; ok {
			firstIndex[num] = i
		}
		lastIndex[num] = i
	}
	minLength := len(slice)
	for num, count := range frequency {
		if count == degree {
			length := lastIndex[num] - firstIndex[num] + 1
			if length < minLength {
				minLength = length
			}
		}
	}
	return minLength
}

// 39. Implement a hash map to store and retrieve hierarchical data (e.g., parent-child relationships).

// 40. Given a slice of integers, find the subarray with sum zero using a hash map.

// Generic type for numbers
type Number interface {
	~int | ~int32 | ~int64 | ~float32 | ~float64
}

func sliceSum[T Number](slice []T) T {
	var sum T
	for _, value := range slice {
		sum += value
	}
	return sum
}

func findSubarrayWithZeroSum(slice []int) []int {
	sumIndex := make(map[int]int)
	sum := 0
	for i, value := range slice {
		sum += value
		if sum == 0 {
			return slice[:i+1]
		}
		if prevIndex, ok := sumIndex[sum]; ok {
			return slice[prevIndex : i+1]
		}
		sumIndex[sum] = i
	}
	return nil
}

// 41. Implement a hash map to efficiently support prefix search for strings.
func isPrefix(s string, prefix string) bool {
	if len(s) < len(prefix) {
		return false
	}
	word := []rune(s)
	// prefixRunes := []rune(prefix)
	for i, char := range prefix {
		if char != word[i] {
			return false
		}
	}
	return true
}

func findPrefixInStrings(words []string, prefix string) []string {
	prefixMap := make(map[string][]string)
	for _, word := range words {
		for i := 0; i <= len(word); i++ {
			pref := word[:i]
			prefixMap[pref] = append(prefixMap[pref], word)
		}
	}
	return prefixMap[prefix]
}

// 42. Given a slice of integers, find the number of subarrays whose sum equals a target value using a hash map.

// 43. Implement a hash map to store timestamps and efficiently query the number of events in a given time range.

// 44. Given a slice of strings, find the longest substring without repeating characters using a hash map.

// 45. Implement a hash map to support undo/redo operations for key-value changes.
