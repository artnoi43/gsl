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
	curr := root
	for {
		switch {
		case curr == nil || !curr.ok:
			curr = node
			curr.ok = true

			return

		case node.value == curr.value:
			// Do nothing if duplicate nodes
			return

		case node.value < curr.value:
			if curr.left == nil {
				curr.left = node
				return
			}

			curr = curr.left

		case node.value > curr.value:
			if curr.right == nil {
				curr.right = node
				return
			}

			curr = curr.right
		}
	}
}

// remove removes target from subtree tree, returning the new root of the subtree
func remove[T constraints.Ordered](root *binTreeNode[T], target T) *binTreeNode[T] {
	switch {
	case root == nil:
		return nil

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

			root.ok = false
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

func BstInsertRecurse[T constraints.Ordered](root *binTreeNode[T], node *binTreeNode[T]) {
	switch {
	case !root.ok:
		root = node
		root.ok = true

		return

	case node.value == root.value:
		// Do nothing if duplicate nodes
		return

	case node.value < root.value:
		if root.left == nil {
			root.left = node
			return
		}

		BstInsertRecurse(root.left, node)

	case node.value > root.value:
		if root.right == nil {
			root.right = node
			return
		}

		BstInsertRecurse(root.right, node)
	}
}
