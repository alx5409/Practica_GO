package hashMaps

import "errors"

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
//    - Given a stream of words, efficiently support:
//      * Adding a word (increment its count)
//      * Querying the top K most frequent words at any time

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
