package linkedlist_test

import (
	"testing"

	"github.com/ajagnic/structures/linkedlist"
)

func TestSingleAppend(t *testing.T) {
	ll := linkedlist.LinkedList{}
	// Verify head and tail are set when there is only one node.
	first := 10
	ll.Append(first)

	head := ll.Head.Value.(int)
	tail := ll.Tail.Value.(int)
	if head != first || tail != first || ll.Length != 1 {
		t.Errorf("Head/Tail/Length not valid:\nHead:%v\nTail:%v\nLength:%v\n", head, tail, ll.Length)
	}
	// Check all nodes are added in order.
	second := 20
	third := 30
	ll.Append(second)
	ll.Append(third)

	if ll.Length != 3 {
		t.Errorf("Incorrect length: %v!=%v\n", ll.Length, 3)
	}

	nodes := [3]int{}
	cntr := 0
	n := ll.Head
	for n.Next != nil {
		nodes[cntr] = n.Value.(int)
		cntr++
		n = n.Next
	}
	nodes[cntr] = n.Value.(int)

	correct := [3]int{first, second, third}
	if nodes != correct {
		t.Errorf("Incorrect number/value of nodes:\nValues:%v\nCorrect:%v\n", nodes, correct)
	}
}

func TestSinglePrepend(t *testing.T) {
	ll := linkedlist.LinkedList{}
	// Verify head is set when null.
	first := "one"
	ll.Prepend(first)

	if ll.Head.Value.(string) != first || ll.Length != 1 {
		t.Errorf("Head not set on null prepend\n")
	}
	// Check all nodes are in the correct order.
	second := "two"
	newFirst := "zero"
	ll.Append(second)
	ll.Prepend(newFirst)

	if ll.Length != 3 {
		t.Errorf("Incorrect length: %v!=%v\n", ll.Length, 3)
	}

	nodes := [3]string{}
	cntr := 0
	n := ll.Head
	for n.Next != nil {
		nodes[cntr] = n.Value.(string)
		cntr++
		n = n.Next
	}
	nodes[cntr] = n.Value.(string)

	correct := [3]string{newFirst, first, second}
	if nodes != correct {
		t.Errorf("Incorrect number/value of nodes:\nValues:%v\nCorrect:%v\n", nodes, correct)
	}
}

func TestSingleDelete(t *testing.T) {
	type testObj struct {
		ID    int
		Data  string
		Valid bool
	}
	ll := linkedlist.LinkedList{}
	// Verify can call delete on null list.
	ll.Delete(0)
	// Verify can delete the only item in list.
	bad := testObj{0, "deleteme", false}
	ll.Append(bad)
	ll.Delete(bad)

	if ll.Head != nil || ll.Length != 0 {
		t.Errorf("Error deleting only item in list:\nHead:%v\nLength:%v\n", ll.Head, ll.Length)
	}
	// Verify head can be deleted and set to the next node.
	first := testObj{1, "one", true}
	second := testObj{2, "two", true}
	ll.Append(first)
	ll.Append(second)
	ll.Delete(first)

	if ll.Head.Value.(testObj) != second || ll.Length != 1 {
		t.Errorf("Head was not reset to the next node when deleted\n")
	}
	// Verify tail can be deleted and set to the previous node.
	ll.Prepend(first)
	ll.Delete(second)

	if ll.Tail.Value.(testObj) != first || ll.Length != 1 {
		t.Errorf("Tail was not reset to the previous node when deleted\n")
	}
	// Verify item can be deleted from middle of list.
	third := testObj{3, "three", true}
	ll.Append(second)
	ll.Append(third)
	ll.Delete(second)

	if ll.Head.Next.Value.(testObj) != third || ll.Length != 2 {
		t.Errorf("Error deleting item from middle of list\n")
	}
}

func TestSingleInsert(t *testing.T) {
	ll := linkedlist.LinkedList{}
	// Verify can insert into head of list.
	first := 10
	newFirst := 1
	ll.Append(first)
	ll.Insert(0, newFirst)

	if ll.Head.Value.(int) != newFirst || ll.Length != 2 {
		t.Errorf("Error inserting at index 0\n")
	}
	// Verify can insert at tail of list.
	second := 20
	ll.Insert(99, second)

	if ll.Tail.Value.(int) != second || ll.Length != 3 {
		t.Errorf("Error inserting at index > length\n")
	}
	// Verify can insert into middle of list.
	inserted := 5
	ll.Insert(1, inserted)

	nodes := [4]int{}
	cntr := 0
	n := ll.Head
	for n.Next != nil {
		nodes[cntr] = n.Value.(int)
		cntr++
		n = n.Next
	}
	nodes[cntr] = n.Value.(int)

	correct := [4]int{newFirst, inserted, first, second}
	if nodes != correct {
		t.Errorf("Incorrect number/value of nodes:\nValues:%v\nCorrect:%v\n", nodes, correct)
	}
}

