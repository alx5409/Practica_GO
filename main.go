package main

import (
	avl "Practica_GO/DSA/graphs/trees/avl"
	"fmt"
)

func main() {
	values := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	tree := avl.ArrayToAVL(values)
	value := 7
	pred, err := tree.Ancestor(value)
	if err != nil {
		fmt.Println("error")
	}
	fmt.Println("The ancestor is ", pred)
}
