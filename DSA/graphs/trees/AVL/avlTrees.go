package avl

// ===================== AVL TREE EXERCISES =====================
// 1. Implement an AVL tree node structure in Go.

// Adelson-Velsky and Landis generic node
import (
	"Practica_GO/DSA/generics"
	binaryTrees "Practica_GO/DSA/graphs/trees/binaryTrees"
	utils "Practica_GO/DSA/utils"
	"errors"
	"fmt"
)

type Number utils.Number

type AVLNode[N Number] struct {
	value N
	left  *AVLNode[N]
	right *AVLNode[N]
}

func (node *AVLNode[N]) isLeaf() bool {
	if node == nil {
		return false
	}
	if node.left == nil && node.right == nil {
		return true
	}
	return false
}

func (node *AVLNode[N]) printValue() {
	if node == nil {
		return
	}

	fmt.Printf("%v ", node.value)
}

// Adelson-Velsky and Landis generic tree
type AVLTree[N Number] struct {
	Root *AVLNode[N]
}

func (tree AVLTree[A]) IsEmpty() bool {
	return tree.Root == nil
}

// Computes the height of the subtree with the node as Root node
func (node *AVLNode[A]) subtreeHeight() int {
	if node == nil {
		return 0
	}
	leftHeight := node.left.subtreeHeight()
	rightHeight := node.right.subtreeHeight()
	if leftHeight > rightHeight {
		return leftHeight + 1
	}
	return rightHeight + 1
}

// Computes the balance factor of a node: the height of the subtree with the left node as root -
// the height of the subtree with the right node as root
func (node *AVLNode[A]) balanceFactor() int {
	if node == nil {
		return 0
	}
	return node.left.subtreeHeight() - node.right.subtreeHeight()
}

// Checks if a node is unbalanced, meaning, its balance factor is outside of {-1, 0, 1} values.
func (node *AVLNode[A]) isBalanced() bool {
	balanceFactor := node.balanceFactor()
	return balanceFactor == -1 || balanceFactor == 0 || balanceFactor == 1
}

// 2. Write a function to insert a value into an AVL tree and maintain balance.

// Perform a left-left rotation (right rotation)
func (node *AVLNode[A]) LLRotation() *AVLNode[A] {
	leftChild := node.left
	node.left = leftChild.right
	leftChild.right = node
	return leftChild
}

// Perform a right-right rotation (left rotation)
func (node *AVLNode[A]) RRRotation() *AVLNode[A] {
	rightChild := node.right
	node.right = rightChild.left
	rightChild.left = node
	return rightChild
}

// Perform a left-right rotation
func (node *AVLNode[A]) LRRotation() *AVLNode[A] {
	node.left = node.left.RRRotation()
	return node.LLRotation()
}

// Perform a right-left rotation
func (node *AVLNode[A]) RLRotation() *AVLNode[A] {
	node.right = node.right.LLRotation()
	return node.RRRotation()
}

// Detect which rotation type use:
//
//	ParentNodeBF | ChildNodeBF | RotationType | Description
//	+2           | 0, 1        | RR           | Left subtree heavy, right rotation
//	+2           | -1          | LR           | Left subtree heavy with bent arm, left-right rotation
//	-2           | -1, 0       | LL           | Right subtree heavy, left rotation
//	-2           | +1          | RL           | Right subtree heavy with bent arm, right-left rotation
func (node *AVLNode[A]) rotationType() string {
	bf := node.balanceFactor()
	if bf > 1 {
		// Left heavy
		if node.left.balanceFactor() >= 0 {
			return "LL"
		} else {
			return "LR"
		}
	} else if bf < -1 {
		// Right heavy
		if node.right.balanceFactor() <= 0 {
			return "RR"
		} else {
			return "RL"
		}
	}
	return ""
}

// rotates according the rotation type
func (node *AVLNode[A]) rotate(rotationType string) *AVLNode[A] {
	switch rotationType {
	case "LL":
		return node.LLRotation()
	case "RR":
		return node.RRRotation()
	case "RL":
		return node.RLRotation()
	case "LR":
		return node.LRRotation()
	default:
		return node
	}
}

