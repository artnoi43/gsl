package list

import (
	"fmt"
	"sync"

	"github.com/artnoi43/mgl/data"
	"golang.org/x/exp/constraints"
)

type TypePQ uint8

const (
	MinHeap TypePQ = iota
	MaxHeap
)

// ItemPQ constraints Valuer to constraints.Ordered, because we need this to compare the values
type ItemPQ[T constraints.Ordered] data.Valuer[T]

type PriorityQueue[T constraints.Ordered] struct {
	Items []ItemPQ[T]
	Type  TypePQ
	mut   *sync.RWMutex
}

func NewPriorityQueue[T constraints.Ordered](t TypePQ) *PriorityQueue[T] {
	return &PriorityQueue[T]{
		Type: t,
		mut:  &sync.RWMutex{},
	}
}

func (self *PriorityQueue[T]) Len() int {
	self.mut.RLock()
	defer self.mut.RUnlock()

	return len(self.Items)
}

func (self *PriorityQueue[T]) Less(i, j int) bool {
	if self.Type == MinHeap {
		return self.Items[i].GetValue() < self.Items[j].GetValue()
	}

	return self.Items[i].GetValue() > self.Items[j].GetValue()
}

func (self *PriorityQueue[T]) Swap(i, j int) {
	self.mut.Lock()
	defer self.mut.Unlock()

	self.Items[i], self.Items[j] = self.Items[j], self.Items[i]
}

func (self *PriorityQueue[T]) Push(x any) {
	self.mut.Lock()
	defer self.mut.Unlock()

	item, ok := x.(ItemPQ[T])
	if !ok {
		typeOfT := fmt.Sprintf("%T", new(T))
		panic(fmt.Sprintf("x is not of type %s", typeOfT))
	}
	self.Items = append(self.Items, item)
}

func (self *PriorityQueue[T]) Pop() any {
	self.mut.Lock()
	defer self.mut.Unlock()

	old := *self
	n := len(old.Items)
	item := old.Items[n-1]
	old.Items[n-1] = nil // avoid memory leak
	self.Items = old.Items[0 : n-1]

	return item
}

func (self *PriorityQueue[T]) IsEmpty() bool {
	self.mut.RLock()
	defer self.mut.RUnlock()

	return self.Len() == 0
}
