package stackqueue

type node struct {
	value interface{}
	next  *node
}

type Stack struct {
	Top    *node
	Length int
}

func (s *Stack) Push(val interface{}) {
	s.Length++
	newNode := &node{value: val}
	newNode.next = s.Top
	s.Top = newNode
}

func (s *Stack) Pop() interface{} {
	popNode := s.Top
	if popNode == nil {
		return nil
	}
	s.Length--
	s.Top = popNode.next
	return popNode.value
}

func (s *Stack) Peek() interface{} {
	return s.Top.value
}