// AVLBalance checks the balance of the AVL tree and performs the necessary rotations to maintain the AVL property.
// func (a *AVLTree[A]) AVLBalance() {
// 	parentNode := a.Root
// 	for parentNode != nil {
// 		parentBalanceFactor := balanceFactor(parentNode)
// 		// if the node is balanced just jump to the next node
// 		if parentBalanceFactor >= -1 && parentBalanceFactor <= 1 {
// 			parentNode = parentNode.left
// 		}
// 		leftChild := parentNode.left
// 		leftChildBalanceFactor := balanceFactor(leftChild)
// 		rightChild := parentNode.right
// 		rightChildBalanceFactor := balanceFactor(rightChild)

// 	}
// }

// Recursive function to insert an element in a BST way and balancing it to keep the AVL structure
func (tree *AVLTree[A]) insertHelper(node *AVLNode[A], value A) *AVLNode[A] {
	if node == nil {
		return &AVLNode[A]{value: value}
	}
	if value < node.value {
		node.left = tree.insertHelper(node.left, value)
	} else {
		node.right = tree.insertHelper(node.right, value)
	}
	// if the balance factor is okay just return the node
	if node.isBalanced() {
		return node
	}
	// rebalance the node with the correct rotation type
	rType := node.rotationType()
	return node.rotate(rType)
}

// Function that inserts a value in the AVL tree keeping the structure of the tree
func (tree *AVLTree[A]) Insert(value A) {
	tree.Root = tree.insertHelper(tree.Root, value)
}

// 3. Write a function to delete a value from an AVL tree and maintain balance.

func (tree AVLTree[A]) delete(value A) error {
	return nil
}

// the following exercises are already done above

// 4. Implement a function to perform a left rotation on a subtree.
// 5. Implement a function to perform a right rotation on a subtree.
// 6. Write a function to get the height of a node in an AVL tree.
// 7. Write a function to get the balance factor of a node in an AVL tree.
// 8. Implement a function to search for a value in an AVL tree.

// returns true if the value is in the AVL tree, and false if not
func (tree AVLTree[N]) search(value N) bool {
	node := tree.Root
	for node != nil {
		if value == node.value {
			return true
		}
		if value < node.value {
			node = node.left
			continue
		}
		if value > node.value {
			node = node.right
		}
	}
	return false
}

// 9. Write a function to find the minimum value in an AVL tree.

// returns the minimum value of an AVL, returns error and 0 if the tree is empty
func (tree AVLTree[N]) min() (N, error) {
	var minimum N
	if tree.IsEmpty() {
		err := errors.New("tree is empty")
		return minimum, err
	}
	node := tree.Root
	for node.left != nil {
		node = node.left
	}
	return node.value, nil
}

// 10. Write a function to find the maximum value in an AVL tree.

// returns the minimum value of an AVL, returns error and 0 if the tree is empty
func (tree AVLTree[N]) max() (N, error) {
	var minimum N
	if tree.IsEmpty() {
		err := errors.New("tree is empty")
		return minimum, err
	}
	node := tree.Root
	for node.right != nil {
		node = node.right
	}
	return node.value, nil
}

// 11. Implement preorder, inorder, and postorder traversals for an AVL tree.

// helper function to apply recursion in preorder
func (node *AVLNode[N]) preTraverse() {
	if node == nil {
		return
	}
	node.printValue()
	node.left.preTraverse()
	node.right.preTraverse()
}

// visit a node, then left subtree, then right subtree
func (tree AVLTree[N]) PreorderTraversal() {
	node := tree.Root
	node.preTraverse()
	fmt.Println()
}

// helper function to apply recursion in postorder
func (node *AVLNode[N]) postTraverse() {
	if node == nil {
		return
	}
	node.left.postTraverse()
	node.right.postTraverse()
	node.printValue()
}

// visit the left subtree, then right subtree, then node
func (tree AVLTree[N]) PostorderTraversal() {
	node := tree.Root
	node.postTraverse()
	fmt.Println()
}

// helper function to apply recursion in inorder
func (node *AVLNode[N]) inTraverse() {
	if node == nil {
		return
	}
	node.left.inTraverse()
	node.printValue()
	node.right.inTraverse()
}

// visit left subtree, then right node, then right subtree
func (tree AVLTree[N]) InorderTraversal() {
	node := tree.Root
	node.inTraverse()
	fmt.Println()
}

// 12. Write a function to check if a given binary tree is a valid AVL tree.

// Return true if is an AVL tree: must be a BST and each node must have |BF| <= 1
func isAVLTree(binaryTree binaryTrees.BinaryTree[int]) bool {
	if !binaryTree.IsValidBST() || !binaryTree.IsBalanced() {
		return false
	}
	return true
}

