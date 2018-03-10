// Queue is thread safe implementation of FIFO list. Nodes contain pointer to Tasks.
// It also contains PopByType method, which will return first node of specific Task type.
package structures

import (
	"sync"
)

type (
	Queue struct {
		first *node
		last  *node
		size  int

		sync.Mutex
	}

	node struct {
		task *Task
		next *node
	}
)

// NewQueue returns FIFO list with pointer to Task in nodes.
func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Push(task *Task) {
	q.Lock()
	defer q.Unlock()

	// Create node.
	n := &node{
		task: task,
	}

	// Add node to end of list. If it is first element mark also first on the list.
	if q.last != nil {
		q.last.next = n
		q.last = n
	} else {
		q.last = n
		q.first = n
	}

	// Update size
	q.size++
}

func (q *Queue) Pop() *Task {
	q.Lock()
	defer q.Unlock()

	// Take first element from list, if exists.
	if q.first == nil {
		return nil
	}
	n := q.first

	// Move first pointer.
	q.first = n.next

	// Check for empty queue.
	if q.first == nil {
		q.last = nil
	}

	// Update size.
	q.size--

	return n.task
}

func (q *Queue) PopByType(types []string) *Task {
	q.Lock()
	defer q.Unlock()

	// Take first element from list, if exists.
	if q.first == nil {
		return nil
	}
	n := q.first

	var np *node

	// Traverse list until right type is found.
	for {
		if n == nil {
			return nil
		}

		// Check if node task types is contained requested types.
		if contains(types, n.task.Type) {
			break
		}

		// Move to next element.
		np = n
		n = n.next
	}

	// Reattach node pointers.
	if np != nil {
		np.next = n.next
	} else {
		q.first = n.next
	}

	// Check for empty queue.
	if q.first == nil {
		q.last = nil
	}

	// Update size
	q.size--

	return n.task
}

func (q *Queue) Len() int {
	q.Lock()
	defer q.Unlock()

	return q.size
}

func contains(slice []string, val string) bool {
	for _, value := range slice {
		if value == val {
			return true
		}
	}
	return false
}