func TestDoubleAppend(t *testing.T) {
	ll := linkedlist.DoublyLinkedList{}
	// Verify head and tail are set when there is only one node.
	first := 10
	ll.Append(first)

	head := ll.Head.Value.(int)
	tail := ll.Tail.Value.(int)
	if head != first || tail != first || ll.Length != 1 {
		t.Errorf("Head/Tail/Length not valid:\nHead:%v\nTail:%v\nLength:%v\n", head, tail, ll.Length)
	}
	// Check all nodes are added in order.
	second := 20
	third := 30
	ll.Append(second)
	ll.Append(third)

	if ll.Length != 3 {
		t.Errorf("Incorrect length: %v!=%v\n", ll.Length, 3)
	}

	nodes := [3]int{}
	cntr := 0
	n := ll.Tail
	for n.Prev != nil {
		nodes[cntr] = n.Value.(int)
		cntr++
		n = n.Prev
	}
	nodes[cntr] = n.Value.(int)

	correct := [3]int{third, second, first}
	if nodes != correct {
		t.Errorf("Incorrect number/value of nodes:\nValues:%v\nCorrect:%v\n", nodes, correct)
	}
}

func TestDoublePrepend(t *testing.T) {
	ll := linkedlist.DoublyLinkedList{}
	// Verify head is set when null.
	first := "one"
	ll.Prepend(first)

	if ll.Head.Value.(string) != first || ll.Length != 1 {
		t.Errorf("Head not set on null prepend\n")
	}
	// Check all nodes are in the correct order.
	second := "two"
	newFirst := "zero"
	ll.Append(second)
	ll.Prepend(newFirst)

	if ll.Length != 3 {
		t.Errorf("Incorrect length: %v!=%v\n", ll.Length, 3)
	}

	nodes := [3]string{}
	cntr := 0
	n := ll.Tail
	for n.Prev != nil {
		nodes[cntr] = n.Value.(string)
		cntr++
		n = n.Prev
	}
	nodes[cntr] = n.Value.(string)

	correct := [3]string{second, first, newFirst}
	if nodes != correct {
		t.Errorf("Incorrect number/value of nodes:\nValues:%v\nCorrect:%v\n", nodes, correct)
	}
}

func TestDoubleDelete(t *testing.T) {
	type testObj struct {
		ID    int
		Data  string
		Valid bool
	}
	ll := linkedlist.DoublyLinkedList{}
	// Verify can call delete on null list.
	ll.Delete(0)
	// Verify can delete the only item in list.
	bad := testObj{0, "deleteme", false}
	ll.Append(bad)
	ll.Delete(bad)

	if ll.Head != nil || ll.Length != 0 {
		t.Errorf("Error deleting only item in list:\nHead:%v\nLength:%v\n", ll.Head, ll.Length)
	}
	// Verify head can be deleted and set to the next node.
	first := testObj{1, "one", true}
	second := testObj{2, "two", true}
	ll.Append(first)
	ll.Append(second)
	ll.Delete(first)

	if ll.Head.Value.(testObj) != second || ll.Length != 1 {
		t.Errorf("Head was not reset to the next node when deleted\n")
	}
	// Verify tail can be deleted and set to the previous node.
	ll.Prepend(first)
	ll.Delete(second)

	if ll.Tail.Value.(testObj) != first || ll.Length != 1 {
		t.Errorf("Tail was not reset to the previous node when deleted\n")
	}
	// Verify item can be deleted from middle of list.
	third := testObj{3, "three", true}
	ll.Append(second)
	ll.Append(third)
	ll.Delete(second)

	if ll.Length != 2 {
		t.Errorf("Incorrect length: %v!=%v\n", ll.Length, 3)
	}

	nodes := [2]testObj{}
	cntr := 0
	n := ll.Tail
	for n.Prev != nil {
		nodes[cntr] = n.Value.(testObj)
		cntr++
		n = n.Prev
	}
	nodes[cntr] = n.Value.(testObj)

	correct := [2]testObj{third, first}
	if nodes != correct {
		t.Errorf("Incorrect number/value of nodes:\nValues:%v\nCorrect:%v\n", nodes, correct)
	}
}

func TestDoubleInsert(t *testing.T) {
	ll := linkedlist.DoublyLinkedList{}
	// Verify can insert into head of list.
	first := 10
	newFirst := 1
	ll.Append(first)
	ll.Insert(0, newFirst)

	if ll.Head.Value.(int) != newFirst || ll.Length != 2 {
		t.Errorf("Error inserting at index 0\n")
	}
	// Verify can insert at tail of list.
	second := 20
	ll.Insert(99, second)

	if ll.Tail.Value.(int) != second || ll.Length != 3 {
		t.Errorf("Error inserting at index > length\n")
	}
	// Verify can insert into first and second half of list.
	third := 30
	ll.Insert(99, third)
	insertedFirstHalf := 5
	insertedSecondHalf := 15
	ll.Insert(1, insertedFirstHalf)
	ll.Insert(3, insertedSecondHalf)

	if ll.Length != 6 {
		t.Errorf("Incorrect length: %v!=%v\n", ll.Length, 3)
	}

	nodes := [6]int{}
	cntr := 0
	n := ll.Tail
	for n.Prev != nil {
		nodes[cntr] = n.Value.(int)
		cntr++
		n = n.Prev
	}
	nodes[cntr] = n.Value.(int)

	correct := [6]int{third, second, insertedSecondHalf, first, insertedFirstHalf, newFirst}
	if nodes != correct {
		t.Errorf("Incorrect number/value of nodes:\nValues:%v\nCorrect:%v\n", nodes, correct)
	}
}
