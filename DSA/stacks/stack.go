package stack

import "errors"

// Stack DSA Practice Exercises:
//
// 1. Implement a generic stack with Push, Pop, Peek, and IsEmpty methods.
type Stack[V comparable] struct {
	data []V
}

func (s Stack[V]) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *Stack[V]) Push(value V) {
	s.data = append(s.data, value)
}

func (s *Stack[V]) Pop() (V, bool) {
	if s.IsEmpty() {
		var zero V
		return zero, false
	}
	val := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return val, true
}

func (s *Stack[V]) Peek() (V, error) {
	if s.IsEmpty() {
		var zero V
		return zero, errors.New("Error, the stack is empty")
	}
	return s.data[len(s.data)-1], nil
}

//
// 2. Reverse a slice using a stack.
func reverseSliceWithStack[V comparable](slice []V) []V {
	var reverserSlice []V
	stack := Stack[V]{}
	// Push the elements in  order into the stack
	for _, element := range slice {
		stack.Push(element)
	}
	for !stack.IsEmpty() {
		value, ok := stack.Pop()
		if !ok {
			continue
		}
		reverserSlice = append(reverserSlice, value)
	}
	return reverserSlice
}

// 3. Check for balanced parentheses in a string using a stack.
func checkBalancedParenthesisWithStack(s string) bool {
	if len(s) == 0 {
		return true
	}
	openParenthesis := '('
	closeParenthesis := ')'
	parenthesisStack := Stack[rune]{}
	for _, char := range s {
		if char == openParenthesis {
			parenthesisStack.Push(openParenthesis)
		}
		if char == closeParenthesis {
			_, ok := parenthesisStack.Pop()
			if !ok {
				return false
			}
		}
	}
	return parenthesisStack.IsEmpty()
}

// 4. Evaluate a postfix (Reverse Polish Notation) expression using a stack.
//
// 5. Sort a stack using only stack operations.
//
// 6. Implement a Min Stack that supports retrieving the minimum element in constant time.
//
// 7. For each element in a slice, find the next greater element to its right using a stack.
//
// 8. Remove all adjacent duplicates in a string using a stack.
//
// 9. Implement a stack using two queues.
//
// 10. Given a histogram (slice of heights), find the largest rectangle area using a stack.
