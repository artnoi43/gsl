package list

import (
	"container/heap"
	"fmt"

	"golang.org/x/exp/constraints"

	"github.com/soyart/gsl/data"
)

// PriorityQueue[T] wraps items []T and implements Go's heap.Interface.
// tree.Heap is an alternative implementation with gsl's implementation of heaps.
type PriorityQueue[T any] struct {
	Items    []data.Getter[T]
	LessFunc data.LessFunc[data.Getter[T]]
}

func NewPriorityQueue[T constraints.Ordered](order data.SortOrder) *PriorityQueue[T] {
	return &PriorityQueue[T]{
		LessFunc: data.FactoryLessFuncOrdered[T](order),
	}
}

// NewPrioirtyQueueCmp[T] returns *PriorityQueue[data.CmpOrdered[T]] with the default lessFunc
// for type T with Cmp(T), i.e. -1 -> less, 0 -> equal, 1 -> greater.
func NewPrioirtyQueueCmp[T data.CmpOrdered[T]](order data.SortOrder) *PriorityQueue[T] {
	return &PriorityQueue[T]{
		LessFunc: data.FactoryLessFuncCmp[T](order),
	}
}

// NewPriorityQueueCustom use the provided lessFunc for the heapify processes.
func NewPriorityQueueCustom[T any](
	order data.SortOrder,
	lessFunc data.LessFunc[data.Getter[T]],
) *PriorityQueue[T] {
	return &PriorityQueue[T]{
		LessFunc: lessFunc,
	}
}

func (q *PriorityQueue[T]) ChangeOrdering(lessFunc data.LessFunc[data.Getter[T]]) {
	q.LessFunc = lessFunc
	heap.Init(q)
}

func (q *PriorityQueue[T]) Len() int {
	return len(q.Items)
}

func (q *PriorityQueue[T]) Less(i, j int) bool {
	return q.LessFunc(q.Items, i, j)
}

func (q *PriorityQueue[T]) Swap(i, j int) {
	q.Items[i], q.Items[j] = q.Items[j], q.Items[i]
}

func (q *PriorityQueue[T]) Push(x any) {
	item, ok := x.(data.Getter[T])
	if !ok {
		typeOfT := fmt.Sprintf("%T", new(T))
		panic(fmt.Sprintf("x is not of type %s", typeOfT))
	}

	q.Items = append(q.Items, item)
}

func (q *PriorityQueue[T]) Pop() any {
	old := *q
	n := len(old.Items)
	item := old.Items[n-1]
	old.Items[n-1] = nil // avoid memory leak
	q.Items = old.Items[0 : n-1]

	return item
}

func (q *PriorityQueue[T]) IsEmpty() bool {
	return q.Len() == 0
}
