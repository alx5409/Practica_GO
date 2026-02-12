package avl

// ===================== AVL TREE EXERCISES =====================
// 1. Implement an AVL tree node structure in Go.

// Adelson-Velsky and Landis generic node
import (
	generics "Practica_GO/DSA/generics"
	binaryTrees "Practica_GO/DSA/graphs/trees/binaryTrees"
	utils "Practica_GO/DSA/utils"
	"errors"
	"fmt"
	"math"
	"strings"
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

func (tree *AVLTree[A]) balanceHelper(node *AVLNode[A]) *AVLNode[A] {
	if node == nil {
		return nil
	}
	node.left = tree.balanceHelper(node.left)
	node.right = tree.balanceHelper(node.right)
	if !node.isBalanced() {
		rType := node.rotationType()
		node = node.rotate(rType)
	}
	return node
}

// AVLBalance checks the balance of the AVL tree and performs the necessary rotations to maintain the AVL property.
func (tree *AVLTree[A]) AVLBalance() {
	tree.Root = tree.balanceHelper(tree.Root)
}

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

func (tree *AVLTree[A]) deleteHelper(node *AVLNode[A], value A) (*AVLNode[A], error) {
	if node == nil {
		return nil, fmt.Errorf("value %v not found", value)
	}
	if value < node.value {
		left, err := tree.deleteHelper(node.left, value)
		if err != nil {
			return nil, err
		}
		node.left = left
	} else if value > node.value {
		right, err := tree.deleteHelper(node.right, value)
		if err != nil {
			return nil, err
		}
		node.right = right
	} else {
		// Node to delete found
		if node.left == nil {
			return node.right, nil
		} else if node.right == nil {
			return node.left, nil
		} else {
			// Two children: find inorder successor (leftmost in right subtree)
			succ := node.right
			for succ.left != nil {
				succ = succ.left
			}
			// Replace value with successor's value
			node.value = succ.value
			// Delete successor
			right, err := tree.deleteHelper(node.right, succ.value)
			if err != nil {
				return nil, err
			}
			node.right = right
		}
	}
	// Rebalance if necessary
	if !node.isBalanced() {
		rType := node.rotationType()
		node = node.rotate(rType)
	}
	return node, nil
}

func (tree *AVLTree[A]) Delete(value A) error {
	root, err := tree.deleteHelper(tree.Root, value)
	if err != nil {
		return err
	}
	tree.Root = root
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

	lcaNode := tree.Root.lcaHelper(value1, value2)
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
	if node1 == nil || node2 == nil {
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

// Converts a generic numeric slice into an avl tree in O(n·log(n)) time complexity.
func ArrayToAVL[N Number](array generics.Slice[N]) AVLTree[N] {
	var avlTree AVLTree[N]
	for _, element := range array {
		avlTree.Insert(element)
	}
	return avlTree
}

// insertSortedHelper recursively builds a balanced AVL tree from a sorted array.
// Returns the root node of the subtree.
func insertSortedHelper[N Number](arr []N, left, right int) *AVLNode[N] {
	if left > right {
		return nil
	}
	mid := (left + right) / 2
	node := &AVLNode[N]{value: arr[mid]}
	node.left = insertSortedHelper[N](arr, left, mid-1)
	node.right = insertSortedHelper[N](arr, mid+1, right)
	return node
}

// InsertSorted builds the AVL tree from a sorted array in O(n) time.
func (tree *AVLTree[N]) InsertSorted(sortedArray []N) {
	if len(sortedArray) == 0 {
		tree.Root = nil
		return
	}
	tree.Root = insertSortedHelper(sortedArray, 0, len(sortedArray)-1)
}

// Converts a sorted array into an AVL tree efficiently, in O(n) time complexity.
// If the array is not sorted returns an empty AVL tree and an error.
func SortedArrayToAVL[N Number](sortedArray generics.ComparableSlice[N]) (AVLTree[N], error) {
	if !sortedArray.IsOrdered() {
		return AVLTree[N]{Root: nil}, errors.New("empty array")
	}
	if len(sortedArray) == 0 {
		return AVLTree[N]{Root: nil}, nil
	}
	var avlTree AVLTree[N]
	avlTree.InsertSorted(sortedArray)
	return avlTree, nil
}

// 24. Write a function to convert an AVL tree to a sorted array (inorder traversal).

func (node *AVLNode[N]) convertToSortedArrayHelper(currentArray *[]N) {
	if node == nil {
		return
	}
	node.left.convertToSortedArrayHelper(currentArray)
	*currentArray = append(*currentArray, node.value)
	node.right.convertToSortedArrayHelper(currentArray)
}

// Inorder traverse the AVL tree to return a sorted array
func (tree AVLTree[N]) ConvertToSortedArray() []N {
	var sortedArray []N
	tree.Root.convertToSortedArrayHelper(&sortedArray)
	return sortedArray
}

// 25. Implement a function to find the predecessor and successor of a given value in an AVL tree.

func (node *AVLNode[N]) predecessorHelper(value N) (*AVLNode[N], error) {
	target := node.dfsHelper(value)
	if target == nil {
		return nil, fmt.Errorf("nil node")
	}
	// Case 1: : predecessor is the rightmost node in left subtree
	if target.left != nil {
		pred := target.left
		for pred.right != nil {
			pred = pred.right
		}
		return pred, nil
	}
	// Case 2: predecessor is the last ancestor where we moved right
	var pred *AVLNode[N]
	curr := node
	for curr != nil {
		if value > curr.value {
			pred = curr
			curr = curr.right
		} else if value < curr.value {
			curr = curr.left
		} else {
			break
		}
	}
	return pred, nil
}

// Returns the in-order predecessor of the given value: the largest value in the tree that is less than the given value.
// If the value does not exist in the tree or has no predecessor, an error is returned
func (tree AVLTree[N]) Predecessor(value N) (N, error) {
	var zero N
	predNode, err := tree.Root.predecessorHelper(value)
	if err != nil {
		return zero, fmt.Errorf("no predecessor for value %v: ", value)
	}
	return predNode.value, nil
}

func (node *AVLNode[N]) ancestorHelper(value N) (*AVLNode[N], error) {
	target := node.dfsHelper(value)
	if target == nil {
		return nil, fmt.Errorf("value %v not found", value)
	}
	// Case 1: successor is the leftmost node in right subtree
	if target.right != nil {
		succ := target.right
		for succ.left != nil {
			succ = succ.left
		}
		return succ, nil
	}
	// Case 2: successor is the last ancestor where we moved left
	var succ *AVLNode[N]
	curr := node
	for curr != nil {
		if value < curr.value {
			succ = curr
			curr = curr.left
		} else if value > curr.value {
			curr = curr.right
		} else {
			break
		}
	}
	if succ == nil {
		return nil, fmt.Errorf("no successor for value %v", value)
	}
	return succ, nil
}

// Returns the in-order ancestor of the given value: the lowest value in the tree that is greater than the given value.
// / If the value does not exist in the tree or has no ancestor, an error is returned
func (tree AVLTree[N]) Ancestor(value N) (N, error) {
	var zero N
	ancestorNode, err := tree.Root.ancestorHelper(value)
	if err != nil {
		return zero, fmt.Errorf("no predecessor for value %v: ", value)
	}
	return ancestorNode.value, nil
}

// 26. Write a function to find the kth smallest element in an AVL tree.

func inorderMinCollect[N Number](node *AVLNode[N], result *[]N, maxLength int) {
	if node == nil || len(*result) >= maxLength {
		return
	}
	inorderMinCollect(node.left, result, maxLength)
	*result = append(*result, node.value)
	inorderMinCollect(node.right, result, maxLength)
}

// Returns the kth smallest element in the AVL tree.
// If the k is greater than the size of the tree returns an error.
func (tree AVLTree[N]) KthSmallestElement(k int) (N, error) {
	var zero N
	if k <= 0 {
		return zero, fmt.Errorf("k must be positive")
	}
	var result []N
	inorderMinCollect(tree.Root, &result, k)
	if k > len(result) {
		return zero, fmt.Errorf("k (%d) exceeds tree size (%d)", k, len(result))
	}
	return result[k-1], nil
}

// 27. Write a function to find the kth largest element in an AVL tree.

func inorderMaxCollect[N Number](node *AVLNode[N], result *[]N, maxLength int) {
	if node == nil || len(*result) >= maxLength {
		return
	}
	inorderMaxCollect(node.right, result, maxLength)
	*result = append(*result, node.value)
	inorderMaxCollect(node.left, result, maxLength)
}

func (tree AVLTree[N]) KthLargestElement(k int) (N, error) {
	var zero N
	if k <= 0 {
		return zero, fmt.Errorf("k must be positive")
	}
	var result []N
	inorderMaxCollect(tree.Root, &result, k)
	if k > len(result) {
		return zero, fmt.Errorf("k (%d) exceeds tree size (%d)", k, len(result))
	}
	return result[k-1], nil
}

// 28. Implement a function to clone (deep copy) an AVL tree.

func (node *AVLNode[N]) cloneNode() *AVLNode[N] {
	if node == nil {
		return nil
	}
	copiedNode := &AVLNode[N]{value: node.value}
	copiedNode.left = node.left.cloneNode()
	copiedNode.right = node.right.cloneNode()
	return copiedNode
}

func (tree AVLTree[N]) Clone() AVLTree[N] {
	var treeCopy AVLTree[N]
	treeCopy.Root = tree.Root.cloneNode()
	return treeCopy
}

// 29. Write a function to merge two AVL trees into a single balanced AVL tree.

// Returns an AVL tree which correspond to the merge of the two AVL trees.
func MergeAVLTrees[N Number](tree1, tree2 AVLTree[N]) AVLTree[N] {
	// convert the trees into sorted array, merge them and then build the new tree
	sortedSlice1 := tree1.ConvertToSortedArray()
	sortedSlice2 := tree2.ConvertToSortedArray()
	// Cast to ComparableSlice[N] to use OrderedMerge
	compSlice1 := generics.ComparableSlice[N](sortedSlice1)
	compSlice2 := generics.ComparableSlice[N](sortedSlice2)
	mergedSlice := compSlice1.OrderedMerge(compSlice2)
	mergedTree, _ := SortedArrayToAVL(mergedSlice)
	return mergedTree
}

// 30. Write a function to split an AVL tree into two AVL trees based on a value.

// splits the current AVL tree into two AVL trees where the first AVL tree consists of all
// the nodes which are less than value and the second AVL tree consists of all
// the nodes which are greater than or equal to value.
func (tree AVLTree[N]) Split(value N) (AVLTree[N], AVLTree[N], error) {
	if tree.IsEmpty() {
		return tree, tree, errors.New("empty tree")
	}
	var nilTree AVLTree[N]
	nodeValue := tree.Root.dfsHelper(value)
	if nodeValue == nil {
		return nilTree, nilTree, errors.New("value not found")
	}
	// converto to sorted array to find where to split
	sortedArray := tree.ConvertToSortedArray()
	index := -1
	for i, v := range sortedArray {
		if v == value {
			index = i
			break
		}
	}
	lessArray := sortedArray[:index]
	greaterArray := sortedArray[index:]
	leftTree, _ := SortedArrayToAVL(generics.ComparableSlice[N](lessArray))
	rightTree, _ := SortedArrayToAVL(generics.ComparableSlice[N](greaterArray))
	return leftTree, rightTree, nil
}

// 31. Implement a function to print the AVL tree in level order (breadth-first traversal).

// prints the node values in level order
func (tree AVLTree[N]) PrintInLevel() {
	if tree.IsEmpty() {
		return
	}
	queue := generics.Queue[*AVLNode[N]]{}
	queue.Enqueue(tree.Root)
	for !queue.IsEmpty() {
		levelSize := len(queue.Data)
		for i := 0; i < levelSize; i++ {
			node, _ := queue.Dequeue()
			fmt.Printf("%v ", node.value)
			if node.left != nil {
				queue.Enqueue(node.left)
			}
			if node.right != nil {
				queue.Enqueue(node.right)
			}
		}
		fmt.Println()
	}
}

// 32. Write a function to check if an AVL tree contains only unique values.

func (node *AVLNode[N]) uniqueValuesHelper(seenMap map[N]bool) bool {
	if node == nil {
		return true
	}
	if seenMap[node.value] {
		return false
	}
	seenMap[node.value] = true
	return node.left.uniqueValuesHelper(seenMap) && node.right.uniqueValuesHelper(seenMap)
}

func (tree AVLTree[N]) HasUniqueValues() bool {
	if tree.IsEmpty() {
		return true
	}
	seenMap := make(map[N]bool)
	return tree.Root.uniqueValuesHelper(seenMap)
}

// 33. Implement a function to remove all leaf nodes from an AVL tree.

func (node *AVLNode[N]) removeLeafsHelper() *AVLNode[N] {
	if node == nil {
		return nil
	}
	if node.isLeaf() {
		return nil
	}
	node.left = node.left.removeLeafsHelper()
	node.right = node.right.removeLeafsHelper()
	if !node.isBalanced() {
		rType := node.rotationType()
		node = node.rotate(rType)
	}
	return node
}

// Removes all leafs nodes in the tree
func (tree *AVLTree[N]) RemoveAllLeafs() {
	tree.Root = tree.Root.removeLeafsHelper()
}

// 34. Write a function to find the distance between two nodes in an AVL tree.

func nodeDepthHelper[N Number](current *AVLNode[N], target *AVLNode[N], depth int) int {
	if current == nil {
		return -1
	}
	if current == target {
		return depth
	}
	if target.value < current.value {
		return nodeDepthHelper(current.left, target, depth+1)
	} else {
		return nodeDepthHelper(current.right, target, depth+1)
	}
}

// returns the depth of the node in the avl tree starting from the root
func (tree *AVLTree[N]) nodeDepth(node *AVLNode[N]) (int, error) {
	return nodeDepthHelper(tree.Root, node, 0), nil
}

func (tree *AVLTree[N]) NodeDistance(nodeValue1, nodeValue2 N) (int, error) {
	node1 := tree.Root.dfsHelper(nodeValue1)
	if node1 == nil {
		return 0, fmt.Errorf("nodeValue: %v not found", nodeValue1)
	}
	node2 := tree.Root.dfsHelper(nodeValue2)
	if node2 == nil {
		return 0, fmt.Errorf("nodeValue: %v not found", nodeValue2)
	}
	// find the LCA of both nodes
	lcaValue, err := tree.LCA(nodeValue1, nodeValue2)
	if err != nil {
		return 0, err
	}
	lcaNode := tree.Root.dfsHelper(lcaValue)

	// compute the depths to find the distance
	node1Depth, _ := tree.nodeDepth(node1)
	node2Depth, _ := tree.nodeDepth(node2)
	lcaDepth, _ := tree.nodeDepth(lcaNode)
	return node1Depth + node2Depth - 2*lcaDepth, nil
}

// 35. Implement a function to serialize and deserialize an AVL tree.

func serializeHelper[N Number](node *AVLNode[N], result *[]string) {
	if node == nil {
		*result = append(*result, "nil")
		return
	}
	*result = append(*result, fmt.Sprintf("%v", node.value))
	serializeHelper(node.left, result)
	serializeHelper(node.right, result)
}

// converts the AVL tree to a string using preorder traversal.
// null nodes are represented as "nil".
func (tree AVLTree[N]) Serialize() string {
	var result []string
	serializeHelper(tree.Root, &result)
	return strings.Join(result, " ")
}

func deserializeHelper[N Number](index *int, tokens []string) *AVLNode[N] {
	if *index >= len(tokens) || tokens[*index] == "nil" {
		*index++
		return nil
	}
	var value N
	fmt.Sscanf(tokens[*index], "%v", &value)
	*index++
	node := &AVLNode[N]{value: value}
	node.left = deserializeHelper[N](index, tokens)
	node.right = deserializeHelper[N](index, tokens)
	return node
}

// reconstructs an AVL tree from a serialized string.
// the string should be space-separeted values from preorder traversal.
func Deserialize[N Number](data string) AVLTree[N] {
	tokens := strings.Fields(data)
	index := 0
	root := deserializeHelper[N](&index, tokens)
	return AVLTree[N]{Root: root}
}

// 36. Write a function to check if an AVL tree is a complete binary tree.

// checkNodeForCompleteness validates a single node based on the completeness flag.
func (tree AVLTree[N]) checkNodeForCompleteness(node *AVLNode[N], hasSeenMissingChild bool) (bool, bool) {
	updatedFlag := hasSeenMissingChild

	// Check left child
	if node.left != nil {
		if updatedFlag {
			return false, false
		}
	} else {
		updatedFlag = true
	}

	// Check right child
	if node.right != nil {
		if updatedFlag {
			return false, false // Invalid: child after missing one
		}
	} else {
		updatedFlag = true
	}

	return true, updatedFlag
}

// processLevel handles one level of BFS, checking nodes and updating the flag.
func (tree AVLTree[N]) processLevel(queue []*AVLNode[N], hasSeenMissingChild bool) (bool, []*AVLNode[N], bool) {
	newQueue := []*AVLNode[N]{}
	updatedFlag := hasSeenMissingChild

	for _, node := range queue {
		isValid, tempFlag := tree.checkNodeForCompleteness(node, updatedFlag)
		if !isValid {
			return false, nil, false
		}
		updatedFlag = tempFlag

		// Enqueue children for next level
		if node.left != nil {
			newQueue = append(newQueue, node.left)
		}
		if node.right != nil {
			newQueue = append(newQueue, node.right)
		}
	}
	return true, newQueue, updatedFlag
}

// checks if the AVL tree is complete: a complete binary tree has all levels fully filled except possibly the last,
// with nodes packed to the left.
func (tree AVLTree[N]) IsComplete() bool {
	if tree.Root == nil {
		return true
	}

	queue := []*AVLNode[N]{tree.Root}
	hasSeenMissingChild := false

	for len(queue) > 0 {
		ok, newQueue, updatedFlag := tree.processLevel(queue, hasSeenMissingChild)
		if !ok {
			return false
		}
		queue = newQueue
		hasSeenMissingChild = updatedFlag
	}
	return true
}

// 37. Write a function to check if an AVL tree is a perfect binary tree:

// checks if the AVL tree is perfect: every level of the tree is full.
func (tree AVLTree[N]) IsPerfect() bool {
	if tree.Root == nil {
		return true
	}
	height := tree.Root.subtreeHeight()
	size := tree.Size()
	// since every level is full there are 1 + 2 + 4 + ... + 2^{height - 1} = 2^{height} - 1
	expectedSize := int(math.Pow(2, float64(height))) - 1
	return size == expectedSize
}

// 38. Implement a function to print the boundary nodes of an AVL tree.

// prints every boundary node values: the root node, the left boundary, the right boundary and the leafs
func (tree AVLTree[N]) PrintBoundaryNodes() {
	if tree.IsEmpty() {
		return
	}

	var boundary []N
	boundary = append(boundary, tree.Root.value)
	tree.addLeftBoundary(tree.Root.left, &boundary)
	tree.addLeaves(tree.Root, &boundary)
	tree.addRightBoundary(tree.Root.right, &boundary)

	for _, v := range boundary {
		fmt.Printf("%v ", v)
	}
	fmt.Println()
}

func (tree AVLTree[N]) addLeftBoundary(node *AVLNode[N], boundary *[]N) {
	if node == nil || node.isLeaf() {
		return
	}
	*boundary = append(*boundary, node.value)
	if node.left != nil {
		tree.addLeftBoundary(node.left, boundary)
	} else {
		tree.addLeftBoundary(node.right, boundary)
	}
}

func (tree AVLTree[N]) addLeaves(node *AVLNode[N], boundary *[]N) {
	if node == nil {
		return
	}
	if node.isLeaf() {
		*boundary = append(*boundary, node.value)
		return
	}
	tree.addLeaves(node.left, boundary)
	tree.addLeaves(node.right, boundary)
}

func (tree AVLTree[N]) addRightBoundary(node *AVLNode[N], boundary *[]N) {
	if node == nil || node.isLeaf() {
		return
	}
	if node.right != nil {
		tree.addRightBoundary(node.right, boundary)
	} else {
		tree.addRightBoundary(node.left, boundary)
	}
	*boundary = append(*boundary, node.value)
}

// 39. Write a function to find the sum of all nodes at a given depth in an AVL tree.

// helper function to sum every node at a certain height
func (node *AVLNode[N]) preSumTravNodes(level int, current int, sum *N) {
	if node == nil {
		return
	}
	if level == current {
		*sum += node.value
		return
	}
	node.left.preSumTravNodes(level, current+1, sum)
	node.right.preSumTravNodes(level, current+1, sum)
}

// Sums every node at the specified height, if the chosen height is empty returns an error
func (tree AVLTree[N]) SumNodesAtHeight(height int) (N, error) {
	var zero N
	if height < 0 {
		return zero, errors.New("height must be a positive integer")
	}
	if tree.IsEmpty() {
		return zero, errors.New("tree is empty")
	}
	if height >= tree.Root.subtreeHeight() {
		return zero, errors.New("the chosen height is greater than the tree height")
	}
	var sum N
	current := 0
	tree.Root.preSumTravNodes(height, current, &sum)
	return sum, nil
}

// 40. Implement a function to find the maximum width of an AVL tree.

// returns the maximum number of nodes at any level in the AVL tree.
func (tree AVLTree[N]) MaxWidth() int {
	if tree.Root == nil {
		return 0
	}
	maxWidth := 0
	queue := []*AVLNode[N]{tree.Root}
	// bfs to find the maximum width
	for len(queue) > 0 {
		levelSize := len(queue)
		maxWidth = max(maxWidth, levelSize)
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.left != nil {
				queue = append(queue, node.left)
			}
			if node.right != nil {
				queue = append(queue, node.right)
			}
		}
	}
	return maxWidth
}
