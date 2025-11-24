package stack

import (
	queue "Practica_GO/DSA/queues"
	"errors"
	"fmt"
	"strings"
)

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
		return zero, errors.New("error, the stack is empty")
	}
	return s.data[len(s.data)-1], nil
}

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

// General number type
type Number interface {
	~int | ~int32 | ~int64 | ~float32 | ~float64
}

func sum[N Number](num1 N, num2 N) (N, error) {
	return num1 + num2, nil
}
func subtract[N Number](num1 N, num2 N) (N, error) {
	return num1 - num2, nil
}
func multiply[N Number](num1 N, num2 N) (N, error) {
	return num1 * num2, nil
}
func divide[N Number](num1 N, num2 N) (N, error) {
	if num2 == 0 {
		var zero N
		return zero, errors.New("division by zero")
	}
	return num1 / num2, nil
}

func isOperator(r rune) bool {
	return r == '+' || r == '-' || r == '*' || r == '/'
}

func whichOperation[N Number](r rune, num1 N, num2 N) (N, error) {
	switch r {
	case '+':
		return sum(num1, num2)
	case '-':
		return subtract(num1, num2)
	case '*':
		return multiply(num1, num2)
	case '/':
		return divide(num1, num2)
	default:
		var zero N
		return zero, errors.New("invalid operator")
	}
}

func evaluatePostfix[N Number](input string) (N, error) {
	var stackNumbers Stack[N]
	elements := strings.Split(input, " ")
	for _, element := range elements {
		r := []rune(element)[0]
		// If it is not an operator, it must be a number
		if !isOperator(r) {
			var num N
			_, err := fmt.Sscan(element, &num)
			if err != nil {
				return num, errors.New("could not convert element to number")
			}
			stackNumbers.Push(num)
			continue
		}
		if len(stackNumbers.data) < 2 {
			return 0, errors.New("not enough operands")
		}
		num2, _ := stackNumbers.Pop()
		num1, _ := stackNumbers.Pop()
		operationResult, err := whichOperation(r, num1, num2)
		if err != nil {
			return 0, err
		}
		stackNumbers.Push(operationResult)
	}
	return stackNumbers.Peek()
}

// 5. Sort a stack using only stack operations.
func sortStack[V Number](s *Stack[V]) {
	// Use auxiliar stack to help sorting
	aux := Stack[V]{}
	for !s.IsEmpty() {
		temp, _ := s.Pop()
		// Move elements from aux back to s if they are greater than temp
		for !aux.IsEmpty() {
			top, _ := aux.Peek()
			if top <= temp {
				break
			}
			val, _ := aux.Pop()
			s.Push(val)
		}
		aux.Push(temp)
	}
	// Move sorted elements back to the original stack
	for !aux.IsEmpty() {
		val, _ := aux.Pop()
		s.Push(val)
	}
}

// 6. Implement a Min Stack that supports retrieving the minimum element in constant time.
type minStack[V Number] struct {
	data Stack[V]
	min  Stack[V]
}

func (m *minStack[V]) Push(val V) {
	m.data.Push(val)
	if m.min.IsEmpty() {
		m.min.Push(val)
	} else {
		minVal, _ := m.min.Peek()
		if val < minVal {
			m.min.Push(val)
		} else {
			m.min.Push(minVal)
		}
	}
}

func (m *minStack[V]) Pop() (V, bool) {
	val, ok := m.data.Pop()
	if ok {
		m.min.Pop()
	}
	return val, ok
}

func (m minStack[V]) retrieveMin() (V, error) {
	return m.min.Peek()
}

// 7. For each element in a slice, find the next greater element to its right using a stack.
func findNextGreaterElement[N Number](s []N) []N {
	result := make([]N, len(s))
	stack := Stack[int]{}

	for i := len(s) - 1; i >= 0; i-- {
		for !stack.IsEmpty() {
			topIdx, _ := stack.Peek()
			if s[topIdx] <= s[i] {
				stack.Pop()
				continue
			}
			break
		}
		if stack.IsEmpty() {
			var zero N
			result[i] = zero
			stack.Push(i)
			continue
		}
		topIdx, _ := stack.Peek()
		result[i] = s[topIdx]
		stack.Push(i)
	}
	return result
}

// 8. Remove all adjacent duplicates in a string using a stack.
func removeAdjacentDuplicates(s string) string {
	stack := Stack[rune]{}
	for _, char := range s {
		top, _ := stack.Peek()
		if char == top && stack.IsEmpty() {
			stack.Pop()
			continue
		}
		stack.Push(char)
	}
	return string(stack.data)
}

// 9. Implement a stack using two queues.
type StackWithQueues[V any] struct {
	q1 queue.Queue[V]
	q2 queue.Queue[V]
}

func (s *StackWithQueues[V]) PushWithQueues(val V) {
	s.q1.Enqueue(val)
}

func (s *StackWithQueues[V]) PopWithQueues() (V, error) {
	var zero V
	if s.q1.IsEmpty() {
		return zero, errors.New("stack is empty")
	}
	// Move all but the last element from q1 to q2
	for len(s.q1.Data) > 1 {
		val, _ := s.q1.Dequeue()
		s.q2.Enqueue(val)
	}
	// The last element in q1 is the "top" of the stack
	val, _ := s.q1.Dequeue()
	// Swap q1 and q2
	s.q1, s.q2 = s.q2, s.q1
	return val, nil
}

// 10. Given a histogram (slice of heights), find the largest rectangle area using a stack.
func findLargestRectangleInHistograms[N Number](histogram []N) N {
	n := len(histogram)
	maxArea := N(0)
	stack := Stack[int]{}

	for i := 0; i <= n; i++ {
		var currHeight N
		if i == n {
			currHeight = 0
		}
		if i != n {
			currHeight = histogram[i]
		}
		for !stack.IsEmpty() {
			topIdx, _ := stack.Peek()
			if currHeight >= histogram[topIdx] {
				break
			}
			stack.Pop()
			height := histogram[topIdx]
			var width int
			if stack.IsEmpty() {
				width = i
				area := height * N(width)
				if area > maxArea {
					maxArea = area
				}
				continue
			}
			prevIdx, _ := stack.Peek()
			width = i - prevIdx - 1
			area := height * N(width)
			if area > maxArea {
				maxArea = area
			}
		}
		stack.Push(i)
	}
	return maxArea
}
