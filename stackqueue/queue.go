package stackqueue

import "github.com/ajagnic/structures/linkedlist"

type Queue struct {
	ll linkedlist.LinkedList
}

func (q *Queue) Enqueue(val interface{}) {
	q.ll.Append(val)
}

func (q *Queue) Dequeue() interface{} {
	first := q.ll.Head
	if first == nil {
		return nil
	}
	q.ll.Length--
	q.ll.Head = first.Next
	return first.Value
}

func (q *Queue) Peek() interface{} {
	return q.ll.Head.Value
}
