package cmp

import (
	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Float | constraints.Integer
}

// Compare function compares its two arguments for order. Returns a negative
// integer, zero, or a positive integer as the first argument is less than,
// equal to, or greater than the second.
type Compare[T any] func(a, b T) bool

// CompareNumber function compares its two arguments for order. Returns a negative
// integer, zero, or a positive integer as the first argument is less than,
// equal to, or greater than the second.
type CompareNumber[T Number] func(a, b T) bool

func Eq[T comparable](a, b T) bool {
	return a == b
}

func Less[T Number](a, b T) bool {
	return a < b
}

func Greater[T Number](a, b T) bool {
	return a > b
}

func CmpNumber[T Number](a, b T) bool {
	return a == b
}