// 13. Write a function to print all nodes at a given level in an AVL tree.

// helper function to print every node at a certain height
func (node *AVLNode[N]) preTravNodes(level int, current int) {
	if node == nil {
		return
	}
	if level == current {
		node.printValue()
		return
	}
	node.left.preTravNodes(level, current+1)
	node.right.preTravNodes(level, current+1)
}

// Print every node at the specified height, if the chosen height is empty returns an error
func (tree AVLTree[N]) printNodesAtHeight(height int) error {
	if height < 0 {
		return errors.New("height must be a positive integer")
	}
	if tree.IsEmpty() {
		return errors.New("tree is empty")
	}
	if height >= tree.Root.subtreeHeight() {
		return errors.New("the chosen height is greater than the tree height")
	}
	if height == 0 {
		tree.Root.printValue()
		return nil
	}
	tree.Root.preTravNodes(height, 0)
	return nil
}

// 14. Implement a function to count the number of nodes in an AVL tree.

func (node *AVLNode[N]) sizeHelper() int {
	if node == nil {
		return 0
	}
	return node.left.sizeHelper() + node.right.sizeHelper() + 1
}

// returns the number of nodes in the tree
func (tree AVLTree[N]) Size() int {
	if tree.IsEmpty() {
		return 0
	}
	return tree.Root.sizeHelper()
}

// 15. Write a function to count the number of leaf nodes in an AVL tree.

func (node *AVLNode[N]) leafNumberHelper() int {
	if node == nil {
		return 0
	}
	if node.left == nil && node.right == nil {
		return 1
	}
	return node.left.leafNumberHelper() + node.right.leafNumberHelper()
}

// returns the number of leaf nodes: nodes without children
func (tree AVLTree[N]) LeafNumber() int {
	if tree.IsEmpty() {
		return 0
	}
	return tree.Root.leafNumberHelper()
}

// 16. Write a function to calculate the sum of all node values in an AVL tree.

func (node *AVLNode[N]) sumNodesHelper(sum *N) N {
	if node == nil {
		var zero N
		return zero
	}
	*sum = *sum + node.value
	return node.left.sumNodesHelper(sum) + node.right.sumNodesHelper(sum)
}

// returns the total sum of every node in the tree
func (tree AVLTree[N]) SumNodes() N {
	var sum N
	if tree.IsEmpty() {
		return sum
	}
	tree.Root.sumNodesHelper(&sum)
	return sum
}

// 17. Implement a function to find the lowest common ancestor (LCA) of two nodes in an AVL tree.

func (node *AVLNode[N]) dfsHelper(value N) *AVLNode[N] {
	if node == nil {
		return nil
	}
	if value == node.value {
		return node
	}
	if left := node.left.dfsHelper(value); left != nil {
		return left
	}
	return node.right.dfsHelper(value)

}

// Deep First Search
func (tree AVLTree[N]) DFS(value N) *AVLNode[N] {
	if tree.IsEmpty() {
		var nilNode *AVLNode[N]
		return nilNode
	}
	node := tree.Root.dfsHelper(value)
	return node
}

func (node *AVLNode[N]) lcaHelper(value1, value2 N) *AVLNode[N] {
	if node == nil {
		return nil
	}
	if value1 < node.value && value2 < node.value {
		return node.left.lcaHelper(value1, value2)
	}
	if value1 > node.value && value2 > node.value {
		return node.right.lcaHelper(value1, value2)
	}
	return node
}

// find the Lowest Common Ancestor (LCA) of two nodes
func (tree AVLTree[N]) LCA(value1, value2 N) (N, error) {
	if tree.IsEmpty() {
		var zero N
		return zero, errors.New("tree is empty")
	}
	// dfsearchs for the nodes with the values provided
	node1 := tree.DFS(value1)
	if node1 == nil {
		var zero N
		return zero, fmt.Errorf("value : %v was not found", value1)
	}
	node2 := tree.DFS(value2)
	if node2 == nil {
		var zero N
		return zero, fmt.Errorf("value : %v was not found", value2)
	}

	lcaNode := node1.lcaHelper(value1, node1.value)
	if lcaNode == nil {
		var zero N
		return zero, errors.New("there is not an LCA")
	}
	return lcaNode.value, nil
}

// 18. Write a function to mirror (invert) an AVL tree.

func (node *AVLNode[N]) invertHelper() {
	if node == nil {
		return
	}
	// swaps the left subtree and the right subtree
	node.left, node.right = node.right, node.left
	node.left.invertHelper()
	node.right.invertHelper()
}

