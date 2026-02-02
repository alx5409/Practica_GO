package main

import (
	"Practica_GO/DSA/graphs/trees/avl"
)

func main() {
	// Create an AVL tree of int
	tree := avl.AVLTree[int]{}

	// Insert some values
	values := []int{30, 20, 40, 10, 25, 35, 50, 5, 8, 9}
	for _, v := range values {
		tree.Insert(v)
	}

	diam := tree.Diameter()
	diam++
}
