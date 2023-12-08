package list

import (
	"fmt"
	"sync"

	"golang.org/x/exp/constraints"

	"github.com/soyart/gsl/data"
)

type Heap[T any] struct {
	Items    []data.GetValuer[T]
	Ordering data.SortOrder

	lessFunc func(items []data.GetValuer[T], t data.SortOrder, i, j int) bool
	mut      *sync.RWMutex
}

func NewHeap[T constraints.Ordered](
	order data.SortOrder,
) *Heap[T] {
	return &Heap[T]{
		Ordering: order,
		lessFunc: lessFuncOrdered[T],
		mut:      new(sync.RWMutex),
	}
}

func (h *Heap[T]) Push(item T) {
	h.mut.Lock()
	defer h.mut.Unlock()

	h.Items = append(h.Items, data.NewGetValuer[T](item))
	h.heapifyUp(len(h.Items) - 1)
}

func (h *Heap[T]) Pop() data.GetValuer[T] {
	h.mut.Lock()
	defer h.mut.Unlock()

	root := h.Items[0]
	h.swap(0, len(h.Items)-1)
	h.Items = h.Items[1:]

	if len(h.Items) < 2 {
		h.Items = []data.GetValuer[T]{}
		return root
	}

	h.heapifyDown(0)

	return root
}

func (h *Heap[T]) heapifyUp(from int) {
	curr := from
	for curr != 0 {
		parent := parentNode(curr)

		if !h.lessFunc(h.Items, h.Ordering, curr, parent) {
			break
		}

		h.swap(curr, parent)
		curr = parent
	}
}

func (h *Heap[T]) heapifyDown(from int) {
	curr := from

	for leftChildNode(curr) < len(h.Items) {
		childLeft := leftChildNode(curr)
		childRight := childLeft + 1

		// Select child to compare
		child := -1
		if h.lessFunc(h.Items, h.Ordering, childLeft, childRight) {
			child = childLeft
		}
		if child == -1 {
			child = childRight
		}

		fmt.Println("from", from, "curr", curr, "l", childLeft, "r", childRight, "child", child)

		if child >= len(h.Items) {
			break
		}

		if !h.lessFunc(h.Items, h.Ordering, curr, child) {
			break
		}

		h.swap(curr, child)

		curr = child
	}
}

func (h *Heap[T]) swap(i, j int) {
	h.Items[j], h.Items[i] = h.Items[i], h.Items[j]
}

func leftChildNode(parent int) int {
	return (2 * parent) + 1
}

func parentNode(child int) int {
	rightChild := child%2 == 0
	if rightChild {
		return (child - 2) / 2
	}

	return (child - 1) / 2
}
