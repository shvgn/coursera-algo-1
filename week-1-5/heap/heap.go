package main

import (
	"fmt"
)

type heap struct {
	elems []int
}

func childrenIndexes(i int) (int, int) {
	return (i+1)*2 - 1, (i + 1) * 2
}

// Returns the element and status (found ot not)
func (h *heap) elem(i int) (int, bool) {
	size := len(h.elems)
	if size <= i {
		return 0, false
	}

	return h.elems[i], true
}

// indexes of children in the heap slice ant their search statuses (left, right, foundLeft , foundRight)
func (h *heap) children(i int) (int, int, bool, bool) {
	li, ri := childrenIndexes(i)

	leftChild, leftOk := h.elem(li)
	rightChild, rightOk := h.elem(ri)

	return leftChild, rightChild, leftOk, rightOk
}

func (h *heap) parent(i int) int {
	return h.elems[i/2]
}

func (h *heap) root() (int, bool) {
	return h.elem(0)
}

// TODO: implement
func (h *heap) extractMin() int {
	return 0
}

func (h *heap) insert(x int) {
	h.elems = append(h.elems, x)
	size := len(h.elems)
	parent := h.parent(size - 1)

	childIdx := size - 1
	parentIdx := (size - 1) / 2
	child, parent := h.elems[childIdx], h.elems[parentIdx]

	for parent > x {
		// Bubble-up `x`: swap parent and child to restore heap property
		h.elems[childIdx] = parent
		h.elems[parentIdx] = child

		parent = h.parent(parentIdx)
		parentIdx /= 2
		childIdx /= 2
	}
}

func main() {
	h := heap{elems: []int{4, 4, 8, 9, 4, 12, 9, 11, 13}}
	fmt.Println(h)

	nums := []int{12, 20, 5}
	for _, x := range nums {
		h.insert(x)
		fmt.Println(h)
	}

}
