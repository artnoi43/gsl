package tree

type nodeWrapper[T comparable] struct {
	value T
	null  bool

	left  BinaryTreeNode[T]
	right BinaryTreeNode[T]
}

type BinaryTreeImpl[NODE comparable] struct {
	count int
	root  nodeWrapper[NODE]
}

func (b *BinaryTreeImpl[T]) Insert(node T) {
	panic("not implemented")
}

func (b *BinaryTreeImpl[T]) Remove(node T) {
	panic("not implemented")
}

func (b *BinaryTreeImpl[T]) Parent(node T) T {
	panic("not implemented")
}

func (b *BinaryTreeImpl[T]) LeftChild(node T) T {
	panic("not implemented")
}

func (b *BinaryTreeImpl[T]) RightChild(node T) T {
	panic("not implemented")
}

func (b *BinaryTreeImpl[T]) Node(node T) T {
	panic("not implemented")
}

func (b *BinaryTreeImpl[T]) IsRoot(node T) bool {
	panic("not implemented")
}

func (b *BinaryTreeImpl[T]) IsNull(node T) bool {
	panic("not implemented")
}
