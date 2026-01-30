package main

import (
	"Practica_GO/DSA/graphs/trees/avl"
	"fmt"
)

func main() {
	// Create an AVL tree of int
	tree := avl.AVLTree[int]{}

	// Insert some values
	values := []int{30, 20, 40, 10, 25, 35, 50}
	for _, v := range values {
		tree.Insert(v)
	}

	if isBalanced := tree.IsBalanced(); isBalanced {
		fmt.Println("tree is balanced")
	} else {
		fmt.Println("tree is unbalanced")
	}
}
