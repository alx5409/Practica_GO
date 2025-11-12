package dsa

import (
	"errors"
	"fmt"
)

// 1. Rotate an array to the right by k steps.
func rotateArray(array []int, steps int) ([]int, error) {
	rotatedArray := array
	if steps <= 0 {
		er := errors.New("The number of steps should be positive.")
		return rotatedArray, er
	}
	for i, _ := range array {
		rotatedArray[(i+steps)%len(array)] = array[i]
	}
	return rotatedArray, nil
}

// 2. Find the longest common prefix among an array of strings.
func longest_common_prefix(stringSlice []string) (string, error) {
	prefix := ""
	if len(stringSlice) == 0 {
		er := errors.New("There are no strings")
		return prefix, er
	}
	return prefix, nil
}

// 3. Implement a function to check if two strings are anagrams.

// 4. Find the first non-repeating character in a string.

// 5. Implement a singly linked list and write a function to reverse it.

// 6. Detect a cycle in a linked list.

// 7. Implement a stack and use it to check for balanced parentheses in a string.

// 8. Implement a queue and use it to simulate a simple task scheduler.

// 9. Find the intersection node of two singly linked lists.

// 10. Implement a function to find the majority element in an array (element that appears more than n/2 times).

// 11. Use a map to count the frequency of words in a paragraph.

// 12. Implement binary search on a sorted array.

// 13. Find the kth largest element in an array.

// 14. Implement a function to merge two sorted linked lists.

// 15. Implement a circular queue with enqueue and dequeue operations.
func main() {
	fmt.Println("Hello")
}
