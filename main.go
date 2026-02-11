package main

import (
	avl "Practica_GO/DSA/graphs/trees/AVL"
	"fmt"
	// "fmt"
)

func main() {
	values := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0, 10, 16, 13}
	tree := avl.ArrayToAVL(values)
	unique := tree.HasUniqueValues()
	fmt.Println(unique)
}
