package tree

import (
	"testing"
)

func TestBstInsertFind(t *testing.T) {
	bst := new(Bst[int])

	limit := 10
	for i := 0; i < limit; i++ {
		bst.Insert(i)
	}

	for i := 0; i < limit; i++ {
		if !bst.Find(i) {
			t.Fatalf("missing node %d", i)
		}
	}

	outOfRange := -1
	if bst.Find(outOfRange) {
		t.Fatalf("unexpected false positive for %d", outOfRange)
	}
}

func TestBstRemove(t *testing.T) {
	bst := new(Bst[int])

	limit := 10
	target := 5
	for i := 0; i < limit; i++ {
		bst.Insert(i)
	}

	if !bst.Remove(target) {
		t.Fatalf("remove returned false on target %d", target)
	}

	if bst.Find(target) {
		t.Fatalf("found removed target %d", target)
	}

	for i := 0; i < limit; i++ {
		if !bst.Find(i) {
			if i == target {
				continue
			}

			t.Fatalf("missing node %d", i)
		}
	}
}

func TestBst(t *testing.T) {
	bst := new(Bst[int])

	targetRemove := 9

	items := []int{2, targetRemove, -1, 1, 7, 5, 30, 8, 3}

	for i := range items {
		bst.Insert(items[i])
	}

	for i := range items {
		item := items[i]

		if !bst.Find(items[i]) {
			t.Fatalf("missing item %d", item)
		}
	}

	if !bst.Remove(targetRemove) {
		t.Fatalf("found removed target %d", targetRemove)
	}

	for i := range items {
		item := items[i]
		if !bst.Find(item) {
			if item == targetRemove {
				continue
			}

			t.Fatalf("missing node %d", item)
		}
	}
}
