package list

import (
	"sync"

	"golang.org/x/exp/constraints"

	"github.com/soyart/gsl/data"
)

const (
	MaxHeap data.SortOrder = data.Descending
	MinHeap data.SortOrder = data.Ascending
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

func NewHeapCmp[T CmpOrdered[T]](
	order data.SortOrder,
) *Heap[T] {
	return &Heap[T]{
		Ordering: order,
		lessFunc: lessFuncCmp[T],
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
	lastIdx := len(h.Items) - 1

	h.Items[0] = h.Items[lastIdx]
	h.Items = h.Items[:lastIdx]

	h.heapifyDown(0)

	return root
}

func (h *Heap[T]) Len() int {
	h.mut.Lock()
	defer h.mut.Unlock()

	return len(h.Items)
}

func (h *Heap[T]) IsEmpty() bool {
	h.mut.RLock()
	defer h.mut.RUnlock()

	return len(h.Items) == 0
}

func (h *Heap[T]) Peek() data.GetValuer[T] {
	h.mut.RLock()
	defer h.mut.RUnlock()

	return h.Items[0]
}

func (h *Heap[T]) PopValue() T {
	return h.Pop().GetValue()
}

func (h *Heap[T]) PeekValue() T {
	return h.Peek().GetValue()
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
	length := len(h.Items)

	for {
		childLeft := leftChildNode(curr)
		if childLeft >= length {
			break
		}

		childRight := childLeft + 1

		// Child to compare
		child := -1

		// Choose left child if:
		// 1) right child is null (out of range)
		// 2) left child has higher priority (lessFunc -> true)
		//
		// Otherwise use right child
		switch {
		case
			childRight >= length,
			h.lessFunc(h.Items, h.Ordering, childLeft, childRight):

			child = childLeft

		default:
			child = childRight
		}

		if h.lessFunc(h.Items, h.Ordering, curr, child) {
			break
		}

		h.swap(curr, child)

		curr = child
	}
}

func (h *Heap[T]) swap(i, j int) {
	h.Items[i], h.Items[j] = h.Items[j], h.Items[i]
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

// Less implementation for constraints.Ordered
func lessFuncOrdered[T constraints.Ordered](
	items []data.GetValuer[T],
	order data.SortOrder,
	i int,
	j int,
) bool {
	if order == MinHeap {
		return items[i].GetValue() < items[j].GetValue()
	}

	return items[i].GetValue() > items[j].GetValue()
}

// Less implementation for CmpOrdered, e.g. *big.Int and *big.Float, and other lib types.
func lessFuncCmp[T CmpOrdered[T]](
	items []data.GetValuer[T],
	order data.SortOrder,
	i int,
	j int,
) bool {
	var cmp int

	switch order {
	case MinHeap:
		cmp = -1

	case MaxHeap:
		cmp = 1
	}

	return items[i].GetValue().Cmp(items[j].GetValue()) == cmp
}