// invert an AVL tree, each node swaps his left subtree and right subtree
func (tree AVLTree[N]) Invert() {
	tree.Root.invertHelper()
}

// 19. Write a function to print all root-to-leaf paths in an AVL tree.

func (node *AVLNode[N]) printAllRootToLeafPathsHelper(path *[]N) {
	if node == nil {
		return
	}
	*path = append(*path, node.value)
	if node.isLeaf() {
		generics.PrintSlice(*path)
		*path = (*path)[:len(*path)-1]
		return
	}
	node.left.printAllRootToLeafPathsHelper(path)
	node.right.printAllRootToLeafPathsHelper(path)
	*path = (*path)[:len(*path)-1]
}

func (tree AVLTree[N]) PrintAllRootToLeafPaths() {
	var path []N
	tree.Root.printAllRootToLeafPathsHelper(&path)
}

// 20. Implement a function to check if two AVL trees are identical.

func areNodesIdentical[N Number](node1, node2 *AVLNode[N]) bool {
	if node1 == nil && node2 == nil {
		return true
	}
	// the following checking returns false if one node is nil and the other one is not, the case when both are
	// nil is covered above
	if node2 == nil || node2 == nil {
		return false
	}
	if node1.value != node2.value {
		return false
	}
	return areNodesIdentical(node1.left, node2.left) && areNodesIdentical(node1.right, node2.right)
}

// returns true if the tree1 and tree2 are identical: have the same structure and the same node values
func areIdentical[N Number](tree1, tree2 AVLTree[N]) bool {
	return areNodesIdentical(tree1.Root, tree2.Root)
}

// 21. Write a function to find the diameter (longest path) of an AVL tree.

// helper function to get the diameter of the tree, it follows the post-order traversal
func (node *AVLNode[N]) diameterHelper(currentDiameter *int) int {
	if node == nil {
		return 0
	}
	leftHeight := node.left.diameterHelper(currentDiameter)
	rightHeight := node.right.diameterHelper(currentDiameter)
	// Update the max diameter if the path through this node is larger
	if leftHeight+rightHeight > *currentDiameter {
		*currentDiameter = leftHeight + rightHeight
	}
	// Return height of this node
	if leftHeight > rightHeight {
		return leftHeight + 1
	}
	return rightHeight + 1
}

// Diameter returns of the tree: the length (number of edges) of the longest path between any two leafs
func (tree AVLTree[N]) Diameter() int {
	if tree.IsEmpty() {
		return 0
	}
	var diameter int
	tree.Root.diameterHelper(&diameter)
	return diameter
}

// 22. Write a function to check if an AVL tree is balanced at every node.

func (node *AVLNode[N]) isBalancedHelper() bool {
	if node == nil {
		return true
	}
	if !node.isBalanced() {
		return false
	}
	return node.right.isBalanced() && node.left.isBalanced()
}

// checks if a AVL tree is balanced: every node has balance factor of -1, 0 or 1
func (tree AVLTree[N]) IsBalanced() bool {
	return tree.Root.isBalancedHelper()
}

// 23. Implement a function to convert a sorted array to a balanced AVL tree.
// 24. Write a function to convert an AVL tree to a sorted array (inorder traversal).
// 25. Implement a function to find the predecessor and successor of a given value in an AVL tree.
// 26. Write a function to find the kth smallest element in an AVL tree.
// 27. Write a function to find the kth largest element in an AVL tree.
// 28. Implement a function to clone (deep copy) an AVL tree.
// 29. Write a function to merge two AVL trees into a single balanced AVL tree.
// 30. Write a function to split an AVL tree into two AVL trees based on a value.
// 31. Implement a function to print the AVL tree in level order (breadth-first traversal).
// 32. Write a function to check if an AVL tree contains only unique values.
// 33. Implement a function to remove all leaf nodes from an AVL tree.
// 34. Write a function to find the distance between two nodes in an AVL tree.
// 35. Implement a function to serialize and deserialize an AVL tree.
// 36. Write a function to check if an AVL tree is a complete binary tree.
// 37. Write a function to check if an AVL tree is a perfect binary tree.
// 38. Implement a function to print the boundary nodes of an AVL tree.
// 39. Write a function to find the sum of all nodes at a given depth in an AVL tree.
// 40. Implement a function to find the maximum width of an AVL tree.
// ==============================================================
