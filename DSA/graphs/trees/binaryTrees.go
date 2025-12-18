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

func (b BinaryTree[T]) IsEmpty() bool {
	return b.Root == nil
}

// Exercise 1: Implement a function to insert a value into a binary search tree.
func (b *BinaryTree[T]) insertValue(value T) {
	// If it is empty adds the value as the root node
	if b.IsEmpty() {
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
	if b.IsEmpty() {
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

func (b BinaryTree[T]) lowestCommonAncestor(node1 *Node[T], node2 *Node[T]) *Node[T] {
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

	lca := tree.lowestCommonAncestor(node1, node2)
	if lca != nil {
		fmt.Printf("LCA of %d and %d is: %d\n", val1, val2, lca.Value)
	} else {
		fmt.Println("LCA not found.")
	}
}

// Exercise 8: Write a function to count the number of leaf nodes in a binary tree.

// Helper function to find the number of leafs of a tree
func numberLeaf[T Number](node *Node[T]) int {
	if node == nil {
		return 0
	}
	if node.Right == nil && node.Left == nil {
		return 1
	}
	return numberLeaf(node.Left) + numberLeaf(node.Right)
}

// main function to find the number of leafs of a tree
func (b BinaryTree[T]) numberOfLeaf() int {
	return numberLeaf(b.Root)
}

func Main8() {
	// Build the tree
	tree := BinaryTree[int]{}
	values := []int{8, 3, 10, 1, 6, 14, 4, 7, 13}
	for _, v := range values {
		tree.insertValue(v)
	}

	// Count the number of leaf nodes
	leafCount := tree.numberOfLeaf()
	fmt.Printf("Number of leaf nodes: %d\n", leafCount)
}

// Exercise 9: Implement a function to check if a binary tree is a valid binary search tree (BST).
func isValidBSTNode[T Number](node *Node[T], min, max *T) bool {
	if node == nil {
		return true
	}
	if (min != nil && node.Value <= *min) || (max != nil && node.Value >= *max) {
		return false
	}
	return isValidBSTNode(node.Left, min, &node.Value) && isValidBSTNode(node.Right, &node.Value, max)
}

func (b BinaryTree[T]) isValidBST() bool {
	if b.IsEmpty() {
		return true
	}
	var zero *T
	return isValidBSTNode(b.Root, zero, zero)
}

func Main9() {
	// Valid BST
	tree := BinaryTree[int]{}
	values := []int{8, 3, 10, 1, 6, 14, 4, 7, 13}
	for _, v := range values {
		tree.insertValue(v)
	}
	fmt.Print("Valid BST: ")
	if tree.isValidBST() {
		fmt.Println("Tree is a valid BST.")
	} else {
		fmt.Println("Tree is NOT a valid BST.")
	}

	// Invalid BST (manually break the BST property)
	badTree := BinaryTree[int]{}
	badTree.Root = &Node[int]{Value: 8}
	badTree.Root.Left = &Node[int]{Value: 3}
	badTree.Root.Right = &Node[int]{Value: 10}
	badTree.Root.Left.Left = &Node[int]{Value: 1}
	badTree.Root.Left.Right = &Node[int]{Value: 9} // 9 > 8, should not be in left subtree

	fmt.Print("Invalid BST: ")
	if badTree.isValidBST() {
		fmt.Println("Tree is a valid BST.")
	} else {
		fmt.Println("Tree is NOT a valid BST.")
	}
}

// Exercise 10: Write a function to mirror (invert) a binary tree.
func invertNode[T Number](node *Node[T]) {
	if node == nil {
		return
	}
	node.Left, node.Right = node.Right, node.Left
	invertNode(node.Left)
	invertNode(node.Right)
}

func (b BinaryTree[T]) mirror() {
	if b.IsEmpty() {
		return
	}
	invertNode(b.Root)
}

// Exercise 11: Implement a function to print all printPaths from root to leaf in a binary tree.
func printPaths[T Number](node *Node[T], currentPath []T) {
	if node == nil {
		return
	}
	currentPath = append(currentPath, node.Value)
	if node.Left == nil && node.Right == nil {
		printSlice(currentPath)
	}
	printPaths(node.Left, currentPath)
	printPaths(node.Right, currentPath)
}

func (b BinaryTree[T]) printAllPaths() {
	if b.IsEmpty() {
		return
	}
	var path []T
	printPaths(b.Root, path)
}

// Exercise 12: Write a function to calculate the sum of all nodes in a binary tree.
func sumNodes[T Number](node *Node[T]) T {
	if node == nil {
		return 0
	}
	return node.Value + sumNodes(node.Left) + sumNodes(node.Right)
}
func (b BinaryTree[T]) sumAllNodes() T {
	return sumNodes(b.Root)
}

// Exercise 13: Implement a function to delete a node from a binary search tree.
func (b *BinaryTree[T]) deleteNode(value T) error {
	if b.IsEmpty() {
		err := errors.New("imposible to delete value, tree is empty")
		return err
	}
	deletedNode := findFirstNodeWithValueDFS(b.Root, value)
	// If the node has no children
	if deletedNode.Right == nil && deletedNode.Left == nil {
		deletedNode = nil
		return nil
	}

	// If the node has only one child
	if (deletedNode.Right == nil && deletedNode.Left != nil) ||
		(deletedNode.Right != nil && deletedNode.Left == nil) {
		// do the implementation
		return nil
	}

	// If the node has two children
	if deletedNode.Right != nil && deletedNode.Left != nil {
		// complete the implementation
		return nil
	}

	return nil
}

// Exercise 14: Write a function to check if two binary trees are identical.
func areIdenticarNodes[T Number](node1 *Node[T], node2 *Node[T]) bool {
	if node1 == nil && node2 == nil {
		return true
	}
	if (node1 == nil && node2 != nil) || (node1 != nil && node2 == nil) {
		return false
	}
	if node1.Value != node2.Value {
		return false
	}

	left := areIdenticarNodes(node1.Left, node2.Left)
	right := areIdenticarNodes(node1.Right, node2.Right)
	return left && right
}

func areIdenticalTrees[T Number](tree1, tree2 BinaryTree[T]) bool {
	return areIdenticarNodes(tree1.Root, tree2.Root)
}

// Exercise 15: Implement a function to find the diameter (longest path) of a binary tree.
func diameter[T Number](node *Node[T], maxDiameter *int) int {
	if node == nil {
		return 0
	}
	leftHeight := diameter(node.Left, maxDiameter)
	rightHeight := diameter(node.Right, maxDiameter)
	if leftHeight+rightHeight > *maxDiameter {
		*maxDiameter = leftHeight + rightHeight
	}
	return 1 + max(leftHeight, rightHeight)
}
func (b BinaryTree[T]) longestPath() int {
	maxDiameter := 0
	diameter(b.Root, &maxDiameter)
	return maxDiameter
}
