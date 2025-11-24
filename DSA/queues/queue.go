package queue

import "errors"

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

//
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
