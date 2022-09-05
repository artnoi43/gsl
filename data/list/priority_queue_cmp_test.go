package list

import (
	"container/heap"
	"math/big"
	"testing"

	"github.com/artnoi43/mgl/data"
)

type bar struct {
	val *big.Int
}

func (b *bar) GetValue() *big.Int {
	return b.val
}

func TestPQCmp(t *testing.T) {
	a := &bar{val: big.NewInt(69)}
	b := &bar{val: big.NewInt(70)}
	c := &bar{val: big.NewInt(100)}
	d := &bar{val: big.NewInt(1000000)}

	t.Run("MaxHeap with Cmp", func(t *testing.T) {
		testPqCmpMax(t, []*bar{a, d, c, b}, d)
	})
	t.Run("MinHeap with Cmp", func(t *testing.T) {
		testPqCmpMin(t, []*bar{a, d, c, b}, a)
	})

	lol(a) // Compiles and no panic
}

func lol(item data.Valuer[*big.Int]) {

}

func testPqCmpMax(t *testing.T, messy []*bar, max *bar) {
	maxPq := NewPriorityQueueCmp[*big.Int](MaxHeap)
	for _, item := range messy {
		heap.Push(maxPq, item)
	}
	if popped := heap.Pop(maxPq); popped != nil {
		actual := popped.(*bar)
		if actual != max {
			t.Fatalf("unexpected max heap result: expected %v, got %v\n", max.GetValue(), actual.GetValue())
		}
	}
}

func testPqCmpMin(t *testing.T, messy []*bar, min *bar) {
	minPq := NewPriorityQueueCmp[*big.Int](MinHeap)
	for _, item := range messy {
		heap.Push(minPq, item)
	}
	if popped := heap.Pop(minPq); popped != nil {
		actual := popped.(*bar)
		if actual != min {
			t.Fatalf("unexpected min heap result: expected %v, got %v\n", min.GetValue(), actual.GetValue())
		}
	}
}
