package tree

import "golang.org/x/exp/constraints"

// Bst is BST implementation with nodeWrapper as node
type Bst[T constraints.Ordered] struct {
	count int
	root  nodeWrapper[T]
}

func (b *Bst[T]) Insert(node T) {
	insert[T](&b.root, &nodeWrapper[T]{
		value: node,
		ok:    false,
		left:  nil,
		right: nil,
	})
}

func (b *Bst[T]) Remove(node T) {
	panic("not implemented")
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

func insert[T constraints.Ordered](root *nodeWrapper[T], node *nodeWrapper[T]) {
	if root.ok {
		root = node
		return
	}

	left := root.left
	right := root.right

	isLeaf := left.IsNull() && right.IsNull()
	hasBoth := !isLeaf
	hasLeft := !left.IsNull()
	hasRight := !right.IsNull()

	var nextParent *nodeWrapper[T]

	switch {
	case isLeaf:
		switch {
		case node.value < root.value:
			root.left = node

		case node.value > root.value:
			root.right = node
		}

	case hasBoth:
		switch {
		case node.value < left.value:
			nextParent = left

		case node.value > right.value:
			nextParent = right
		}

	case hasLeft:
		nextParent = left

	case hasRight:
		nextParent = right
	}

	if nextParent == nil {
		panic("unexpected nil next parent")
	}

	insert(nextParent, node)
}
