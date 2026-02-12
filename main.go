package main

import (
	avl "Practica_GO/DSA/graphs/trees/AVL"
)

func main() {
	values := []int{4, 2, 6, 1, 3, 5, 7, 8, 9, 10}
	tree := avl.ArrayToAVL(values)
	tree.PrintBoundaryNodes()
}
