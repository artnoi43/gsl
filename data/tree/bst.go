package tree

import (
	"fmt"

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

func (b *Bst[T]) Find(target T) bool {
	curr := &b.root
	for {
		if curr == nil || curr.IsNull() {
			return false
		}

		v := curr.value

		switch {
		case target == v:
			return true

		case target < v:
			curr = curr.left

		case target > v:
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
	case node.value == root.value:
		// Do nothing if duplicate nodes
		return

	case node.value < root.value:
		if root.left == nil {
			root.left = node
			return
		}

		insert(root.left, node)

	case node.value > root.value:
		if root.right == nil {
			root.right = node
			return
		}

		insert(root.right, node)
	}
}

func remove[T constraints.Ordered](root *binTreeNode[T], target T) *binTreeNode[T] {
	if root == nil || !root.ok {
		return nil
	}

	switch {
	case target < root.value:
		root.left = remove(root.left, target)

	case target > root.value:
		root.right = remove(root.right, target)

	// Do the actual removal
	default:
		switch {
		case root.left == nil:
			return root.right

		case root.right == nil:
			return root.left

		default:
			replacement := digLeft(root.right)
			root.value = replacement.value

			root.right = remove(root.right, replacement.value)
		}
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
