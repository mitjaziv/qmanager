package structures

import (
	"testing"
)

func Test_Queue(t *testing.T) {
	// Create Queue.
	q := NewQueue()
	if q == nil {
		t.Error("expected Queue not to be nil")
	}

	// Create tasks for test.
	task1 := NewTask()
	task1.Type = "fibonacci"

	task2 := NewTask()
	task2.Type = "fibonacci"

	task3 := NewTask()
	task3.Type = "fibonacci"

	task4 := NewTask()
	task4.Type = "encoder"

	task5 := NewTask()
	task5.Type = "encoder"

	task6 := NewTask()
	task6.Type = "arithmetic"

	task7 := NewTask()
	task7.Type = "arithmetic"

	// Push tasks to Queue.
	q.Push(&task1)
	q.Push(&task2)
	q.Push(&task3)
	q.Push(&task4)
	q.Push(&task5)
	q.Push(&task6)

	// Check size of Queue.
	size := q.Len()
	if size != 6 {
		t.Error("expected Queue size to be 6")
	}

	// Pop first task from Queue.
	t1 := q.Pop()
	if t1 == nil {
		t.Error("expected Task not to be nil")
	}
	if t1.Id != task1.Id {
		t.Error("expected to recive Task1 from Queue")
	}

	// Pop first encoder task from Queue.
	e1 := q.PopByType([]string{"encoder", "arithmetic"})
	if e1 == nil {
		t.Error("expected Task not to be nil")
	}
	if e1.Id != task4.Id {
		t.Error("expected to recive Task4 from Queue")
	}

	// Pop missing Task type from Queue.
	u1 := q.PopByType([]string{"missing"})
	if u1 != nil {
		t.Error("expected Task to be nil")
	}

	// Pop next task from Queue
	t2 := q.Pop()
	if t2 == nil {
		t.Error("expected Task not to be nil")
	}
	if t2.Id != task2.Id {
		t.Error("expected to recive Task2 from Queue")
	}

	// Pop next task from Queue
	t3 := q.Pop()
	if t3 == nil {
		t.Error("expected Task not to be nil")
	}
	if t3.Id != task3.Id {
		t.Error("expected to recive Task3 from Queue")
	}

	// Pop next task from Queue, should be Task5 as we already Poped Task4
	t5 := q.Pop()
	if t5 == nil {
		t.Error("expected Task not to be nil")
	}
	if t5.Id != task5.Id {
		t.Error("expected to recive Task5 from Queue")
	}

	// Pop last task from Queue
	t6 := q.Pop()
	if t6 == nil {
		t.Error("expected Task not to be nil")
	}
	if t6.Id != task6.Id {
		t.Error("expected to recive Task6 from Queue")
	}

	// Push new Task to queue
	q.Push(&task7)

	// Pop last task from Queue
	t7 := q.PopByType([]string{"fibonacci", "arithmetic"})
	if t7 == nil {
		t.Error("expected Task not to be nil")
	}
	if t7.Id != task7.Id {
		t.Error("expected to recive Task7 from Queue")
	}

	// Pop task from empty Queue
	empty1 := q.Pop()
	if empty1 != nil {
		t.Error("expected Task to be nil")
	}

	// Pop task by type from empty Queue
	empty2 := q.PopByType([]string{"encoder"})
	if empty2 != nil {
		t.Error("expected Task to be nil")
	}
}
