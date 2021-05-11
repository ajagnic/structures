package tree

type Node struct {
	Left  *Node
	Right *Node
	Value int
}

type BinarySearchTree struct {
	Root *Node
}

func (t *BinarySearchTree) Insert(val int) {
	newNode := &Node{Value: val}
	if t.Root == nil {
		t.Root = newNode
		return
	}
	thisNode := t.Root
	for {
		if val < thisNode.Value {
			if thisNode.Left == nil {
				thisNode.Left = newNode
				break
			}
			thisNode = thisNode.Left
		} else if val > thisNode.Value {
			if thisNode.Right == nil {
				thisNode.Right = newNode
				break
			}
			thisNode = thisNode.Right
		}
	}
}

func (t *BinarySearchTree) Lookup(val int) *Node {
	thisNode := t.Root
	for thisNode != nil {
		if val == thisNode.Value {
			break
		}
		if val < thisNode.Value {
			thisNode = thisNode.Left
			continue
		}
		thisNode = thisNode.Right
	}
	return thisNode
}

func (t *BinarySearchTree) Remove(val int) {
	thisNode := &t.Root
	for *thisNode != nil {
		n := *thisNode
		if n.Value == val {
			if ans, child := n.hasZeroOrOneChild(); ans == true {
				*thisNode = child
				break
			}
			// Traverse to the next lowest value after thisNode
			s := &n.Right
			successor := *s
			for {
				if successor.Left == nil {
					break
				}
				s = &successor.Left
				successor = *s
			}
			// Reassign children, and replace thisNode with successor
			*s = successor.Right
			successor.Left, successor.Right = n.Left, n.Right
			*thisNode = successor
			break
		}

		if val < n.Value {
			thisNode = &n.Left
			continue
		}
		thisNode = &n.Right
	}
}

func (n *Node) hasZeroOrOneChild() (bool, *Node) {
	if n.Left == nil && n.Right == nil {
		return true, nil
	} else if n.Left != nil && n.Right == nil {
		return true, n.Left
	} else if n.Right != nil && n.Left == nil {
		return true, n.Right
	}
	return false, nil
}
