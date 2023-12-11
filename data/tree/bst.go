package tree

import "golang.org/x/exp/constraints"

// Bst is BST implementation with nodeWrapper as node
type Bst[T constraints.Ordered] struct {
	count int
	root  binTreeNode[T]
}

func (b *Bst[T]) Insert(node T) {
	insert[T](&b.root, &binTreeNode[T]{
		value: node,
		ok:    false,
		left:  nil,
		right: nil,
	})
}

func (b *Bst[T]) Remove(node T) bool {
	return remove(&b.root, node)
}

func (b *Bst[T]) Find(node T) bool {
	curr := &b.root
	for {
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

	left := root.left
	right := root.right

	rootIsLeaf := !left.ok && !right.ok
	rootHasBoth := !rootIsLeaf
	rootHasLeft := left.ok
	rootHasRight := right.ok

	var nextParent *binTreeNode[T]

	switch {
	case rootIsLeaf:
		switch {
		case node.value < root.value:
			root.left = node

		case node.value > root.value:
			root.right = node
		}

	case rootHasBoth:
		switch {
		case node.value < left.value:
			nextParent = left

		case node.value > right.value:
			nextParent = right
		}

	case rootHasLeft:
		nextParent = left

	case rootHasRight:
		nextParent = right
	}

	if nextParent == nil {
		panic("unexpected nil next parent")
	}

	insert(nextParent, node)
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
			replacement := findLeafRight[T](left)
			node.value = replacement.value

			nextRoot = left

		case nodeHasLeft:
			replacement := findLeafRight[T](left)
			node.value = replacement.value

			nextRoot = left

		case nodeHasRight:
			replacement := findLeafLeft[T](right)
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

func findLeafLeft[T constraints.Ordered](root *binTreeNode[T]) *binTreeNode[T] {
	curr := root
	for !curr.left.ok {
		curr = curr.left
	}

	return curr
}

func findLeafRight[T constraints.Ordered](root *binTreeNode[T]) *binTreeNode[T] {
	curr := root
	for !curr.right.ok {
		curr = curr.right
	}

	return curr
}
