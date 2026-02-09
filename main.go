package main

import (
	avl "Practica_GO/DSA/graphs/trees/AVL"
	"fmt"
)

func main() {
	values := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	tree := avl.ArrayToAVL(values)
	k := 2
	kmin, merror := tree.KthSmallestElement(k)
	kmax, Merror := tree.KthLargestElement(k)
	if merror != nil {
		fmt.Println("error at kth smallest element")
		return
	}
	if Merror != nil {
		fmt.Println("error at kth largest element")
		return
	}
	fmt.Println(k, "-th smallest = ", kmin)
	fmt.Println(k, "-th largest = ", kmax)
}
