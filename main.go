package main

import (
	avl "Practica_GO/DSA/graphs/trees/AVL"
	"fmt"
)

func main() {
	values := []int{4, 2, 6, 1, 3, 5, 7}
	tree := avl.ArrayToAVL(values)
	if tree.IsComplete() {
		fmt.Println("complete tree")
		return
	}
	fmt.Println("incomplete tree")
}
