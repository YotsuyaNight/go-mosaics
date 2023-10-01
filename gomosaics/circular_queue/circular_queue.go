package circular_queue

type CircularQueue[T comparable] struct {
	size int
	head int
	q    []T
}

func NewCircularQueue[T comparable](size int) *CircularQueue[T] {
	return &CircularQueue[T]{
		size: size,
		q:    make([]T, size),
	}
}

func (q *CircularQueue[T]) Contains(elem T) bool {
	for _, val := range q.q {
		if val == elem {
			return true
		}
	}
	return false
}

func (q *CircularQueue[T]) Append(elem T) {
	q.head = (q.head + 1) % q.size
	q.q[q.head] = elem
}
