package arithmetic

import "collection/cmp"

type Arithmetic[T any] func(a, b T) T

type AvgFn[T any] func([]T) T

func Sum[T cmp.Number](a, b T) T {
	return a + b

}

func Avg[T cmp.Number](arr []T) T {
	var sum T
	for _, item := range arr {
		sum += item
	}
	return sum / T(len(arr))
}
