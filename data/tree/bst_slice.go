package tree

import (
	"golang.org/x/exp/constraints"
)

// Binary search tree implementation with backing array
type BstSlice[T constraints.Ordered] struct {
	backing []T
	dups    map[T]struct{}
}

func NewBstSlice[T constraints.Ordered]() BinaryTree[int, T] {
	return &BstSlice[T]{}
}

func (b *BstSlice[T]) Insert(node T) {
	if b.allowDups() {
		panic("not implemented duplicate nodes")
	}

	if _, duplicate := b.dups[node]; duplicate {
		return
	}

	b.insert(0, node)
}

func (b *BstSlice[T]) Remove(node T) bool {
	targetPos := b.find(node)
	if targetPos == -1 {
		return false
	}

	panic("not implemented")
}

func (b *BstSlice[T]) Find(target T) bool {
	if len(b.backing) == 0 {
		return false
	}

	curr := 0
	for {
		if b.NodeIsNull(curr) {
			return false
		}

		node := b.Node(curr)
		switch {
		case target == node:
			return true

		case target > node:
			curr = b.RightChild(curr)

		case target < node:
			curr = b.LeftChild(curr)

		default:
			panic("unhandled case")
		}
	}
}

func (b *BstSlice[T]) Parent(pos int) int {
	return ParentIdx(pos)
}

func (b *BstSlice[T]) LeftChild(pos int) int {
	return LeftChildIdx(pos)
}

func (b *BstSlice[T]) RightChild(pos int) int {
	return RightChildIdx(pos)
}

func (b *BstSlice[T]) Node(pos int) T {
	return b.backing[pos]
}

func (b *BstSlice[T]) NodeIsRoot(pos int) bool {
	return pos == 0
}

func (b *BstSlice[T]) NodeIsNull(pos int) bool {
	return pos >= len(b.backing)
}

func (b *BstSlice[T]) insert(root int, node T) {
	if b.NodeIsNull(root) {
		b.addToBacking(root, node)
		return
	}

	left := b.LeftChild(root)
	right := b.RightChild(root)
	nextParent := -1

	rootIsLeaf := b.NodeIsNull(left) && b.NodeIsNull(right)
	rootHasBoth := !rootIsLeaf
	rootHasLeft := !b.NodeIsNull(left)
	rootHasRight := !b.NodeIsNull(right)

	switch {
	case rootIsLeaf, rootHasBoth:
		insertIdx := -1

		parentNode := b.Node(root)
		switch {
		case node < parentNode:
			insertIdx = left

		case node > parentNode:
			insertIdx = right

		default:
			panic("unexpected case node == parent")
		}

		if insertIdx == -1 {
			panic("unexpected -1 insertIdx")
		}

		if rootHasBoth {
			b.insert(insertIdx, node)
			return
		}

		b.addToBacking(insertIdx, node)

	case rootHasLeft:
		leftNode := b.Node(left)
		if node < leftNode {
			b.addToBacking(left, node)
			return
		}

		nextParent = left

	case rootHasRight:
		rightNode := b.Node(right)
		if node > rightNode {
			b.addToBacking(right, node)
			return
		}

		nextParent = right
	}

	b.insert(nextParent, node)
}

func (b *BstSlice[T]) addToBacking(insertIdx int, node T) {
	l := len(b.backing)
	if insertIdx >= l {
		newBacking := make([]T, l)
		copy(newBacking, b.backing)

		b.backing = newBacking
	}

	b.backing[insertIdx] = node
}

func (b *BstSlice[T]) find(target T) int {
	for i := range b.backing {
		if b.backing[i] == target {
			return i
		}
	}

	return -1
}

func (b *BstSlice[T]) allowDups() bool {
	return b.dups == nil
}

func (b *BstSlice[T]) isDuplicate(node T) bool {
	_, duplicate := b.dups[node]

	return duplicate
}
