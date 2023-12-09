package tree

type BinaryTree[POS any, NODE any] interface {
	Parent(node POS) POS
	LeftChild(node POS) POS
	RightChild(node POS) POS
	Node(pos POS) NODE

	IsRoot(node POS) bool
	IsLeaf(node POS) bool
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
