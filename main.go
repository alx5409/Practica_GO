package main

import (
	avl "Practica_GO/DSA/graphs/trees/AVL"
	"fmt"
)

func main() {
	values := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0, 10, 16, 13}
	tree := avl.ArrayToAVL(values)
	d, err := tree.NodeDistance(3, 8)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("The distance between the node values is: ", d)
}
