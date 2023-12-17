package tree

import (
	"golang.org/x/exp/constraints"
)

// Bst is BST implementation with nodeWrapper as node
type Bst[T constraints.Ordered] struct {
	root binTreeNode[T]
}

func (b *Bst[T]) Insert(item T) {
	node := binTreeNode[T]{
		value: item,
		ok:    true,
		left:  nil,
		right: nil,
	}

	if b.root.IsNull() {
		b.root = node
		return
	}

	insert(&b.root, &node)
}

func (b *Bst[T]) Remove(node T) bool {
	return remove(&b.root, node) != nil
}

func (b *Bst[T]) Find(node T) bool {
	curr := &b.root
	for {
		if curr == nil || curr.IsNull() {
			return false
		}

		v := curr.value

		switch {
		case node == v:
			return true

		case node < v:
			curr = curr.left

		case node > v:
			curr = curr.right

		default:
			panic("unhandled case")
		}
	}
}

func insert[T constraints.Ordered](root *binTreeNode[T], node *binTreeNode[T]) {
	if !root.ok {
		root = node
		root.ok = true

		return
	}

	switch {
	case root.IsLeaf():
		switch {
		case node.value < root.value:
			root.left = node

		case node.value > root.value:
			root.right = node
		}

		return

	case root.hasBoth():
		switch {
		case node.value < root.left.value:
			insert(root.left, node)

		case node.value > root.right.value:
			insert(root.right, node)
		}

	case root.hasLeft():
		insert(root.left, node)

	case root.hasRight():
		insert(root.right, node)
	}
}

func remove[T constraints.Ordered](root *binTreeNode[T], target T) *binTreeNode[T] {
	if root == nil || !root.ok {
		return nil
	}

	switch {
	case target == root.value:
		switch {
		case root.hasBoth():
			replacement := digRight(root.left)
			return replacement

		case root.hasLeft():
			return root.left

		case root.hasRight():
			return root.right
		}

	case target < root.value:
		root.left = remove(root.left, target)

	case target > root.value:
		root.right = remove(root.right, target)
	}

	return root
}

// Returns left leaf of a tree root
func digLeft[T constraints.Ordered](root *binTreeNode[T]) *binTreeNode[T] {
	curr := root
	for curr.left != nil && curr.left.ok {
		curr = curr.left
	}

	return curr
}

// Returns right leaf of a tree root
func digRight[T constraints.Ordered](root *binTreeNode[T]) *binTreeNode[T] {
	curr := root
	for curr.right != nil && !curr.right.ok {
		curr = curr.right
	}

	return curr
}
