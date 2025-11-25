package queue

import (
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
//
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
//
// 6. Implement a queue using two stacks.
//
// 7. Find the first non-repeating character in a stream using a queue.
//
// 8. Implement a priority queue.
//
// 9. Simulate a round-robin scheduler using a queue.
//
// 10. Given a sliding window size, find the maximum in each window using a queue.
