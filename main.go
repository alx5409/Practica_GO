package main

import (
	avl "Practica_GO/DSA/graphs/trees/avl"
	"fmt"
)

func main() {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	tree, err := avl.SortedArrayToAVL(values)
	if err != nil {
		fmt.Println("error at creating the tree")
	}
	tree.InorderTraversal()
}
