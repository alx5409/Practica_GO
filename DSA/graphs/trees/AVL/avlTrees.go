package avl

// ===================== AVL TREE EXERCISES =====================
// 1. Implement an AVL tree node structure in Go.

// Adelson-Velsky and Landis generic node
import (
	utils "Practica_GO/DSA/utils"
)

type Number utils.Number

type AVLNode[A Number] struct {
	value A
	left  *AVLNode[A]
	right *AVLNode[A]
}

// Adelson-Velsky and Landis generic tree
type AVLTree[A Number] struct {
	Root *AVLNode[A]
}

func (a AVLTree[A]) isEmpty() bool {
	return a.Root == nil
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
	return node.left.subtreeHeight() - node.right.subtreeHeight()
}

// Checks if a node is unbalanced, meaning, its balance factor is outside of {-1, 0, 1} values.
func (node *AVLNode[A]) isBalanced() bool {
	balanceFactor := node.balanceFactor()
	return balanceFactor == -1 || balanceFactor == 0 || balanceFactor == 1
}

// 2. Write a function to insert a value into an AVL tree and maintain balance.

// Perform a left-left rotation (right rotation)
func (n *AVLNode[A]) LLRotation() *AVLNode[A] {
	leftChild := n.left
	n.left = leftChild.right
	leftChild.right = n
	return leftChild
}

// Perform a right-right rotation (left rotation)
func (n *AVLNode[A]) RRRotation() *AVLNode[A] {
	rightChild := n.right
	n.right = rightChild.left
	rightChild.left = n
	return rightChild
}

// Perform a left-right rotation
func (n *AVLNode[A]) LRRotation() *AVLNode[A] {
	n.left = n.left.RRRotation()
	return n.LLRotation()
}

// Perform a right-left rotation
func (n *AVLNode[A]) RLRotation() *AVLNode[A] {
	n.right = n.right.LLRotation()
	return n.RRRotation()
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
func (n *AVLNode[A]) rotate(rotationType string) {
	switch rotationType {
	case "LL":
		n.LLRotation()
	case "RR":
		n.RRRotation()
	case "RL":
		n.RLRotation()
	case "LR":
		n.LRRotation()
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
func (a AVLTree[A]) insertHelper(node *AVLNode[A], value A) *AVLNode[A] {
	if node == nil {
		return &AVLNode[A]{value: value}
	}
	if value < node.value {
		node.left = a.insertHelper(node.left, value)
	} else {
		node.right = a.insertHelper(node.right, value)
	}
	// if the balance factor is okay just return the node
	if node.isBalanced() {
		return node
	}
	// rebalance the node with a rotation
	rType := node.rotationType()
	node.rotate(rType)
	return node
}

// Function that inserts a value in the AVL tree keeping the structure of the tree
func (a AVLTree[A]) insert(value A) {
	a.Root = a.insertHelper(a.Root, value)
}

// 3. Write a function to delete a value from an AVL tree and maintain balance.

func (a AVLTree[A]) delete(value A) error {

	return nil
}

// 4. Implement a function to perform a left rotation on a subtree.
// 5. Implement a function to perform a right rotation on a subtree.
// 6. Write a function to get the height of a node in an AVL tree.
// 7. Write a function to get the balance factor of a node in an AVL tree.
// 8. Implement a function to search for a value in an AVL tree.
// 9. Write a function to find the minimum value in an AVL tree.
// 10. Write a function to find the maximum value in an AVL tree.
// 11. Implement preorder, inorder, and postorder traversals for an AVL tree.
// 12. Write a function to check if a given binary tree is a valid AVL tree.
// 13. Write a function to print all nodes at a given level in an AVL tree.
// 14. Implement a function to count the number of nodes in an AVL tree.
// 15. Write a function to count the number of leaf nodes in an AVL tree.
// 16. Write a function to calculate the sum of all node values in an AVL tree.
// 17. Implement a function to find the lowest common ancestor (LCA) of two nodes in an AVL tree.
// 18. Write a function to mirror (invert) an AVL tree.
// 19. Write a function to print all root-to-leaf paths in an AVL tree.
// 20. Implement a function to check if two AVL trees are identical.
// 21. Write a function to find the diameter (longest path) of an AVL tree.
// 22. Write a function to check if an AVL tree is balanced at every node.
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
