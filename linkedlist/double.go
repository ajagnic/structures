package linkedlist

type dllNode struct {
	Value interface{}
	Prev  *dllNode
	Next  *dllNode
}

type DoublyLinkedList struct {
	Head   *dllNode
	Tail   *dllNode
	Length int
}

func (l *DoublyLinkedList) Append(val interface{}) {
	l.Length++
	newNode := &dllNode{Value: val}
	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
		return
	}
	newNode.Prev = l.Tail
	l.Tail.Next = newNode
	l.Tail = newNode
}

func (l *DoublyLinkedList) Prepend(val interface{}) {
	if l.Head == nil {
		l.Append(val)
		return
	}
	l.Length++
	newNode := &dllNode{Value: val, Next: l.Head}
	l.Head.Prev = newNode
	l.Head = newNode
}

func (l *DoublyLinkedList) Delete(val interface{}) {
	if l.Head == nil {
		return
	} else if l.Head.Value == val {
		l.Length--
		if l.Head = l.Head.Next; l.Head != nil {
			l.Head.Prev = nil
		}
		return
	} else if l.Tail.Value == val {
		l.Length--
		l.Tail = l.Tail.Prev
		l.Tail.Next = nil
		return
	}
	thisNode := l.Head.Next
	for thisNode.Next != nil {
		if thisNode.Value == val {
			l.Length--
			thisNode.Prev.Next = thisNode.Next
			thisNode.Next.Prev = thisNode.Prev
			break
		}
		thisNode = thisNode.Next
	}
}

func (l *DoublyLinkedList) Insert(i int, val interface{}) {
	switch {
	case i <= 0:
		l.Prepend(val)
	case i >= l.Length:
		l.Append(val)
	case i < l.Length/2:
		newNode := &dllNode{Value: val}
		thisNode := l.Head.Next
		cntr := 1
		for thisNode.Next != nil {
			if cntr == i {
				l.Length++
				insertAt(thisNode, newNode)
				break
			}
			cntr++
			thisNode = thisNode.Next
		}
	case i >= l.Length/2:
		newNode := &dllNode{Value: val}
		thisNode := l.Tail
		cntr := l.Length - 1
		for thisNode.Prev != nil {
			if cntr == i {
				l.Length++
				insertAt(thisNode, newNode)
				break
			}
			cntr--
			thisNode = thisNode.Prev
		}
	}
}

func insertAt(replace, new *dllNode) {
	new.Next = replace
	new.Prev = replace.Prev
	replace.Prev.Next = new
	replace.Prev = new
}
