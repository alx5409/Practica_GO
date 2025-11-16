package hashMaps

import (
	"errors"
	"fmt"
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

// 13. Find the first non-repeating element in a slice using a hash map.

// 14. Merge two hash maps into one.

// 15. Find the intersection of two hash maps (keys present in both).

// 16. Find the union of two hash maps (all unique keys).

// 17. Reverse the key-value pairs in a hash map (values become keys).

// 18. Find the key with the maximum value in the hash map.

// 19. Find the key with the minimum value in the hash map.

// 20. Remove all entries with a specific value from the hash map.

// 21. Group strings by their length using a hash map.

// 22. Implement a simple cache using a hash map.

// 23. Check if two slices are anagrams using a hash map.

// 24. Find duplicate values in a hash map.

// 25. Implement a hash map with custom struct keys.

// 26. Implement a hash map to count word occurrences in a string.

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
