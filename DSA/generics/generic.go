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

// Generic Map structure
type Map[K comparable, V any] map[K]V
