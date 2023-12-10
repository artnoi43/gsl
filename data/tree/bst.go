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

func (b *BstSlice[T]) allowDups() bool {
	return b.dups == nil
}

func (b *BstSlice[T]) isDuplicate(node T) bool {
	_, duplicate := b.dups[node]

	return duplicate
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

func (b *BstSlice[T]) Remove(node T) {
	targetPos := b.find(node)
	if targetPos == -1 {
		return
	}

	panic("not implemented")
}

func (b *BstSlice[T]) Find(node T) bool {
	return b.find(node) >= 0
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
	left := b.LeftChild(root)
	right := b.RightChild(root)
	nextParent := -1

	switch {
	case b.NodeIsNull(left):
		leftNode := b.Node(left)
		if node < leftNode {
			b.addToBacking(left, node)
			return
		}

		nextParent = left

	case b.NodeIsNull(right):
		rightNode := b.Node(right)
		if node > rightNode {
			b.addToBacking(right, node)
			return
		}

		nextParent = right

	case b.NodeIsNull(left) && b.NodeIsNull(right):
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

		b.addToBacking(insertIdx, node)
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

func LeftChildIdx(parent int) int {
	return (2 * parent) + 1
}

func RightChildIdx(parent int) int {
	return LeftChildIdx(parent) + 1
}

func ParentIdx(child int) int {
	rightChild := child%2 == 0
	if rightChild {
		return (child - 2) / 2
	}

	return (child - 1) / 2
}
