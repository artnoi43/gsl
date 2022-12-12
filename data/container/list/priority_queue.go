package list

import (
	"fmt"
	"sync"

	"golang.org/x/exp/constraints"

	"github.com/artnoi43/gsl/data"
)

type HeapType uint8

const (
	MinHeap HeapType = iota
	MaxHeap
)

// PriorityQueue implements heap.Interface,
// and can be use with container/heap.Push, container/heap.Pop, and container/heap.Init.
// I'm working on a new implementation that wouldn't require the heap package.
type PriorityQueue[T any] struct {
	Items    []data.GetValuer[T]
	HeapType HeapType
	// lessFunc depends on T, and the New* functions below
	lessFunc func(items []data.GetValuer[T], t HeapType, i, j int) bool
	mut      *sync.RWMutex
}

// CmpOrdered represents any type T with `Cmp(T) int` method.
// Examples of types that implement this interface include *big.Int and *big.Float.
type CmpOrdered[T any] interface {
	Cmp(T) int
}

// NewPriorityQueue returns *PriorityQueue[constraints.Ordered], and this instance
// uses lessOrdered as lessFunc, which means that the priority type must be able to compare ordering
// using greater than (>) and lesser than (<) family of signs.
func NewPriorityQueue[T constraints.Ordered](t HeapType) *PriorityQueue[T] {
	return &PriorityQueue[T]{
		HeapType: t,
		lessFunc: lessOrdered[T],
		mut:      &sync.RWMutex{},
	}
}

// NewPriorityQueueCmp[T] returns *PriorityQueue[CmpOrdered[T]], and this instance
// uses lessCmp as lessFunc, which means that the priority type must be able to compare ordering
// using Cmp(T) int function.
func NewPriorityQueueCmp[T CmpOrdered[T]](t HeapType) *PriorityQueue[T] {
	return &PriorityQueue[T]{
		HeapType: t,
		lessFunc: lessCmp[T],
		mut:      &sync.RWMutex{},
	}
}

// If the priority type for your priority queue does not implement constraints.Ordered or CmpOrdered interface,
// then you can provide your own lessFunc to determine ordering.
func NewPriorityQueueCustom[T any](
	t HeapType,
	lessFunc func(items []data.GetValuer[T], t HeapType, i, j int) bool,
) *PriorityQueue[T] {
	return &PriorityQueue[T]{
		HeapType: t,
		lessFunc: lessFunc,
		mut:      &sync.RWMutex{},
	}
}

// Less implementation for constraints.Ordered
func lessOrdered[T constraints.Ordered](items []data.GetValuer[T], t HeapType, i, j int) bool {
	if t == MinHeap {
		return items[i].GetValue() < items[j].GetValue()
	}

	return items[i].GetValue() > items[j].GetValue()
}

// Less implementation for CmpOrdered, e.g. *big.Int and *big.Float, and other lib types.
func lessCmp[T CmpOrdered[T]](items []data.GetValuer[T], t HeapType, i, j int) bool {
	var cmp int
	switch t {
	case MinHeap:
		cmp = -1
	case MaxHeap:
		cmp = 1
	}

	return items[i].GetValue().Cmp(items[j].GetValue()) == cmp
}

func (q *PriorityQueue[T]) Len() int {
	q.mut.RLock()
	defer q.mut.RUnlock()

	return len(q.Items)
}

func (q *PriorityQueue[T]) Less(i, j int) bool {
	q.mut.RLock()
	defer q.mut.RUnlock()

	return q.lessFunc(q.Items, q.HeapType, i, j)
}

func (q *PriorityQueue[T]) Swap(i, j int) {
	q.mut.Lock()
	defer q.mut.Unlock()

	q.Items[i], q.Items[j] = q.Items[j], q.Items[i]
}

func (q *PriorityQueue[T]) Push(x any) {
	q.mut.Lock()
	defer q.mut.Unlock()

	item, ok := x.(data.GetValuer[T])
	if !ok {
		typeOfT := fmt.Sprintf("%T", new(T))
		panic(fmt.Sprintf("x is not of type %s", typeOfT))
	}
	q.Items = append(q.Items, item)
}

func (q *PriorityQueue[T]) Pop() any {
	q.mut.Lock()
	defer q.mut.Unlock()

	old := *q
	n := len(old.Items)
	item := old.Items[n-1]
	old.Items[n-1] = nil // avoid memory leak
	q.Items = old.Items[0 : n-1]

	return item
}

func (q *PriorityQueue[T]) IsEmpty() bool {
	q.mut.RLock()
	defer q.mut.RUnlock()

	return q.Len() == 0
}
