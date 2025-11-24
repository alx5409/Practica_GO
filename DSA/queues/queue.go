package queue

import "errors"

// Queue DSA Practice Exercises:
//
// 1. Implement a generic queue with Enqueue, Dequeue, Peek, and IsEmpty methods.
type Queue[V any] struct {
	data []V
}

func (q Queue[V]) IsEmpty() bool {
	return len(q.data) == 0
}

func (q Queue[V]) Peek() (V, error) {
	if q.IsEmpty() {
		var zero V
		return zero, errors.New("Error, the queue is empty")
	}
	return q.data[0], nil
}

func (q *Queue[V]) Enqueue(value V) {
	q.data = append(q.data, value)
}

func (q *Queue[V]) Dequeue() []V {
	if !q.IsEmpty() {
		q.data = q.data[1:]
	}
	return q.data
}

//
// 2. Reverse a queue using recursion.
func reverseQueue[V any](queue Queue[V]) Queue[V] {
	reversedQueue := Queue[V]{}
	for !queue.IsEmpty() {
		value := queue.data[len(queue.data)-1]
		reversedQueue.Enqueue(value)
		queue.data = queue.data[:len(queue.data)-1]
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
//
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
