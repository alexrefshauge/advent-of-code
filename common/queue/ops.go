package queue

func (q *Queue[T]) Push(element T) {
	node := &node[T]{
		data: element,
		next: nil,
	}
	if q.head == nil { q.head = node }
	if q.tail != nil {
		q.tail.next = node
	}
	q.tail = node
	q.size++
}

func (q *Queue[T]) Next() T {
	head := q.head
	if q.head == q.tail {
		q.head = nil
		q.tail = nil
	} else {
		q.head = head.next
	}
	q.size--
	return head.data
}

func (q *Queue[T]) Size() int {
	return q.size
}

func (q *Queue[T]) All() []T {
	all := make([]T, q.size, q.size)
	node := q.head
	for range q.size {
		all = append(all, node.data)
		node = node.next
	}
	return all
}