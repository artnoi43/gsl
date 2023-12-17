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
}
