package tree

type BinaryHeap struct {
	A []int
}

func (h *BinaryHeap) Insert(val int) {
	h.A = append(h.A, val)
	newIndex := len(h.A) - 1
	heapifyUp(h.A, newIndex)
}

func (h *BinaryHeap) ExtractMax() int {
	max := h.A[0]
	h.A = extract(h.A, 0)
	heapify(h.A, 0)
	return max
}

func (h *BinaryHeap) Delete(i int) {
	current := h.A[i]
	new := h.A[len(h.A)-1]
	h.A = extract(h.A, i)
	if new > current {
		heapifyUp(h.A, i)
	} else if new < current {
		heapify(h.A, i)
	}
}

func (h *BinaryHeap) ModifyValue(i, newValue int) {
	current := h.A[i]
	h.A[i] = newValue
	if newValue > current {
		heapifyUp(h.A, i)
	} else if newValue < current {
		heapify(h.A, i)
	}
}

func (h *BinaryHeap) Search(val int) (i int, ok bool) {
	for i = 0; i < len(h.A); i++ {
		if h.A[i] == val {
			return i, true
		}
	}
	return
}

func heapify(a []int, i int) {
	for {
		left, right := i*2+1, i*2+2
		if left >= len(a) {
			break
		}
		child := largest(a, left, right)
		if a[i] < a[child] {
			a[i], a[child] = a[child], a[i]
		}

		i = child
	}
}

func heapifyUp(a []int, i int) {
	parent := parentIdx(i)
	for i > 0 && a[i] > a[parent] {
		a[i], a[parent] = a[parent], a[i]

		i = parent
		parent = parentIdx(i)
	}
}

func extract(a []int, i int) []int {
	last := len(a) - 1
	a[i], a[last] = a[last], a[i]
	a = a[:last]
	return a
}

func largest(a []int, x, y int) int {
	size := len(a)
	if x >= size && y < size {
		return y
	} else if x < size && y >= size {
		return x
	} else if a[x] > a[y] {
		return x
	} else {
		return y
	}
}

func parentIdx(i int) int {
	return (i+1)/2 - 1
}
