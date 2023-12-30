package collection

import (
	"collection/arithmetic"
	"collection/cmp"
)

type collector[T any] struct {
	collection []T
}

func (c collector[T]) ToVec() []T {
	return c.collection
}

func (c collector[T]) IsEmpty() bool {
	return len(c.collection) == 0
}

func (c collector[T]) IsNotEmpty() bool {
	return len(c.collection) != 0
}

func (c collector[T]) Len() int {
	return len(c.collection)
}

func (c collector[T]) Clear() Collection[T] {
	//TODO implement me
	panic("implement me")
}

func (c collector[T]) Dedup(cmp cmp.Compare[T]) Collection[T] {
	if len(c.collection) == 0 {
		return c
	}

	j := 0
	for k := 1; k < len(c.collection); k++ {
		if cmp(c.collection[j], c.collection[k]) {
			j++
			c.collection[j] = c.collection[k]
		}
	}
	c.collection = c.collection[:j+1]
	return &c
}

func (c collector[T]) ExtractIf(fn func(T) bool) Collection[T] {
	var result []T
	for _, v := range c.collection {
		if fn(v) {
			result = append(result, v)
		}
	}
	c.collection = result
	return &c
}

func (c collector[T]) First() T {
	if len(c.collection) > 0 {
		return c.collection[0]
	}

	var zero T
	return zero
}

func (c collector[T]) SplitFirst() (T, Collection[T]) {
	var first T
	if len(c.collection) > 0 {
		first = c.collection[0]
		c.collection = c.collection[1:]
	}
	return first, c
}

func (c collector[T]) SplitLast() (T, Collection[T]) {
	var last T
	length := len(c.collection)
	if length > 0 {
		last = c.collection[length-1]
		c.collection = c.collection[:length-1]
	}
	return last, c
}

func (c collector[T]) Last() T {
	if len(c.collection) > 0 {
		return c.collection[len(c.collection)-1]
	}

	var zero T
	return zero
}

func (c collector[T]) Get(j int) T {
	if j >= 0 && j < len(c.collection) {
		return c.collection[j]
	}

	var zero T
	return zero
}

func (c collector[T]) Swap(j int, k int) Collection[T] {
	if j < len(c.collection) && k < len(c.collection) && j >= 0 && k >= 0 {
		c.collection[j], c.collection[k] = c.collection[k], c.collection[j]
	}

	return c
}

func (c collector[T]) Reverse() Collection[T] {
	for left, right := 0, len(c.collection)-1; left < right; left, right = left+1, right-1 {
		c.collection[left], c.collection[right] = c.collection[right], c.collection[left]
	}

	return c
}

func (c collector[T]) Iter() Iterator[T] {
	return &iterator[T]{slice: c.collection}
}

func (c collector[T]) Windows(i int) Split[T] {
	return &windows[T]{slice: c.collection, window: i}
}

func (c collector[T]) GroupBy(fn func(T, T) bool) Split[T] {
	return &GroupBy[T]{slice: c.collection, comparator: fn}
}

func (c collector[T]) SplitAt(j int) (Collection[T], Collection[T]) {
	if j < 0 {
		j = 0
	}
	if j > len(c.collection) {
		j = len(c.collection)
	}

	firstHalf := c.collection[:j]
	secondHalf := c.collection[j:]

	return collector[T]{collection: firstHalf}, collector[T]{collection: secondHalf}
}

func (c collector[T]) Split(fn func(T) bool) Split[T] {
	var result [][]T
	var currentSlice []T

	for _, num := range c.collection {
		if fn(num) {
			result = append(result, currentSlice)
			currentSlice = nil
			continue
		}
		currentSlice = append(currentSlice, num)
	}

	result = append(result, currentSlice)
	return &splitter[T]{slices: result}
}

func (c collector[T]) SplitInclusive(fn func(T) bool) Split[T] {
	var result [][]T
	var currentSlice []T

	for _, item := range c.collection {
		currentSlice = append(currentSlice, item)

		if fn(item) {
			result = append(result, currentSlice)
			currentSlice = []T{item}
		}
	}

	if len(currentSlice) > 0 {
		result = append(result, currentSlice)
	}

	return &splitter[T]{slices: result}
}

func (c collector[T]) Avg(fn arithmetic.AvgFn[T]) T {
	return fn(c.collection)
}

func (c collector[T]) Sum(arithmetic arithmetic.Arithmetic[T]) T {
	var sum T
	for _, item := range c.collection {
		sum = arithmetic(sum, item)
	}
	return sum
}

func (c collector[T]) Min(cmp cmp.Compare[T]) T {
	if len(c.collection) == 0 {
		var zero T
		return zero
	}

	min := c.collection[0]
	for _, item := range c.collection[1:] {
		if cmp(item, min) {
			min = item
		}
	}
	return min
}

func (c collector[T]) Max(cmp cmp.Compare[T]) T {
	if len(c.collection) == 0 {
		var zero T
		return zero
	}

	max := c.collection[0]
	for _, item := range c.collection[1:] {
		if cmp(item, max) {
			max = item
		}
	}
	return max
}
