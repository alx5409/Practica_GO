package generics

// Generic Pair structure
type Pair[A any] struct {
	First  A
	Second A
}

// Generic number interface
type Number interface {
	~int | ~int32 | ~int64 | ~float32 | ~float64
}

// Generic slice
type Slice[T any] []T

// Generic hashmap
type HashMap[K comparable, V any] map[K]V

// Generic stack
type Stack[V any] struct {
	data []V
}

// Generic queue
type Queue[V any] struct {
	data []V
}

// Generic linked list node
type ListNode[T any] struct {
	Value T
	Next  *ListNode[T]
}

// Generic linked list
type LinkedList[T any] struct {
	Head *ListNode[T]
}

// Generic binary tree node
type TreeNode[T any] struct {
	Value T
	Left  *TreeNode[T]
	Right *TreeNode[T]
}

// Generic binary tree
type BinaryTree[T any] struct {
	Root *TreeNode[T]
}

// Generic set
type Set[T comparable] map[T]struct{} // Using empty struct to save memory (each value occupies 0 bytes)
