package tree

import (
	"github.com/soyart/gsl/data"
)

type BstCustom[T data.CmpOrdered[T]] struct {
	Root binTreeNode[T]
}

func NewBstCustom[T data.CmpOrdered[T]]() *BstCustom[T] {
	return &BstCustom[T]{}
}

func (b *BstCustom[T]) Insert(item T) bool {
	node := binTreeNode[T]{
		value: item,
		ok:    true,
		left:  nil,
		right: nil,
	}

	if !b.Root.ok {
		b.Root = node
		return true
	}

	return BstCustomInsert(
		&b.Root,
		&node,
	)
}

func (b *BstCustom[T]) Find(target T) bool {
	return BstCustomFind[T](&b.Root, target)
}

func (b *BstCustom[T]) Remove(target T) bool {
	return BstCustomRemove(&b.Root, target) != nil
}

func BstCustomInsert[T data.CmpOrdered[T]](root, node *binTreeNode[T]) bool {
	curr := root

	for {
		if curr == nil {
			curr = node

			return true
		}

		cmp := node.value.Cmp(curr.value)

		switch {
		case cmp == 0:
			exists := curr.ok
			node.ok = true
			curr = node

			return !exists

		case cmp < 0:
			if curr.left == nil {
				curr.left = node
				return true
			}

			curr = curr.left

		case cmp > 0:
			if curr.right == nil {
				curr.right = node
				return true
			}

			curr = curr.right
		}
	}
}

func BstCustomFind[T data.CmpOrdered[T]](root *binTreeNode[T], target T) bool {
	curr := root

	for {
		if curr == nil || curr.IsNull() {
			return false
		}

		cmp := target.Cmp(curr.value)

		switch {
		case cmp == 0:
			return true

		case cmp < 0:
			curr = curr.left

		case cmp > 0:
			curr = curr.right

		default:
			panic("unhandled case")
		}
	}
}

// BstRemove removes target from subtree tree, returning the new root of the subtree
func BstCustomRemove[T data.CmpOrdered[T]](root *binTreeNode[T], target T) *binTreeNode[T] {
	if root == nil {
		return nil
	}

	cmp := target.Cmp(root.value)

	switch {
	case cmp < 0:
		root.left = BstCustomRemove(root.left, target)

	case cmp > 0:
		root.right = BstCustomRemove(root.right, target)

	// Do the actual removal
	case cmp == 0:
		switch {
		case root.left == nil:
			return root.right

		case root.right == nil:
			return root.left

		default:
			replacement := digLeft(root.right)

			root.ok = false
			root.value = replacement.value
			root.right = BstCustomRemove(root.right, replacement.value)
		}
	}

	return root
}
