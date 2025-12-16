package binaryTrees

import "errors"

type Number interface {
	~int | ~float64
}

type Node[T Number] struct {
	Value T
	Left  *Node[T]
	Right *Node[T]
}

type BinaryTree[T Number] struct {
	Root *Node[T]
}

// Exercise 1: Implement a function to insert a value into a binary search tree.
func (b *BinaryTree[T]) insertValue(value T) {
	// If it is empty adds the value as the root node
	if b.Root == nil {
		b.Root = &Node[T]{Value: value}
		return
	}

	// When is not empty compares the value with each node
	currentNode := b.Root
	for {
		if value < currentNode.Value {
			if currentNode.Left == nil {
				currentNode.Left = &Node[T]{Value: value}
				return
			}
			currentNode = currentNode.Left
			continue
		}
		if currentNode.Right == nil {
			currentNode.Right = &Node[T]{Value: value}
			return
		}
		currentNode = currentNode.Right
	}
}

// Exercise 2: Write a function to search for a value in a binary tree (not necessarily a BST).
func recursiveDFS[T Number](node *Node[T], value T) bool {
	if node == nil {
		return false
	}
	if node.Value == value {
		return true
	}
	return recursiveDFS(node.Left, value) || recursiveDFS(node.Right, value)
}

func (b BinaryTree[T]) deepFirstSearch(value T) bool {
	return recursiveDFS(b.Root, value)
}

// Exercise 3: Implement a function to find the minimum and maximum value in a binary search tree.
func (b BinaryTree[T]) minMaxBST() (min T, max T, err error) {
	if b.Root == nil {
		err = errors.New("binary tree is empty")
		var zero T
		min = zero
		max = zero
		return min, max, err
	}
	minNode := b.Root
	for minNode.Left != nil {
		minNode = minNode.Left
	}
	maxNode := b.Root
	for maxNode.Right != nil {
		maxNode = maxNode.Right
	}

	min = minNode.Value
	max = maxNode.Value
	return min, max, nil
}

// Exercise 4: Write a function to calculate the height (maximum depth) of a binary tree.
func recursiveHeight[T Number](node *Node[T]) int {
	if node == nil {
		return 0
	}
	left := recursiveHeight(node.Left)
	right := recursiveHeight(node.Right)

	if left > right {
		return left + 1
	}
	return right + 1
}

func (b BinaryTree[T]) maximumDepth() int {
	return recursiveHeight(b.Root)
}

// Exercise 5: Implement functions for preorder, inorder, and postorder traversals of a binary tree.

// Exercise 6: Write a function to check if a binary tree is balanced.

// Exercise 7: Implement a function to find the lowest common ancestor (LCA) of two nodes in a binary tree.

// Exercise 8: Write a function to count the number of leaf nodes in a binary tree.

// Exercise 9: Implement a function to check if a binary tree is a valid binary search tree (BST).

// Exercise 10: Write a function to mirror (invert) a binary tree.

// Exercise 11: Implement a function to print all paths from root to leaf in a binary tree.

// Exercise 12: Write a function to calculate the sum of all nodes in a binary tree.

// Exercise 13: Implement a function to delete a node from a binary search tree.

// Exercise 14: Write a function to check if two binary trees are identical.

// Exercise 15: Implement a function to find the diameter (longest path) of a binary tree.
