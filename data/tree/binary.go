package tree

import (
	"github.com/soyart/gsl/data/list"
)

// BinaryTreeBasic is basic, minimal binary tree with node type NODE
type BinaryTreeBasic[NODE any] interface {
	Insert(node NODE)
	Remove(node NODE) bool
	Find(node NODE) bool
}

// BinaryTree have extra methods to work with nodes.
// POS is any type used for indexing a node,
// e.g. a bintree with backing arrays may use int as POS.
type BinaryTree[POS any, NODE any] interface {
	BinaryTreeBasic[NODE]

	Parent(node POS) POS
	LeftChild(node POS) POS
	RightChild(node POS) POS
	Node(pos POS) NODE

	NodeIsRoot(node POS) bool
	NodeIsNull(node POS) bool
}

type BinaryTreeNode[T any] interface {
	Value() T
	Left() BinaryTreeNode[T]
	Right() BinaryTreeNode[T]
	IsNull() bool
}

type binTreeNode[T comparable] struct {
	value T
	ok    bool

	left  *binTreeNode[T]
	right *binTreeNode[T]
}

func (n *binTreeNode[T]) Left() BinaryTreeNode[T] { return n.left }

func (n *binTreeNode[T]) Right() BinaryTreeNode[T] { return n.right }

func (n *binTreeNode[T]) Value() T { return n.value }

func (n *binTreeNode[T]) IsNull() bool {
	return !n.ok &&
		n.left == nil &&
		n.right == nil
}

func (n *binTreeNode[T]) IsLeaf() bool {
	return n.left == nil && n.right == nil
}

func (n *binTreeNode[T]) hasBoth() bool {
	return n.left != nil && n.right != nil
}

func (n *binTreeNode[T]) hasLeft() bool {
	return n.left != nil
}

func (n *binTreeNode[T]) hasRight() bool {
	return n.right != nil
}

func Inorder[POS any, NODE any](
	tree BinaryTree[POS, NODE],
	node POS,
	f func(NODE) error,
) error {
	stack := list.NewStack[POS]()
	curr := node

	for !tree.NodeIsNull(curr) || !stack.IsEmpty() {

		for !tree.NodeIsNull(curr) {
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

func InorderNode[T any, NODE BinaryTreeNode[T]](
	tree BinaryTreeBasic[NODE],
	node NODE,
	f func(NODE) error,
) error {
	stack := list.NewStack[NODE]()
	curr := node

	for !curr.IsNull() || !stack.IsEmpty() {
		for !curr.IsNull() {
			stack.Push(curr)
			curr = curr.Left().(NODE)
		}

		if err := f(curr); err != nil {
			return err
		}

		curr = *stack.Pop()
		curr = curr.Right().(NODE)
	}

	return nil
}

func InorderNodeRecurse[T any, NODE BinaryTreeNode[T]](
	node NODE,
	f func(NODE) error,
) error {
	if err := InorderNodeRecurse[T, NODE](node.Left().(NODE), f); err != nil {
		return err
	}

	if err := f(node); err != nil {
		return err
	}

	if err := InorderNodeRecurse[T, NODE](node.Right().(NODE), f); err != nil {
		return err
	}

	return nil
}

// BinTreeMax digs for smallest values in the subtree, returning
// the node as well as the height.
func BinTreeMax[T any](node BinaryTreeNode[T]) (BinaryTreeNode[T], uint) {
	curr := node
	var height uint

	for curr.Right() != nil && !curr.Right().IsNull() {
		curr = curr.Right()
		height++
	}

	return curr, height
}

// BinTreeMin digs for smallest values in the subtree, returning
// the node as well as the height.
func BinTreeMin[T any](root BinaryTreeNode[T]) (BinaryTreeNode[T], uint) {
	curr := root
	var height uint

	for curr.Left() != nil && !curr.Left().IsNull() {
		curr = curr.Left()
		height++
	}

	return curr, height
}
