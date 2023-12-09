package list

import (
	"fmt"
	"sync"

	"golang.org/x/exp/constraints"

	"github.com/soyart/gsl/data"
)

const (
	minHeap = data.Ascending
	maxHeap = data.Descending
)

// GoHeapImpl implements heap.Interface,
// and can be use with container/heap.Push, container/heap.Pop, and container/heap.Init.
//
// Ordering can be changed arbitrarily, but users must explicitly calls container/heap.Init.
type GoHeapImpl[T any] struct {
	Items    []data.Getter[T]
	Ordering data.SortOrder

	// CmpFunc depends on T, and the New* functions below
	CmpFunc func(items []data.Getter[T], order data.SortOrder, i, j int) bool
	mut     *sync.RWMutex
}

// NewHeapImpl returns *PriorityQueue[constraints.Ordered], and this instance
// uses lessFuncOrdered as lessFunc, which means that the priority type must be able to compare ordering
// using greater than (>) and lesser than (<) family of signs.
func NewHeapImpl[T constraints.Ordered](order data.SortOrder) *GoHeapImpl[T] {
	return &GoHeapImpl[T]{
		Ordering: order,
		CmpFunc:  lessFuncOrdered[T],
		mut:      new(sync.RWMutex),
	}
}

// NewHealImplCmp[T] returns *PriorityQueue[CmpOrdered[T]], and this instance
// uses lessFuncCmp as lessFunc, which means that the priority type must be able to compare ordering
// using Cmp(T) int function.
func NewHealImplCmp[T data.CmpOrdered[T]](order data.SortOrder) *GoHeapImpl[T] {
	return &GoHeapImpl[T]{
		Ordering: order,
		CmpFunc:  lessFuncCmp[T],
		mut:      new(sync.RWMutex),
	}
}

// If the priority type for your priority queue does not implement constraints.Ordered or CmpOrdered interface,
// then you can provide your own lessFunc to determine ordering.
func NewHeapImplCustom[T any](
	order data.SortOrder,
	lessFunc func(items []data.Getter[T], order data.SortOrder, i, j int) bool,
) *GoHeapImpl[T] {
	return &GoHeapImpl[T]{
		Ordering: order,
		CmpFunc:  lessFunc,
		mut:      new(sync.RWMutex),
	}
}

func (q *GoHeapImpl[T]) Len() int {
	q.mut.RLock()
	defer q.mut.RUnlock()

	return len(q.Items)
}

func (q *GoHeapImpl[T]) Less(i, j int) bool {
	q.mut.RLock()
	defer q.mut.RUnlock()

	return q.CmpFunc(q.Items, q.Ordering, i, j)
}

func (q *GoHeapImpl[T]) Swap(i, j int) {
	q.mut.Lock()
	defer q.mut.Unlock()

	q.Items[i], q.Items[j] = q.Items[j], q.Items[i]
}

func (q *GoHeapImpl[T]) Push(x any) {
	q.mut.Lock()
	defer q.mut.Unlock()

	item, ok := x.(data.Getter[T])
	if !ok {
		typeOfT := fmt.Sprintf("%T", new(T))
		panic(fmt.Sprintf("x is not of type %s", typeOfT))
	}

	q.Items = append(q.Items, item)
}

func (q *GoHeapImpl[T]) Pop() any {
	q.mut.Lock()
	defer q.mut.Unlock()

	old := *q
	n := len(old.Items)
	item := old.Items[n-1]
	old.Items[n-1] = nil // avoid memory leak
	q.Items = old.Items[0 : n-1]

	return item
}

func (q *GoHeapImpl[T]) IsEmpty() bool {
	q.mut.RLock()
	defer q.mut.RUnlock()

	return q.Len() == 0
}

// Less implementation for constraints.Ordered
func lessFuncOrdered[T constraints.Ordered](
	items []data.Getter[T],
	order data.SortOrder,
	i int,
	j int,
) bool {
	if order == minHeap {
		return items[i].GetValue() < items[j].GetValue()
	}

	return items[i].GetValue() > items[j].GetValue()
}

// Less implementation for CmpOrdered, e.g. *big.Int and *big.Float, and other lib types.
func lessFuncCmp[T data.CmpOrdered[T]](
	items []data.Getter[T],
	order data.SortOrder,
	i int,
	j int,
) bool {
	var cmp int

	switch order {
	case minHeap:
		cmp = -1

	case maxHeap:
		cmp = 1
	}

	return items[i].GetValue().Cmp(items[j].GetValue()) == cmp
}
