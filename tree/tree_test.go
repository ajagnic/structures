package tree_test

import (
	"testing"

	"github.com/ajagnic/structures/tree"
)

func TestBinaryHeapInsert(t *testing.T) {
	h := tree.BinaryHeap{}
	// Verify the order of values when inserted.
	input := []int{2, 7, 26, 25, 19, 17, 1, 90, 3, 36}
	for _, n := range input {
		h.Insert(n)
	}

	values := [10]int{}
	copy(values[:], h.A)
	correct := [10]int{90, 36, 17, 25, 26, 7, 1, 2, 3, 19}
	if values != correct || len(h.A) != 10 {
		t.Errorf("Incorrect number/value of nodes:\nValues:%v\nCorrect:%v\n", values, correct)
	}
}

func TestBinaryHeapExtractMax(t *testing.T) {
	h := tree.BinaryHeap{}
	// Verify max values are extracted in the correct order.
	input := []int{2, 7, 26, 25, 19, 17, 1, 90, 3, 36}
	for _, n := range input {
		h.Insert(n)
	}

	correct := [10]int{90, 36, 26, 25, 19, 17, 7, 3, 2, 1}
	l := len(h.A)
	for i := 0; i < l; i++ {
		max := h.ExtractMax()
		if correct[i] != max {
			t.Errorf("Max value is not correct:\nValue:%v\nExpected:%v\n", max, correct[i])
		}
	}

	if len(h.A) != 0 {
		t.Errorf("Length not correct:\n%v!=%v\n", len(h.A), 0)
	}
}

func TestBinaryHeapDelete(t *testing.T) {
	h := tree.BinaryHeap{}
	// Verify value is deleted and heap is maintained upwards.
	input := []int{2, 7, 26, 25, 19, 17, 1, 90, 3, 36}
	for _, n := range input {
		h.Insert(n)
	}

	h.Delete(6)
	values := [9]int{}
	copy(values[:], h.A)
	correct := [9]int{90, 36, 19, 25, 26, 7, 17, 2, 3}
	if values != correct || len(h.A) != 9 {
		t.Errorf("(1)Incorrect number/value of nodes:\nValues:%v\nCorrect:%v\n", values, correct)
	}
	// Verify value is deleted and heap is maintained downwards.
	h.Delete(1)
	values2 := [8]int{}
	copy(values2[:], h.A)
	correct2 := [8]int{90, 26, 19, 25, 3, 7, 17, 2}
	if values2 != correct2 || len(h.A) != 8 {
		t.Errorf("(2)Incorrect number/value of nodes:\nValues:%v\nCorrect:%v\n", values, correct)
	}
}

func TestBinaryHeapModifyValue(t *testing.T) {
	h := tree.BinaryHeap{}
	// Verify modified value moves up the tree.
	input := []int{2, 7, 26, 25, 19, 17, 1, 90, 3, 36}
	for _, n := range input {
		h.Insert(n)
	}

	h.ModifyValue(5, 100)
	values := [10]int{}
	copy(values[:], h.A)
	correct := [10]int{100, 36, 90, 25, 26, 17, 1, 2, 3, 19}
	if values != correct || len(h.A) != 10 {
		t.Errorf("(1)Incorrect number/value of nodes:\nValues:%v\nCorrect:%v\n", values, correct)
	}
	// Verify modified value moves down the tree.
	h.ModifyValue(1, 15)
	copy(values[:], h.A)
	correct = [10]int{100, 26, 90, 25, 19, 17, 1, 2, 3, 15}
	if values != correct || len(h.A) != 10 {
		t.Errorf("(2)Incorrect number/value of nodes:\nValues:%v\nCorrect:%v\n", values, correct)
	}
}

func TestBinaryHeapSearch(t *testing.T) {
	h := tree.BinaryHeap{}
	// Verify existing value can be found.
	input := []int{2, 7, 26, 25, 19, 17, 1, 90, 3, 36}
	for _, n := range input {
		h.Insert(n)
	}

	i, ok := h.Search(3)
	if ok == false || i != 8 {
		t.Errorf("Failed to find inserted value:\nOk:%v\nIndex:%v\n", ok, i)
	}
	// Verify Search returns false for non-existent value.
	i, ok = h.Search(99)
	if ok == true || i != len(h.A) {
		t.Errorf("Returned true for non-existent value:\nOk:%v\nIndex:%v\n", ok, i)
	}
}

func TestBinarySearchTreeInsert(t *testing.T) {
	b := tree.BinarySearchTree{}
	// Verify root node is set.
	b.Insert(50)

	if b.Root == nil {
		t.Errorf("Root node is nil\n")
	}
	// Verify nodes are in the correct order once inserted.
	input := []int{20, 80, 10, 30, 70, 90, 15, 25, 75, 85}
	for _, v := range input {
		b.Insert(v)
	}

	values := &[]int{b.Root.Value}
	traverseTree(b.Root, values)

	valuesArr := [11]int{}
	copy(valuesArr[:], *values)

	correct := [11]int{50, 20, 10, 15, 30, 25, 80, 70, 75, 90, 85}
	if valuesArr != correct {
		t.Errorf("Incorrect value of nodes:\nValues:%v\nCorrect:%v\n", valuesArr, correct)
	}
}

