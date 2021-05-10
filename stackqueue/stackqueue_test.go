package stackqueue_test

import (
	"testing"

	"github.com/ajagnic/structures/stackqueue"
)

func TestStackPush(t *testing.T) {
	s := stackqueue.Stack{}
	// Verify length and check top is not null.
	s.Push(10)
	if s.Top == nil || s.Length != 1 {
		t.Errorf("Top node is nil or length is wrong:\nLength:%v!=%v\n", s.Length, 1)
	}
}

func TestStackPop(t *testing.T) {
	s := stackqueue.Stack{}
	// Check can pop on single item stack.
	first := "one"
	s.Push(first)
	value := s.Pop()
	if value != first || s.Length != 0 || s.Top != nil {
		t.Errorf("Incorrect value/length/top:\nLength:%v\nTop:%v\n%v==%v?\n", s.Length, s.Top, value, first)
	}
	// Verify order of multiple popped items.
	second := "two"
	third := "three"
	s.Push(first)
	s.Push(second)
	s.Push(third)

	values := [3]string{}
	cntr := 0
	for s.Length > 0 {
		values[cntr] = s.Pop().(string)
		cntr++
	}

	correct := [3]string{third, second, first}
	if values != correct {
		t.Errorf("Incorrect values:\nValues:%v\nCorrect:%v\n", values, correct)
	}
}

func TestStackPeek(t *testing.T) {
	type testObj struct {
		ID    int
		Data  string
		Valid bool
	}
	s := stackqueue.Stack{}
	obj := testObj{1, "one", true}
	s.Push(obj)
	value := s.Peek()
	if value != obj || s.Length != 1 {
		t.Errorf("Incorrect value from peek:\n%v!=%v\n", value, obj)
	}
}

func TestQueueDequeue(t *testing.T) {
	q := stackqueue.Queue{}
	// Verify order of dequeued items.
	first, second, third := 10, 20, 30
	q.Enqueue(first)
	q.Enqueue(second)
	q.Enqueue(third)

	values := [3]int{}
	for i := 0; i < 3; i++ {
		values[i] = q.Dequeue().(int)
	}

	correct := [3]int{first, second, third}
	if values != correct {
		t.Errorf("Incorrect values:\nValues:%v\nCorrect:%v\n", values, correct)
	}
	// Verify queue is empty.
	empty := q.Dequeue()
	if empty != nil {
		t.Errorf("Queue not empty: %v\n", empty)
	}
}

func TestQueuePeek(t *testing.T) {
	type testObj struct {
		ID    int
		Data  string
		Valid bool
	}
	s := stackqueue.Queue{}
	obj := testObj{1, "one", true}
	s.Enqueue(obj)
	value := s.Peek()
	if value != obj {
		t.Errorf("Incorrect value from peek:\n%v!=%v\n", value, obj)
	}
}
