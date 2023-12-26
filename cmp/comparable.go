package cmp

import "cmp"

// Compare function compares its two arguments for order. Returns a negative
// integer, zero, or a positive integer as the first argument is less than,
// equal to, or greater than the second.
type Compare[T any] func(a, b T) bool

func Cmp[T cmp.Ordered](a, b T) bool {
	return a == b
}
