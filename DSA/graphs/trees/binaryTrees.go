package binaryTrees

import (
	"errors"
	"fmt"
	"math"
)

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
func printSlice[A any](slice []A) {
	for i, value := range slice {
		fmt.Print(value)
		if i < len(slice)-1 {
			fmt.Print(" -> ")
		}
	}
	fmt.Println()
}

func preTraverse[T Number](node *Node[T], slice *[]T) {
	if node == nil {
		return
	}

	*slice = append(*slice, node.Value)
	preTraverse(node.Left, slice)
	preTraverse(node.Right, slice)
}

func (b BinaryTree[T]) preorderTraversal() []T {
	var result []T
	firstNode := b.Root
	preTraverse(firstNode, &result)
	return result
}

func inTraverse[T Number](node *Node[T], slice *[]T) {
	if node == nil {
		return
	}

	inTraverse(node.Left, slice)
	*slice = append(*slice, node.Value)
	inTraverse(node.Right, slice)
}

func (b BinaryTree[T]) inorderTraversal() []T {
	var result []T
	inTraverse(b.Root, &result)
	return result
}

func postTraverse[T Number](node *Node[T], slice *[]T) {
	if node == nil {
		return
	}

	postTraverse(node.Left, slice)
	postTraverse(node.Right, slice)
	*slice = append(*slice, node.Value)
}

func (b BinaryTree[T]) postorderTraversal() []T {
	var result []T
	postTraverse(b.Root, &result)
	return result
}

func Main5() {
	// Create a binary tree and insert some values
	tree := BinaryTree[int]{}
	values := []int{8, 3, 10, 1, 6, 14, 4, 7, 13}
	for _, v := range values {
		tree.insertValue(v)
	}

	// Test the 3 types of traversals
	pre := tree.preorderTraversal()
	in := tree.inorderTraversal()
	post := tree.postorderTraversal()

	fmt.Println("Preorder traversal:")
	printSlice(pre)
	fmt.Println("Inorder traversal:")
	printSlice(in)
	fmt.Println("Postorder traversal:")
	printSlice(post)
}

// Exercise 6: Write a function to check if a binary tree is balanced.
func subtreeHeight[T Number](node *Node[T]) int {
	if node == nil {
		return 0
	}
	left := subtreeHeight(node.Left)
	right := subtreeHeight(node.Right)
	if left > right {
		return left + 1
	}
	return right + 1
}

func isBalancedHelper[T Number](node *Node[T]) bool {
	if node == nil {
		return true
	}
	leftHeight := subtreeHeight(node.Left)
	rightHeight := subtreeHeight(node.Right)
	if math.Abs(float64(leftHeight-rightHeight)) > 1 {
		return false
	}
	return isBalancedHelper(node.Left) && isBalancedHelper(node.Right)
}

func (b BinaryTree[T]) isBalanced() bool {
	return isBalancedHelper(b.Root)
}

func Main6() {
	// Example 1: Balanced tree (should print "The tree is balanced.")
	tree1 := BinaryTree[int]{}
	values1 := []int{8, 3, 10, 1, 6, 14, 4, 7, 13}
	for _, v := range values1 {
		tree1.insertValue(v)
	}
	fmt.Print("Example 1: ")
	if tree1.isBalanced() {
		fmt.Println("The tree is balanced.")
	} else {
		fmt.Println("The tree is not balanced.")
	}

	// Example 2: Unbalanced tree (right-skewed, should print "not balanced")
	tree2 := BinaryTree[int]{}
	values2 := []int{1, 2, 3, 4, 5}
	for _, v := range values2 {
		tree2.insertValue(v)
	}
	fmt.Print("Example 2: ")
	if tree2.isBalanced() {
		fmt.Println("The tree is balanced.")
	} else {
		fmt.Println("The tree is not balanced.")
	}

	// Example 3: Unbalanced tree (left-skewed, should print "not balanced")
	tree3 := BinaryTree[int]{}
	values3 := []int{5, 4, 3, 2, 1}
	for _, v := range values3 {
		tree3.insertValue(v)
	}
	fmt.Print("Example 3: ")
	if tree3.isBalanced() {
		fmt.Println("The tree is balanced.")
	} else {
		fmt.Println("The tree is not balanced.")
	}

	// Example 4: Small balanced tree (should print "balanced")
	tree4 := BinaryTree[int]{}
	values4 := []int{2, 1, 3}
	for _, v := range values4 {
		tree4.insertValue(v)
	}
	fmt.Print("Example 4: ")
	if tree4.isBalanced() {
		fmt.Println("The tree is balanced.")
	} else {
		fmt.Println("The tree is not balanced.")
	}
}

// Exercise 7: Implement a function to find the lowest common ancestor (LCA) of two nodes in a binary tree.

func findFirstNodeWithValueDFS[T Number](node *Node[T], value T) *Node[T] {
	if node == nil {
		return nil
	}
	if node.Value == value {
		return node
	}
	if left := findFirstNodeWithValueDFS(node.Left, value); left != nil {
		return left
	}
	return findFirstNodeWithValueDFS(node.Right, value)
}

// TODO: finish implementation
// func (b BinaryTree[T]) lowestCommontAncestorByvalue(value1 T, value2 T) T {
// 	var result T
// 	return result
// }

// Helper function to find LCA given two nodes
func LCA[T Number](currentNode, node1, node2 *Node[T]) *Node[T] {
	if currentNode == nil {
		return nil
	}
	if currentNode == node1 || currentNode == node2 {
		return currentNode
	}
	left := LCA(currentNode.Left, node1, node2)
	right := LCA(currentNode.Right, node1, node2)
	if left != nil && right != nil {
		return currentNode
	}
	if left != nil {
		return left
	}
	return right
}

func (b BinaryTree[T]) lowestCommontAncestor(node1 *Node[T], node2 *Node[T]) *Node[T] {
	return LCA(b.Root, node1, node2)
}

func Main7() {
	// Build the tree
	tree := BinaryTree[int]{}
	values := []int{8, 3, 10, 1, 6, 14, 4, 7, 13}
	for _, v := range values {
		tree.insertValue(v)
	}

	// Choose two nodes to find LCA for
	val1, val2 := 1, 7
	node1 := findFirstNodeWithValueDFS(tree.Root, val1)
	node2 := findFirstNodeWithValueDFS(tree.Root, val2)

	if node1 == nil || node2 == nil {
		fmt.Printf("One or both nodes not found: %d, %d\n", val1, val2)
		return
	}

	lca := tree.lowestCommontAncestor(node1, node2)
	if lca != nil {
		fmt.Printf("LCA of %d and %d is: %d\n", val1, val2, lca.Value)
	} else {
		fmt.Println("LCA not found.")
	}
}

// Exercise 8: Write a function to count the number of leaf nodes in a binary tree.

// Exercise 9: Implement a function to check if a binary tree is a valid binary search tree (BST).

// Exercise 10: Write a function to mirror (invert) a binary tree.

// Exercise 11: Implement a function to print all paths from root to leaf in a binary tree.

// Exercise 12: Write a function to calculate the sum of all nodes in a binary tree.

// Exercise 13: Implement a function to delete a node from a binary search tree.

// Exercise 14: Write a function to check if two binary trees are identical.

// Exercise 15: Implement a function to find the diameter (longest path) of a binary tree.
