package tree

type BinaryTree[POS any, NODE any] interface {
	Parent(node POS) POS
	LeftChild(node POS) POS
	RightChild(node POS) POS
	Node(pos POS) NODE

	IsRoot(node POS) bool
	IsLeaf(node POS) bool
}
