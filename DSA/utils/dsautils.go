package utils

import "errors"

// Stack structure
type Stack[V any] struct {
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

// Queue structure
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
