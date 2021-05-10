package linkedlist

type node struct {
	Value interface{}
	Next  *node
}

type LinkedList struct {
	Head   *node
	Tail   *node
	Length int
}

func (l *LinkedList) Append(val interface{}) {
	l.Length++
	newNode := &node{Value: val}
	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
		return
	}
	l.Tail.Next = newNode
	l.Tail = newNode
}

func (l *LinkedList) Prepend(val interface{}) {
	if l.Head == nil {
		l.Append(val)
		return
	}
	l.Length++
	newNode := &node{Value: val, Next: l.Head}
	l.Head = newNode
}

func (l *LinkedList) Delete(val interface{}) {
	if l.Head == nil {
		return
	} else if l.Head.Value == val {
		l.Length--
		l.Head = l.Head.Next
		return
	}
	thisNode := l.Head
	for thisNode.Next != nil {
		nextNode := thisNode.Next
		if nextNode.Value == val {
			l.Length--
			if nextNode.Next == nil { // Removing tail
				thisNode.Next = nil
				l.Tail = thisNode
				break
			}
			thisNode.Next = nextNode.Next
			break
		}
		thisNode = nextNode
	}
}

func (l *LinkedList) Insert(i int, val interface{}) {
	switch {
	case i <= 0:
		l.Prepend(val)
	case i >= l.Length:
		l.Append(val)
	default:
		newNode := &node{Value: val}
		thisNode := l.Head
		cntr := 0
		for thisNode.Next != nil {
			if cntr+1 == i { // Insert after thisNode
				l.Length++
				newNode.Next = thisNode.Next
				thisNode.Next = newNode
				break
			}
			cntr++
			thisNode = thisNode.Next
		}
	}
}

func (l *LinkedList) Reverse() {
	var prevNode *node
	thisNode := l.Head
	for thisNode.Next != nil {
		nextNode := thisNode.Next
		thisNode.Next = prevNode
		prevNode = thisNode
		thisNode = nextNode
	}
	l.Tail.Next = prevNode
	l.Head, l.Tail = l.Tail, l.Head
}
