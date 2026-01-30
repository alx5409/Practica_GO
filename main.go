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

	sum := tree.SumNodes()
	fmt.Println("The sum of the nodes is ", sum)
	lca, err := tree.LCA(10, 50)
	if err != nil {
		fmt.Println("there is no lca")
	}
	fmt.Println("the lca is ", lca)
}
