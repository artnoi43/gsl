package data

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type TypePQ uint8

const (
	MinHeap TypePQ = iota
	MaxHeap
)

// ItemPQ constraints Valuer to constraints.Ordered, because we need this to compare the values
type ItemPQ[T constraints.Ordered] Valuer[T]

type PriorityQueue[T constraints.Ordered] struct {
	Items []ItemPQ[T]
	Type  TypePQ
}

func NewPriorityQueue[T constraints.Ordered](t TypePQ) *PriorityQueue[T] {
	return &PriorityQueue[T]{
		Type: t,
	}
}

func (self *PriorityQueue[T]) Len() int { return len(self.Items) }

func (self *PriorityQueue[T]) Less(i, j int) bool {
	if self.Type == MinHeap {
		return self.Items[i].Value() < self.Items[j].Value()
	}
	return self.Items[i].Value() > self.Items[j].Value()
}

func (self *PriorityQueue[T]) Swap(i, j int) {
	self.Items[i], self.Items[j] = self.Items[j], self.Items[i]
}

func (self *PriorityQueue[T]) Push(x any) {
	item, ok := x.(ItemPQ[T])
	if !ok {
		typeOfT := fmt.Sprintf("%T", new(T))
		panic(fmt.Sprintf("x is not of type %s", typeOfT))
	}
	self.Items = append(self.Items, item)
}

func (self *PriorityQueue[T]) Pop() any {
	old := *self
	n := len(old.Items)
	item := old.Items[n-1]
	old.Items[n-1] = nil // avoid memory leak
	self.Items = old.Items[0 : n-1]
	return item
}

func (self *PriorityQueue[T]) IsEmpty() bool {
	return self.Len() == 0
}
