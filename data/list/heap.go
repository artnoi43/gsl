package list

import (
	"golang.org/x/exp/constraints"

	"github.com/soyart/gsl/data"
)

const (
	MaxHeap data.SortOrder = data.Descending
	MinHeap data.SortOrder = data.Ascending
)

type (
	lessFunc[T any] func(items []T, i, j int) bool
)

type Heap[T any] struct {
	Items    []data.Getter[T]
	CmpFunc  lessFunc[data.Getter[T]]
	Ordering data.SortOrder
}

func NewHeap[T constraints.Ordered](
	order data.SortOrder,
) *Heap[T] {
	return &Heap[T]{
		Ordering: order,
		CmpFunc:  factoryLessFuncOrdered[T](order),
	}
}

func NewHeapCmp[T CmpOrdered[T]](
	order data.SortOrder,
) *Heap[T] {
	return &Heap[T]{
		Ordering: order,
		CmpFunc:  factoryLessFuncCmp[T](order),
	}
}

func (h *Heap[T]) Push(item T) {
	h.Items = append(h.Items, data.NewGetValuer[T](item))
	h.heapifyUp(len(h.Items) - 1)
}

func (h *Heap[T]) PushGetter(getter data.Getter[T]) {
}

func (h *Heap[T]) Pop() data.Getter[T] {
	root := h.Items[0]
	lastIdx := len(h.Items) - 1

	h.Items[0] = h.Items[lastIdx]
	h.Items = h.Items[:lastIdx]

	h.heapifyDown(0)

	return root
}

func (h *Heap[T]) Len() int {
	return len(h.Items)
}

func (h *Heap[T]) IsEmpty() bool {
	return len(h.Items) == 0
}

func (h *Heap[T]) Peek() data.Getter[T] {
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

		if !h.CmpFunc(h.Items, curr, parent) {
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
			h.CmpFunc(h.Items, childLeft, childRight):

			child = childLeft

		default:
			child = childRight
		}

		if h.CmpFunc(h.Items, curr, child) {
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
func factoryLessFuncOrdered[T constraints.Ordered](
	order data.SortOrder,
) lessFunc[data.Getter[T]] {
	if order == MinHeap {
		return func(items []data.Getter[T], i, j int) bool {
			return items[i].GetValue() < items[j].GetValue()
		}
	}

	return func(items []data.Getter[T], i, j int) bool {
		return items[i].GetValue() > items[j].GetValue()
	}
}

func factoryLessFuncCmp[T CmpOrdered[T]](
	order data.SortOrder,
) lessFunc[data.Getter[T]] {
	if order == MinHeap {
		return func(items []data.Getter[T], i, j int) bool {
			return items[i].GetValue().Cmp(items[j].GetValue()) < 0
		}
	}

	return func(items []data.Getter[T], i, j int) bool {
		return items[i].GetValue().Cmp(items[j].GetValue()) > 0
	}
}
