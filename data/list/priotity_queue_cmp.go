package list

import (
	"fmt"
	"reflect"
	"sync"

	"github.com/artnoi43/mgl/data"
)

// CmpOrdered represents any type T with `Cmp(T) int` method.
// Examples of types that implement this interface include *big.Int and *big.Float.
type CmpOrdered[T any] interface {
	Cmp(T) int
}

// An ItemPQCmp is like a structure whose GetValue() method returns CmpOrdered type.
// And example would be `struct foo { value *big.Int }`, with method `GetValue() *big.Int`
type ItemPQCmp[T any] data.Valuer[T]

type PriorityQueueCmp[T CmpOrdered[T]] struct {
	Items []ItemPQCmp[T]
	Type  TypePQ
	mut   *sync.RWMutex
}

func NewPriorityQueueCmp[T CmpOrdered[T]](t TypePQ) *PriorityQueueCmp[T] {
	return &PriorityQueueCmp[T]{
		Type: t,
		mut:  &sync.RWMutex{},
	}
}

func (pq *PriorityQueueCmp[T]) Len() int {
	pq.mut.RLock()
	defer pq.mut.RUnlock()

	return len(pq.Items)
}

func (pq *PriorityQueueCmp[T]) Less(i, j int) bool {
	pq.mut.RLock()
	defer pq.mut.RUnlock()

	var cmp int
	switch pq.Type {
	case MinHeap:
		cmp = -1
	case MaxHeap:
		cmp = 1
	}

	return pq.Items[i].GetValue().Cmp(pq.Items[j].GetValue()) == cmp
}

func (pq *PriorityQueueCmp[T]) Swap(i, j int) {
	pq.mut.Lock()
	defer pq.mut.Unlock()

	pq.Items[i], pq.Items[j] = pq.Items[j], pq.Items[i]
}

func (pq *PriorityQueueCmp[T]) Push(x any) {
	pq.mut.Lock()
	defer pq.mut.Unlock()
	item, ok := x.(ItemPQCmp[T])
	if !ok {
		typeOfT := reflect.TypeOf(pq.Items).String()[2:]
		panic(fmt.Sprintf("x is not of type %s", typeOfT))
	}
	pq.Items = append(pq.Items, item)
}

func (pq *PriorityQueueCmp[T]) Pop() any {
	pq.mut.Lock()
	defer pq.mut.Unlock()

	old := *pq
	n := len(old.Items)
	item := old.Items[n-1]
	old.Items[n-1] = nil // avoid memory leak
	pq.Items = old.Items[0 : n-1]

	return item
}

func (pq *PriorityQueueCmp[T]) IsEmpty() bool {
	pq.mut.RLock()
	defer pq.mut.RUnlock()

	return pq.Len() == 0
}
