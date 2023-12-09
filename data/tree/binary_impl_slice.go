package tree

// Binary tree implementation with backing array
type BinaryTreeSlice[T any] struct {
	backing []T
}

func NewBinaryTreeSlice[T any]() BinaryTree[int, T] {
	return &BinaryTreeSlice[T]{}
}

func (b *BinaryTreeSlice[T]) Parent(node int) int {
	return ParentIdx(node)
}

func (b *BinaryTreeSlice[T]) LeftChild(node int) int {
	return LeftChildIdx(node)
}

func (b *BinaryTreeSlice[T]) RightChild(node int) int {
	return RightChildIdx(node)
}

func (b *BinaryTreeSlice[T]) Node(node int) T {
	return b.backing[node]
}

func (b *BinaryTreeSlice[T]) IsRoot(node int) bool {
	return node == 0
}

func (b *BinaryTreeSlice[T]) IsNull(node int) bool {
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
