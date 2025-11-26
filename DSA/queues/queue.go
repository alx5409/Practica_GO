package queue

import (
	utils "Practica_GO/DSA/utils"
	"errors"
	"fmt"
)

// Queue DSA Practice Exercises:
//
// 1. Implement a generic queue with Enqueue, Dequeue, Peek, and IsEmpty methods.
type Queue[V any] struct {
	Data []V
}

func (q Queue[V]) IsEmpty() bool {
	return len(q.Data) == 0
}

func (q Queue[V]) Peek() (V, error) {
	if q.IsEmpty() {
		var zero V
		return zero, errors.New("the queue is empty")
	}
	return q.Data[0], nil
}

func (q *Queue[V]) Enqueue(value V) {
	q.Data = append(q.Data, value)
}

func (q *Queue[V]) Dequeue() (V, error) {
	var val V
	if q.IsEmpty() {
		return val, errors.New("queue is empty")
	}
	val = q.Data[0]
	q.Data = q.Data[1:]
	return val, nil
}

// 2. Reverse a queue using recursion.
func ReverseQueue[V any](queue Queue[V]) Queue[V] {
	reversedQueue := Queue[V]{}
	for !queue.IsEmpty() {
		value := queue.Data[len(queue.Data)-1]
		reversedQueue.Enqueue(value)
		queue.Data = queue.Data[:len(queue.Data)-1]
	}
	return reversedQueue
}

func recursiveReverseQueue[V any](queue Queue[V]) Queue[V] {
	if queue.IsEmpty() {
		return queue
	}
	front, _ := queue.Peek()
	queue.Dequeue()
	reversed := recursiveReverseQueue(queue)
	reversed.Enqueue(front)
	return reversed
}

// 3. Generate binary numbers from 1 to N using a queue.
func binaryNumbersFrom1ToN(number int) []string {
	var result []string
	zero := "0"
	one := "1"
	if number < 1 {
		return result
	}
	q := Queue[string]{}
	q.Enqueue("1")
	for i := 0; i < number; i++ {
		value, _ := q.Dequeue()
		result = append(result, value)
		q.Enqueue(value + zero)
		q.Enqueue(value + one)
	}
	return result
}

// 4. Implement a circular queue.
type CircularQueue[V any] struct {
	data  []V
	cap   int
	front int
	rear  int
	size  int
}

func NewCircularQueue[V any](capacity int) *CircularQueue[V] {
	return &CircularQueue[V]{
		data:  make([]V, capacity),
		cap:   capacity,
		front: 0,
		rear:  0,
		size:  0,
	}
}

func (c CircularQueue[V]) IsEmpty() bool {
	return c.size == 0
}

func (c *CircularQueue[V]) Enqueue(value V) error {
	if c.size == c.cap {
		return fmt.Errorf("max capacity reached, can not enqueue the value %v", value)
	}
	c.data[c.rear] = value
	c.rear = (c.rear + 1) % c.cap
	c.size++
	return nil
}

func (c *CircularQueue[V]) Dequeue() (V, error) {
	var result V
	if c.IsEmpty() {
		return result, errors.New("queue is empty, could not dequeue")
	}
	result = c.data[c.front]
	c.front = (c.front + 1) % c.cap
	c.size--
	return result, nil
}

func (c *CircularQueue[V]) Peek() (V, error) {
	var result V
	if c.IsEmpty() {
		return result, errors.New("queue is empty")
	}
	result = c.data[c.front]
	return result, nil
}

// 5. Check if a given sequence of brackets can be balanced using a queue.
func checkBalancedBracketsWithQueue(s string) bool {
	queue := Queue[rune]{}
	openBracket := '{'
	closeBracket := '}'
	for _, char := range s {
		if char == openBracket {
			queue.Enqueue(char)
		}
		if char == closeBracket {
			_, err := queue.Dequeue()
			if err != nil {
				return false
			}
		}
	}
	return queue.IsEmpty()
}

// 6. Implement a queue using two stacks.
type QueueWithStacks[V any] struct {
	s1 utils.Stack[V]
	s2 utils.Stack[V]
}

func (q QueueWithStacks[V]) IsEmpty() bool {
	return q.s1.IsEmpty() && q.s2.IsEmpty()
}

func (q *QueueWithStacks[V]) EnqueueWithStacks(value V) {
	q.s1.Push(value)
}

func (q *QueueWithStacks[V]) DequeueWithStacks() (V, error) {
	if q.IsEmpty() {
		var zero V
		return zero, errors.New("queue is empty")
	}
	// If s2 is empty, move all elements from s1 to s2
	if q.s2.IsEmpty() {
		for !q.s1.IsEmpty() {
			val, _ := q.s1.Pop()
			q.s2.Push(val)
		}
	}
	value, _ := q.s2.Pop()
	return value, nil
}

