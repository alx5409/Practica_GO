package main

import (
	avl "Practica_GO/DSA/graphs/trees/avl"
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

	sum := tree.SumNodes()
	fmt.Println("The sum of the nodes is ", sum)
}
