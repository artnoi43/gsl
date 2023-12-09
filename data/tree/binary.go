package tree

import "github.com/soyart/gsl/data/list"

type BinaryTree[POS any, NODE any] interface {
	Parent(node POS) POS
	LeftChild(node POS) POS
	RightChild(node POS) POS
	Node(pos POS) NODE

	IsRoot(node POS) bool
	IsNull(node POS) bool
}

func Inorder[POS any, NODE any](
	tree BinaryTree[POS, NODE],
	node POS,
	f func(NODE) error,
) error {
	stack := list.NewStack[POS]()
	curr := node

	for !tree.IsNull(curr) || !stack.IsEmpty() {
		for !tree.IsNull(curr) {
			stack.Push(curr)
			curr = tree.LeftChild(curr)
		}

		curr = *stack.Pop()
		if err := f(tree.Node(curr)); err != nil {
			return err
		}

		curr = tree.RightChild(curr)
	}

	return nil
}

func InorderRecurse[POS any, NODE any](
	tree BinaryTree[POS, NODE],
	node POS,
	f func(NODE) error,
) error {
	if err := InorderRecurse(tree, tree.LeftChild(node), f); err != nil {
		return err
	}

	if err := f(tree.Node(node)); err != nil {
		return err
	}

	return InorderRecurse(tree, tree.RightChild(node), f)
}