func TestBinarySearchTreeLookup(t *testing.T) {
	b := tree.BinarySearchTree{}
	// Verify that an existing value can be found.
	input := []int{50, 20, 80, 10, 30, 70, 90, 15, 25, 75, 85}
	for _, v := range input {
		b.Insert(v)
	}

	find := 90
	n := b.Lookup(find)
	if n.Value != find {
		t.Errorf("Found node value not correct:\n%v!=%v\n", n.Value, find)
	}
	// Verify that a non-existent value returns nil.
	n = b.Lookup(999)
	if n != nil {
		t.Errorf("Non-existent node != nil:\nNode:%v\n", n)
	}
}

func TestBinarySearchTreeRemove(t *testing.T) {
	b := tree.BinarySearchTree{}
	// Verify a node with <2 children can be removed correctly.
	input := []int{50, 20, 80, 10, 30, 90, 15, 25, 85, 27, 26, 28}
	for _, v := range input {
		b.Insert(v)
	}

	removeSingle := 80
	b.Remove(removeSingle)

	values := &[]int{b.Root.Value}
	traverseTree(b.Root, values)
	valuesArr := [11]int{}
	copy(valuesArr[:], *values)

	correct := [11]int{50, 20, 10, 15, 30, 25, 27, 26, 28, 90, 85}
	if valuesArr != correct {
		t.Errorf("(1)Incorrect value of nodes:\nValues:%v\nCorrect:%v\n", valuesArr, correct)
	}
	// Verify node with 2 children can be removed correctly.
	removeDouble := 20
	b.Remove(removeDouble)

	values = &[]int{b.Root.Value}
	traverseTree(b.Root, values)
	valuesArr2 := [10]int{}
	copy(valuesArr2[:], *values)

	correct2 := [10]int{50, 25, 10, 15, 30, 27, 26, 28, 90, 85}
	if valuesArr2 != correct2 {
		t.Errorf("(2)Incorrect value of nodes:\nValues:%v\nCorrect:%v\n", valuesArr, correct)
	}
}

func TestTrieSearch(t *testing.T) {
	tr := tree.Trie{}
	// Verify search doesn't match partial keys.
	tr.Insert("target")
	tr.Insert("tarp")
	tr.Insert("tangent")
	tr.Insert("tangy")

	partials := []string{"t", "tar", "tan", "tangential", "targets"}
	for _, v := range partials {
		res := tr.Search(v)
		if res == true {
			t.Errorf("Search returned true for partial match: %v", v)
		}
	}
	// Verify search finds inserted keys.
	exists := []string{"target", "tarp", "tangent", "tangy"}
	for _, v := range exists {
		res := tr.Search(v)
		if res == false {
			t.Errorf("Search cound not find inserted key: %v", v)
		}
	}
}

func TestTrieKeys(t *testing.T) {
	tr := tree.Trie{}
	// Verify the amount of keys and their values.
	tr.Insert("target")
	tr.Insert("tarp")
	tr.Insert("tangent")
	tr.Insert("tangy")
	tr.Insert("ace")

	keys := tr.Keys()
	keysArr := [5]string{}
	copy(keysArr[:], keys)
	correct := [5]string{"ace", "tangent", "tangy", "target", "tarp"}
	if keysArr != correct || len(keys) != 5 {
		t.Errorf("Incorrect value of keys:\nValues:%v\nCorrect:%v\n", keysArr, correct)
	}
}

func TestTrieDelete(t *testing.T) {
	tr := tree.Trie{}
	tr.Insert("a")
	tr.Insert("tar")
	tr.Insert("target")
	tr.Insert("tarp")
	tr.Insert("tan")
	tr.Insert("tangent")
	tr.Insert("tangy")

	// Verify:
	// Partial match is kept.
	tr.Delete("ace")
	// Match with children is deleted correctly.
	tr.Delete("tar")
	// Match with parents is deleted correctly.
	tr.Delete("tangent")

	keys := tr.Keys()
	keysArr := [5]string{}
	copy(keysArr[:], keys)
	correct := [5]string{"a", "tan", "tangy", "target", "tarp"}
	if keysArr != correct {
		t.Errorf("Incorrect value of keys:\nValues:%v\nCorrect:%v\n", keysArr, correct)
	}
}

func traverseTree(node *tree.Node, a *[]int) {
	if node.Left != nil {
		*a = append(*a, node.Left.Value)
		traverseTree(node.Left, a)
	}
	if node.Right != nil {
		*a = append(*a, node.Right.Value)
		traverseTree(node.Right, a)
	}
}
