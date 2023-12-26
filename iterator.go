package collection

import "collection/cmp"

type iterator[T any] struct {
	collection []T
}

func (i iterator[T]) ToVec() []T {
	return i.collection
}

func (i iterator[T]) IsEmpty() bool {
	return len(i.collection) == 0
}

func (i iterator[T]) IsNotEmpty() bool {
	return len(i.collection) != 0
}

func (i iterator[T]) Len() int {
	return len(i.collection)
}

func (i iterator[T]) Clear() Collection[T] {
	//TODO implement me
	panic("implement me")
}

func (i iterator[T]) Dedup(cmp cmp.Compare[T]) Collection[T] {
	if len(i.collection) == 0 {
		return i
	}

	j := 0
	for k := 1; k < len(i.collection); k++ {
		if !cmp(i.collection[j], i.collection[k]) {
			j++
			i.collection[j] = i.collection[k]
		}
	}
	i.collection = i.collection[:j+1]
	return &i
}

func (i iterator[T]) ExtractIf(fn func(T) bool) Collection[T] {
	var result []T
	for _, v := range i.collection {
		if fn(v) {
			result = append(result, v)
		}
	}
	i.collection = result
	return &i
}

func (i iterator[T]) First() T {
	if len(i.collection) > 0 {
		return i.collection[0]
	}

	var zero T
	return zero
}

func (i iterator[T]) SplitFirst() (T, Collection[T]) {
	//TODO implement me
	panic("implement me")
}

func (i iterator[T]) SplitLast() (T, Collection[T]) {
	//TODO implement me
	panic("implement me")
}

func (i iterator[T]) Last() T {
	if len(i.collection) > 0 {
		return i.collection[len(i.collection)-1]
	}

	var zero T
	return zero
}

func (i iterator[T]) Get(i2 int) T {
	//TODO implement me
	panic("implement me")
}

func (i iterator[T]) Swap(j int, k int) Collection[T] {
	if j < len(i.collection) && k < len(i.collection) && j >= 0 && k >= 0 {
		i.collection[j], i.collection[k] = i.collection[k], i.collection[j]
	}

	return i
}

func (i iterator[T]) Reverse() Collection[T] {
	for left, right := 0, len(i.collection)-1; left < right; left, right = left+1, right-1 {
		i.collection[left], i.collection[right] = i.collection[right], i.collection[left]
	}

	return i
}

func (i iterator[T]) Iter() Iterator {
	//TODO implement me
	panic("implement me")
}

func (i iterator[T]) Windows() Iterator {
	//TODO implement me
	panic("implement me")
}

func (i iterator[T]) GroupBy(fn func() bool) Iterator {
	//TODO implement me
	panic("implement me")
}

func (i iterator[T]) SplitAt(i2 int) (Collection[T], Collection[T]) {
	//TODO implement me
	panic("implement me")
}

func (i iterator[T]) Split(fn func() bool) Split {
	//TODO implement me
	panic("implement me")
}

func (i iterator[T]) SplitInclusive(fn func() bool) Split {
	//TODO implement me
	panic("implement me")
}

func (i iterator[T]) Avg() T {
	//TODO implement me
	panic("implement me")
}

func (i iterator[T]) Sum() T {
	//TODO implement me
	panic("implement me")
}

func (i iterator[T]) Min() T {
	//TODO implement me
	panic("implement me")
}

func (i iterator[T]) Max() T {
	//TODO implement me
	panic("implement me")
}
