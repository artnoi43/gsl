package tree

type BinaryTree interface {
	Parent(node int) int
	LeftChild(node int) int
	RightChild(node int) int

	IsRoot(node int) bool
	IsLeaf(node int) bool
}

type binaryTreeSlice[T any] struct {
	backing []T
}

func (b *binaryTreeSlice[T]) Parent(node int) int {
	return ParentIdx(node)
}

func (b *binaryTreeSlice[T]) LeftChild(node int) int {
	return LeftChildIdx(node)
}

func (b *binaryTreeSlice[T]) RightChild(node int) int {
	return RightChildIdx(node)
}

func (b *binaryTreeSlice[T]) IsRoot(node int) bool {
	return node == 0
}

func (b *binaryTreeSlice[T]) IsLeaf(node int) bool {
	return node >= len(b.backing)
}

func LeftChildIdx(parent int) int {
	return (2 * parent) + 1
}

func RightChildIdx(parent int) int {
	return LeftChildIdx(parent) + 1
}

func ParentIdx(child int) int {
	rightChild := child%2 == 0
	if rightChild {
		return (child - 2) / 2
	}

	return (child - 1) / 2
}
