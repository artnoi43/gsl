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
	return remove(&b.root, node)
}

func (b *Bst[T]) Find(node T) bool {
	curr := &b.root
	for {

		if curr == nil {
			return false
		}

		if curr.IsNull() {
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

	rootIsLeaf := root.IsLeaf()
	rootHasBoth := root.HasBoth()
	rootHasLeft := root.left != nil
	rootHasRight := root.right != nil

	switch {
	case rootIsLeaf:
		switch {
		case node.value < root.value:
			root.left = node

		case node.value > root.value:
			root.right = node
		}

		return

	case rootHasBoth:
		switch {
		case node.value < root.left.value:
			insert(root.left, node)

		case node.value > root.right.value:
			insert(root.right, node)
		}

	case rootHasLeft:
		insert(root.left, node)

	case rootHasRight:
		insert(root.right, node)
	}
}

func remove[T constraints.Ordered](node *binTreeNode[T], target T) bool {
	if !node.ok {
		return false
	}

	left, right := node.left, node.right

	nodeIsLeaf := !left.ok && !right.ok
	nodeHasBoth := !nodeIsLeaf
	nodeHasLeft := left.ok
	nodeHasRight := right.ok

	var nextRoot *binTreeNode[T]

	switch {
	case node.value == target:
		switch {
		case nodeIsLeaf:
			// TODO: do hard deletes
			node.ok = false

		case nodeHasBoth:
			replacement := digRight[T](left)
			node.value = replacement.value

			nextRoot = left

		case nodeHasLeft:
			replacement := digRight[T](left)
			node.value = replacement.value

			nextRoot = left

		case nodeHasRight:
			replacement := digLeft[T](right)
			node.value = replacement.value

			nextRoot = right

		}

	case node.value < target:
		if nodeHasBoth || nodeHasLeft {
			nextRoot = left
		}

	case node.value < target:
		if nodeHasBoth || nodeHasRight {
			nextRoot = right
		}

	}

	if nextRoot == nil {
		return false
	}

	return remove(nextRoot, target)
}

// Returns left leaf of a tree root
func digLeft[T constraints.Ordered](root *binTreeNode[T]) *binTreeNode[T] {
	curr := root
	for !curr.left.ok {
		curr = curr.left
	}

	return curr
}

// Returns right leaf of a tree root
func digRight[T constraints.Ordered](root *binTreeNode[T]) *binTreeNode[T] {
	curr := root
	for !curr.right.ok {
		curr = curr.right
	}

	return curr
}
