package stack

import (
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
		return zero, errors.New("Error, the stack is empty")
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
		return zero, errors.New("Division by zero")
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
		return zero, errors.New("Invalid operator")
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
				return num, errors.New("Could not convert element to number")
			}
			stackNumbers.Push(num)
			continue
		}
		if len(stackNumbers.data) < 2 {
			return 0, errors.New("Not enough operands")
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
