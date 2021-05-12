# structures
__A collection of various data structures implemented in Go.__


## Linked List
![](_readme/linkedlist.png)

```go
package linkedlist // import "github.com/ajagnic/structures/linkedlist"

type node struct {
	Value interface{}
	Next  *node
}

type LinkedList struct {
        Head   *node
        Tail   *node
        Length int
}

func (l *LinkedList) Append(val interface{})
func (l *LinkedList) Delete(val interface{})
func (l *LinkedList) Insert(i int, val interface{})
func (l *LinkedList) Prepend(val interface{})
```


## Doubly Linked List
![](_readme/doublylinkedlist.png)

```go
package linkedlist // import "github.com/ajagnic/structures/linkedlist"

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

func (l *DoublyLinkedList) Append(val interface{})
func (l *DoublyLinkedList) Delete(val interface{})
func (l *DoublyLinkedList) Insert(i int, val interface{})
func (l *DoublyLinkedList) Prepend(val interface{})
```


## Stack
![](_readme/stack.png)

```go
package stackqueue // import "github.com/ajagnic/structures/stackqueue"

type node struct {
	value interface{}
	next  *node
}

type Stack struct {
        Top    *node
        Length int
}

func (s *Stack) Peek() interface{}
func (s *Stack) Pop() interface{}
func (s *Stack) Push(val interface{})
```


## Queue
![](_readme/queue.png)

```go
package stackqueue // import "github.com/ajagnic/structures/stackqueue"

type Queue struct {
	ll linkedlist.LinkedList
}

func (q *Queue) Dequeue() interface{}
func (q *Queue) Enqueue(val interface{})
func (q *Queue) Peek() interface{}
```


## Binary Search Tree
![](_readme/bst.png)

```go
package tree // import "github.com/ajagnic/structures/tree"

type Node struct {
	Left  *Node
	Right *Node
	Value int
}

type BinarySearchTree struct {
        Root *Node
}

func (t *BinarySearchTree) Insert(val int)
func (t *BinarySearchTree) Lookup(val int) *Node
func (t *BinarySearchTree) Remove(val int)
```


## Binary Heap
![](_readme/heap.png)

```go
package tree // import "github.com/ajagnic/structures/tree"

type BinaryHeap struct {
        A []int
}

func (h *BinaryHeap) Delete(i int)
func (h *BinaryHeap) ExtractMax() int
func (h *BinaryHeap) Insert(val int)
func (h *BinaryHeap) ModifyValue(i, newValue int)
func (h *BinaryHeap) Search(val int) (i int, ok bool)
```


## Trie
![](_readme/trie.png)

```go
package tree // import "github.com/ajagnic/structures/tree"

type trieNode struct {
	children    [26]*trieNode
	isEndOfWord bool
}

type Trie struct {
        Root *trieNode
}

func (t *Trie) Delete(key string)
func (t *Trie) Insert(key string)
func (t *Trie) Keys() []string
func (t *Trie) Search(key string) bool
```

# Authors
Adrian Agnic [ [GitHub](https://github.com/ajagnic) ]
