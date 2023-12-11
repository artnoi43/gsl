package list

import (
	"container/heap"
	"fmt"
	"sync"

	"golang.org/x/exp/constraints"

	"github.com/soyart/gsl/data"
)

// GoHeapImpl implements heap.Interface,
// and can be use with container/heap.Push, container/heap.Pop, and container/heap.Init.
//
// Ordering can be changed arbitrarily, but users must explicitly calls container/heap.Init.
type GoHeapImpl[T any] struct {
	Items   []data.Getter[T]
	CmpFunc data.LessFunc[data.Getter[T]]

	ordering data.SortOrder
	mut      *sync.RWMutex
}

// NewPriorityQueue returns *PriorityQueue[constraints.Ordered], and this instance
// uses lessFuncOrdered as lessFunc, which means that the priority type must be able to compare ordering
// using greater than (>) and lesser than (<) family of signs.
func NewPriorityQueue[T constraints.Ordered](order data.SortOrder) *GoHeapImpl[T] {
	return &GoHeapImpl[T]{
		CmpFunc:  data.FactoryLessFuncOrdered[T](order),
		ordering: order,
		mut:      new(sync.RWMutex),
	}
}

// NewPrioirtyQueueCmp[T] returns *PriorityQueue[CmpOrdered[T]] with the default lessFunc
// for type T with Cmp(T), i.e. -1 -> less, 0 -> equal, 1 -> greater.
func NewPrioirtyQueueCmp[T data.CmpOrdered[T]](order data.SortOrder) *GoHeapImpl[T] {
	return &GoHeapImpl[T]{
		CmpFunc:  data.FactoryLessFuncCmp[T](order),
		ordering: order,
		mut:      new(sync.RWMutex),
	}
}

// NewPriorityQueueCustom use the provided lessFunc for the heapify processes.
func NewPriorityQueueCustom[T any](
	order data.SortOrder,
	lessFunc data.LessFunc[data.Getter[T]],
) *GoHeapImpl[T] {
	return &GoHeapImpl[T]{
		CmpFunc:  lessFunc,
		ordering: order,
		mut:      new(sync.RWMutex),
	}
}

func (q *GoHeapImpl[T]) ChangeOrdering(lessFunc data.LessFunc[data.Getter[T]]) {
	q.CmpFunc = lessFunc
	heap.Init(q)
}

func (q *GoHeapImpl[T]) Len() int {
	q.mut.RLock()
	defer q.mut.RUnlock()

	return len(q.Items)
}

func (q *GoHeapImpl[T]) Less(i, j int) bool {
	q.mut.RLock()
	defer q.mut.RUnlock()

	return q.CmpFunc(q.Items, i, j)
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
