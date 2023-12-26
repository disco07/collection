package collection

import "collection/cmp"

type Collection[T any] interface {
	// ToVec returns a Slice containing all the elements of this Collection.
	ToVec() []T

	// IsEmpty returns true if the value is empty; otherwise, false is returned.
	IsEmpty() bool

	// IsNotEmpty returns true if the value is not empty; otherwise, false is returned.
	IsNotEmpty() bool

	// Len return the length of the value.
	Len() int

	// Clear removes all values of the value
	Clear() Collection[T]

	// Dedup Removes consecutive repeated elements in the value
	Dedup(cmp cmp.Compare[T]) Collection[T]

	// ExtractIf Creates an iterator which uses a closure to determine if an element should be removed
	ExtractIf(fn func(T) bool) Collection[T]

	// First Returns the first element of the slice, or nil if it is empty.
	First() T

	// SplitFirst Returns the first and all the rest of the elements of the slice, or nil if it is empty.
	SplitFirst() (T, Collection[T])

	// SplitLast Returns the last and all the rest of the elements of the slice, or nil if it is empty.
	SplitLast() (T, Collection[T])

	// Last Returns the last element of the slice, or nil if it is empty.
	Last() T

	// Get Returns a reference to an element or subslice depending on the type of index, or nil if it is empty.
	Get(int) T

	// Swap Swaps two elements in the slice by index.
	Swap(int, int) Collection[T]

	// Reverse reverses the order of the value's items, preserving the original keys.
	Reverse() Collection[T]

	// Iter Returns an iterator over the slice.
	Iter() Iterator

	// Windows Returns an iterator over all contiguous windows of length size. The windows overlap.
	// If the slice is shorter than size, the iterator returns no values.
	Windows() Iterator

	// GroupBy Returns an iterator over the slice producing non-overlapping runs of elements using the predicate to separate them.
	GroupBy(fn func() bool) Iterator

	// SplitAt Divides one slice into two at an index.
	SplitAt(int) (Collection[T], Collection[T])

	// Split Returns an iterator over subslices separated by elements that match pred. The matched element is not contained in the subslices.
	Split(fn func() bool) Split

	// SplitInclusive Returns an iterator over subslices separated by elements that match pred.
	// The matched element is contained in the end of the previous subslice as a terminator.
	SplitInclusive(fn func() bool) Split

	// Avg returns the average value of a given key.
	Avg() T

	// Sum returns the sum of all items in the value.
	Sum() T

	// Min returns the minimum value of a given key.
	Min() T

	// Max returns the maximum value of a given key.
	Max() T
}

type Iterator interface {
	Next() (any, bool)
}

type Split interface {
	Next() (any, bool)
}

func NewVec[T any](vec []T) Collection[T] {
	return &iterator[T]{
		collection: vec,
	}
}
