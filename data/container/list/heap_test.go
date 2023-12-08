package list

import (
	"fmt"
	"testing"

	"github.com/soyart/gsl/data"
)

func TestHeapifyUp(t *testing.T) {
	ints := []int{6, 1, 3, 3, 2, 4, 5}
	pq := NewHeap[int](MaxHeap)

	for i := range ints {
		pq.Push(ints[i])
	}

	pq.Push(3)
	pq.Push(0)
	pq.Push(9)

	l := len(pq.Items)
	for curr := range pq.Items {
		fmt.Println("index", curr, "value", pq.Items[curr].GetValue())

		parent := parentNode(curr)
		if parent >= 0 {
			if pq.lessFunc(pq.Items, MaxHeap, curr, parent) {
				valueCurr := pq.Items[curr].GetValue()
				valueParent := pq.Items[parent].GetValue()

				t.Logf(
					"node at %d is less than parent at %d: node %v vs parent %v",
					curr, parent, valueCurr, valueParent,
				)

				t.Fatalf("Unexpected value")
			}
		}

		childLeft := leftChildNode(curr)
		if childLeft >= l {
			continue
		}

		if !pq.lessFunc(pq.Items, MaxHeap, curr, childLeft) {
			valueCurr := pq.Items[curr].GetValue()
			valueLeft := pq.Items[childLeft].GetValue()

			t.Logf(
				"node at %d is less than left child at %d: node %v vs parent %v",
				curr, childLeft, valueCurr, valueLeft,
			)

			t.Fatalf("Unexpected value")
		}

		childRight := childLeft + 1
		if childRight >= l {
			continue
		}

		if !pq.lessFunc(pq.Items, MaxHeap, curr, childRight) {
			valueCurr := pq.Items[curr].GetValue()
			valueRight := pq.Items[childRight].GetValue()

			t.Logf(
				"node at %d is less than left child at %d: node %v vs parent %v",
				curr, childRight, valueCurr, valueRight,
			)

			t.Fatalf("Unexpected heap shape")
		}
	}
}

func TestHeapifyDown(t *testing.T) {
	ints := []int{6, 1, 3, 3, 2, 4, 5}
	pq := NewHeap[int](MaxHeap)

	for i := range ints {
		pq.Push(ints[i])
	}

	intsHeap := make([]int, len(pq.Items))
	for i := range pq.Items {
		intsHeap[i] = pq.Items[i].GetValue()
	}

	t.Log("intsHeap", intsHeap)

	var c int
	for l := len(pq.Items); l > 0; l = len(pq.Items) {
		max := data.MaxValuer(pq.Items)
		popped := pq.Pop().GetValue()

		items := make([]int, l)
		for i := range pq.Items {
			items[i] = pq.Items[i].GetValue()
		}

		t.Log("l", l, "max", max, "pop", popped, "items", items)

		if max != popped {
			t.Logf("pop #%d expecting max %d, got %d", c+1, max, popped)
			t.Fatal("Unexpected popped value")
		}

		c++
	}
}
