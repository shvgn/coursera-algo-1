package main

import (
	"fmt"
	"math"
	"strings"
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

func (h *heap) smallestChild(i int) (int, bool) {

	leftChild, rightChild, leftOk, rightOk := h.children(i)
	if leftOk && rightOk {
		return int(math.Min(float64(leftChild), float64(rightChild))), true
	}

	if rightOk {
		return rightChild, true
	}

	if leftOk {
		return leftChild, true
	}

	return 0, false
}

func (h *heap) bubbleDown(i int) {
	current, _ := h.elem(i)

	li, ri := childrenIndexes(i)
	left, leftOk := h.elem(li)
	right, rightOk := h.elem(ri)

	fmt.Printf("bubble-down  current = %2d (%2d)   left = %2d (%2d)   right = %2d (%2d)\n",
		current, i, left, li, right, ri)

	if leftOk && rightOk {
		// take the index of the smalles child
		smallest := left
		si := li
		if left > right {
			si = ri
			smallest = right
		}

		if smallest >= current {
			return
		}

		// swap
		h.elems[si], h.elems[i] = h.elems[i], h.elems[si]
		h.bubbleDown(si)
		return
	}

	if leftOk {
		if left >= current {
			return
		}
		h.elems[li], h.elems[i] = h.elems[i], h.elems[li]
		h.bubbleDown(li)
		return
	}

	if rightOk {
		if right >= current {
			return
		}
		h.elems[ri], h.elems[i] = h.elems[i], h.elems[ri]
		h.bubbleDown(ri)
		return
	}
}

func (h *heap) parent(i int) int {
	return h.elems[i/2]
}

func (h *heap) root() (int, bool) {
	return h.elem(0)
}

func (h *heap) extractMin() int {
	size := len(h.elems)
	last := h.elems[size-1]

	root := h.elems[0]
	h.elems[0] = last

	h.bubbleDown(0)

	h.elems = h.elems[0 : size-1]
	return root
}

func (h *heap) insert(x int) {
	fmt.Println("Insert", x)

	h.elems = append(h.elems, x)
	size := len(h.elems)
	parent := h.parent(size - 1)

	childIdx := size - 1
	parentIdx := (size - 1) / 2
	child, parent := h.elems[childIdx], h.elems[parentIdx]

	for parent > x {
		// Bubble-up `x`: swap parent and child to restore heap property
		fmt.Println("Swap", parent, "and", child)

		h.elems[childIdx] = parent
		h.elems[parentIdx] = child

		parent = h.parent(parentIdx)
		parentIdx /= 2
		childIdx /= 2
	}
}

func (h heap) String() string {

	b := new(strings.Builder)
	levelSize := 1
	for i, n := range h.elems {
		b.WriteString(fmt.Sprintf("%4d", n))

		if (i+2)%levelSize == 0 {
			levelSize *= 2
			b.WriteString(strings.Repeat(" ", 4))
		}
	}
	return b.String()
}

func main() {
	h := heap{elems: []int{4, 4, 8, 9, 4, 12, 9, 11, 13}}
	fmt.Println(h)

	nums := []int{12, 20, 5}
	for _, x := range nums {
		h.insert(x)
		fmt.Println(h)
		// fmt.Println()
	}

	min := h.extractMin()
	fmt.Println("min extracted", min)

	fmt.Println(h)

}
