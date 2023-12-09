package tree

type BinaryTree interface {
	Parent(node int) int
	LeftChild(node int) int
	RightChild(node int) int

	IsRoot(node int) int
	IsLead(node int) int
}
