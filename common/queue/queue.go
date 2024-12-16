package queue

type node[T any] struct {
	data T
	next *node[T]
}

// Queue
//
// FILO - First In Last Out
type Queue[T any] struct {
	nodes []node[T]
	head *node[T]
	tail *node[T]
	size int
}

func New[T any]() *Queue[T] {
	return &Queue[T]{
		nodes: make([]node[T], 0),
		head: nil,
		tail: nil,
		size: 0,
	}
}