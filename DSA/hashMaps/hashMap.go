package hashMaps

import (
	"errors"
	"fmt"
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

// 28. Swap keys and values in a hash map (assume values are unique).

// 29. Implement a hash map to store student grades (name to grade).

// 30. Find the most frequent value in a hash map.

// 31. Implement an LRU (Least Recently Used) cache using a hash map and a doubly linked list.

// 32. Given a slice of integers, find all pairs that sum to a target value using a hash map.

// 33. Find the longest consecutive sequence in a slice using a hash map.

// 34. Given two slices, find the elements that appear in both more than once using hash maps.

// 35. Implement a hash map that supports constant time get, set, and delete, and can return a random key in constant time.

// 36. Given a slice of strings, group all anagrams together using a hash map.

// 37. Implement a hash map to track the frequency of rolling window elements in a slice.

// 38. Given a slice of integers, find the length of the smallest subarray with the same degree as the original slice using a hash map.

// 39. Implement a hash map to store and retrieve hierarchical data (e.g., parent-child relationships).

// 40. Given a slice of integers, find the subarray with sum zero using a hash map.

// 41. Implement a hash map to efficiently support prefix search for strings.

// 42. Given a slice of integers, find the number of subarrays whose sum equals a target value using a hash map.

// 43. Implement a hash map to store timestamps and efficiently query the number of events in a given time range.

// 44. Given a slice of strings, find the longest substring without repeating characters using a hash map.

// 45. Implement a hash map to support undo/redo operations for key-value changes.
