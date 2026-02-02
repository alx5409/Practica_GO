package main

import (
	"Practica_GO/DSA/generics"
	avl "Practica_GO/DSA/graphs/trees/avl"
)

func main() {
	values := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	tree := avl.ArrayToAVL(values)
	sortedArray := tree.ConvertToSortedArray()
	generics.PrintSlice(sortedArray)
}
