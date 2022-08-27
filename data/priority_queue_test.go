package data

import (
	"container/heap"
	"reflect"
	"testing"

	"golang.org/x/exp/constraints"
)

type foo[T constraints.Ordered] struct {
	name  string
	value T
}

func (self foo[T]) Value() T {
	return self.value
}

func TestPq(t *testing.T) {
	highest := foo[int]{name: "b", value: 100}
	lowest := foo[int]{name: "d", value: 0}

	items := []foo[int]{
		{name: "a", value: 69},
		highest,
		{name: "c", value: 12},
		lowest,
	}

	outMinHeap := testPop(t, MinHeap, items)
	if outMinHeap != lowest {
		t.Fatalf("unexpected MinHeap results - expected %+v, got %+v\n", lowest, outMinHeap)
	}
	outMaxHeap := testPop(t, MaxHeap, items)
	if outMaxHeap != highest {
		t.Fatalf("unexpected MaxHeap results - expected %+v, got %+v\n", highest, outMaxHeap)
	}

	testArbitaryUpdate(t)
}

func testPop[T constraints.Ordered](t *testing.T, pqType TypePQ, items []foo[T]) foo[T] {
	pq := NewPriorityQueue[T](pqType)
	for _, item := range items {
		heap.Push(pq, item)
	}

	p := heap.Pop(pq)
	popped, ok := p.(foo[T])
	if !ok {
		t.Fatalf("type assertion to *foo failed, got type: %s", reflect.TypeOf(p))
	}
	return popped
}

func testArbitaryUpdate(t *testing.T) {
	// Test with Valuer[float64]
	hundred := foo[float64]{name: "hundred", value: 100}
	seventy := foo[float64]{name: "seventy", value: 70}
	zero := foo[float64]{name: "zero", value: 0}
	foosFloat := []foo[float64]{
		{name: "a", value: 69},
		hundred,
		{name: "b", value: 71},
		zero,
		seventy,
	}

	// Arbitary pushes and inits
	pq := NewPriorityQueue[float64](MaxHeap)
	for _, item := range foosFloat {
		heap.Push(pq, item)
	}

	p := heap.Pop(pq)
	popped, ok := p.(foo[float64])
	if !ok {
		t.Fatalf("type assertion to *foo failed, got type: %s", reflect.TypeOf(p))
	}
	if popped != hundred {
		t.Fatalf("unexpected MaxHeap results - expected %+v, got %+v\n", hundred, popped)
	}

	pq.Type = MinHeap
	heap.Init(pq)
	p = heap.Pop(pq)
	popped, ok = p.(foo[float64])
	if !ok {
		t.Fatalf("type assertion to *foo failed, got type: %s", reflect.TypeOf(p))
	}
	if popped != zero {
		t.Fatalf("unexpected MinHeap results - expected %+v, got %+v\n", zero, popped)
	}
}