func (q QueueWithStacks[V]) PeekWithStacks() (V, error) {
	if q.IsEmpty() {
		var zero V
		return zero, errors.New("queue is empty")
	}
	// If s2 is empty, move all elements from s1 to s2
	if q.s2.IsEmpty() {
		for !q.s1.IsEmpty() {
			val, _ := q.s1.Pop()
			q.s2.Push(val)
		}
	}
	value, _ := q.s2.Peek()
	return value, nil
}

// 7. Find the first non-repeating character in a string using a queue.

func FirstNonRepeatingCharStream(s string) (rune, error) {
	hmap := make(map[rune]int)
	for _, char := range s {
		hmap[char]++
	}
	for char, count := range hmap {
		if count == 1 {
			return char, nil
		}
	}
	return ' ', errors.New("character not found")
}

func FirstNonRepeatingCharStreamWithQueue(s string) (rune, error) {
	q := Queue[rune]{}
	count := make(map[rune]int)
	for _, char := range s {
		count[char]++
		q.Enqueue(char)
		for !q.IsEmpty() {
			front, _ := q.Peek()
			if count[front] > 1 {
				q.Dequeue()
			} else {
				return front, nil
			}
		}
	}
	return ' ', errors.New("no non-repeating character found")
}

// 8. Implement a priority queue.
type PriorityPair[V any] struct {
	Value    V
	Priority int
}

type PriorityQueue[V any] struct {
	data []PriorityPair[V]
}

func (p PriorityQueue[V]) IsEmpty() bool {
	return len(p.data) == 0
}

func maxInIntSlice(s []int) int {
	if len(s) == 0 {
		return 0
	}
	max := s[0]
	for _, value := range s {
		if max < value {
			max = value
		}
	}
	return max
}

func (p PriorityQueue[V]) FindPriorityPair() (PriorityPair[V], error) {
	if p.IsEmpty() {
		var zero PriorityPair[V]
		return zero, errors.New("queue is empty")
	}
	maxPriority := p.data[0].Priority
	maxValue := p.data[0].Value
	for _, pair := range p.data {
		if maxPriority < pair.Priority {
			maxPriority = pair.Priority
			maxValue = pair.Value
		}
	}
	return PriorityPair[V]{Value: maxValue, Priority: maxPriority}, nil
}

func (p PriorityQueue[V]) Peek() (V, error) {
	if p.IsEmpty() {
		var zero V
		return zero, errors.New("queue is empty")
	}
	PriorityPair, err := p.FindPriorityPair()
	return PriorityPair.Value, err
}

func (p *PriorityQueue[V]) Enqueue(value V, priority int) {
	pair := PriorityPair[V]{Value: value, Priority: priority}
	p.data = append(p.data, pair)
}

func (p PriorityQueue[V]) maxIndexOfPriorityPair() (int, error) {
	if p.IsEmpty() {
		var zero int
		return zero, errors.New("queue is empty")
	}
	maxIndex := 0
	maxPriority := p.data[0].Priority
	for index, pair := range p.data {
		if maxPriority < pair.Priority {
			maxPriority = pair.Priority
			maxIndex = index
		}
	}
	return maxIndex, nil
}

func (p *PriorityQueue[V]) Dequeue() (V, error) {
	if p.IsEmpty() {
		var zero V
		return zero, errors.New("queue is empty")
	}
	i, _ := p.maxIndexOfPriorityPair()
	deletedValue := p.data[i].Value
	p.data = append(p.data[:i], p.data[i+1:]...)
	return deletedValue, nil
}

// 9. Simulate a round-robin scheduler using a queue.
type Task[V any] struct {
	Value    V
	TimeLeft int
}

func RoundRobinScheduler[V any](tasks []Task[V], quantum int) []V {
	q := Queue[Task[V]]{}
	var result []V

	// Initialization of the task queue
	for _, task := range tasks {
		q.Enqueue(task)
	}
	// keeps a quantum amount of time in each task
	for !q.IsEmpty() {
		task, _ := q.Dequeue()
		// Tracks the finished tasks
		if task.TimeLeft < quantum {
			result = append(result, task.Value)
			continue
		}
		// If task is not finished returns to the queue
		task.TimeLeft -= quantum
		q.Enqueue(task)
	}
	return result
}

// 10. Given a sliding window size, find the maximum in each window using a queue.
func findMaxInEachSlideWindow(s []int, windowSize int) []int {
	var result []int
	sLenght := len(s)
	if sLenght == 0 || windowSize <= 0 {
		return result
	}
	if sLenght <= windowSize {
		result = append(result, maxInIntSlice(s))
		return result
	}
	q := Queue[int]{}
	for i := 0; i < windowSize; i++ {
		q.Enqueue(s[i])
	}
	result = append(result, maxInIntSlice(q.Data))
	for i := windowSize; i < sLenght; i++ {
		q.Dequeue()
		q.Enqueue(s[i+1])
		result = append(result, maxInIntSlice(q.Data))
	}
	return result
}
