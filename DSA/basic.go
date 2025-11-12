package dsa

import (
	"errors"
	"fmt"
)

// 1. Rotate a slice to the right by k steps.
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

// 2. Find the longest common prefix among a slice of strings.
func character_position_is_in_every_string(stringSlice []string, char rune, position int) bool {
	if len(stringSlice) == 0 {
		return false
	}
	for _, wordString := range stringSlice {
		word := []rune(wordString)
		if word[position] != char {
			return false
		}
	}
	return true
}

func longest_common_prefix(stringSlice []string) (string, error) {
	prefix := ""
	if len(stringSlice) == 0 {
		er := errors.New("There are no strings")
		return prefix, er
	}
	firstString := stringSlice[0]
	firstWord := []rune(firstString)
	for index, char := range firstWord {
		if !(character_position_is_in_every_string(stringSlice, char, index)) {
			break
		}
		prefix += string(char)
	}
	return prefix, nil
}

// 3. Implement a function to check if two strings are anagrams.
func countLetters(word string) (map[rune]int, error) {
	letterCounter := make(map[rune]int)
	if len(word) == 0 {
		er := errors.New("Empty string")
		return letterCounter, er
	}
	runes := []rune(word)
	for _, char := range runes {
		letterCounter[char]++
	}
	return letterCounter, nil
}

func areMapEqual(map1 map[rune]int, map2 map[rune]int) bool {
	for key1, value1 := range map1 {
		value2, ok := map2[key1]
		if !ok {
			return false
		}
		if value1 != value2 {
			return false
		}
	}
	for key2, value2 := range map2 {
		value1, ok := map1[key2]
		if !ok {
			return false
		}
		if value1 != value2 {
			return false
		}

	}
	return true
}

func areAnagrams(firstWord string, secondWord string) bool {
	firstLetterCounter, _ := countLetters(firstWord)
	secondLetterCounter, _ := countLetters(secondWord)
	if !areMapEqual(firstLetterCounter, secondLetterCounter) {
		return false
	}
	return true
}

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
