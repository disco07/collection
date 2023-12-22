package collections

import (
	"reflect"
)

type Collection interface {
	Value() any

	// All returns the underlying array represented by the value.
	All() []any

	// IsEmpty returns true if the value is empty; otherwise, false is returned.
	IsEmpty() bool

	// IsNotEmpty returns true if the value is not empty; otherwise, false is returned.
	IsNotEmpty() bool

	// Len return the length of the value.
	Len() int

	// Clear removes all values of the value
	Clear() Collection

	// Dedup Removes consecutive repeated elements in the value
	Dedup() Collection

	// ExtractIf Creates an iterator which uses a closure to determine if an element should be removed
	ExtractIf(fn func()) Collection

	// Flatten Takes a [[T; N]], and flattens it to a [T].
	Flatten() Collection

	// First Returns the first element of the slice, or nil if it is empty.
	First() any

	// SplitFirst Returns the first and all the rest of the elements of the slice, or nil if it is empty.
	SplitFirst() (any, Collection)

	// SplitLast Returns the last and all the rest of the elements of the slice, or nil if it is empty.
	SplitLast() (any, Collection)

	// Last Returns the last element of the slice, or nil if it is empty.
	Last() any

	// Get Returns a reference to an element or subslice depending on the type of index, or nil if it is empty.
	Get(int) any

	// Swap Swaps two elements in the slice by index.
	Swap(int, int) Collection

	// Reverse reverses the order of the value's items, preserving the original keys.
	Reverse() Collection

	// Iter Returns an iterator over the slice.
	Iter() Iterator

	// Windows Returns an iterator over all contiguous windows of length size. The windows overlap.
	// If the slice is shorter than size, the iterator returns no values.
	Windows() Iterator

	// GroupBy Returns an iterator over the slice producing non-overlapping runs of elements using the predicate to separate them.
	GroupBy(func()) Iterator

	// SplitAt Divides one slice into two at an index.
	SplitAt(int) (Collection, Collection)

	// Split Returns an iterator over subslices separated by elements that match pred. The matched element is not contained in the subslices.
	Split(func()) Split

	// SplitInclusive Returns an iterator over subslices separated by elements that match pred.
	// The matched element is contained in the end of the previous subslice as a terminator.
	SplitInclusive(func()) Split

	// Avg returns the average value of a given key.
	Avg(key ...string) any

	// Sum returns the sum of all items in the value.
	Sum(key ...string) any

	// Min returns the minimum value of a given key.
	Min(key ...string) any

	// Max returns the maximum value of a given key.
	Max(key ...string) any
}

type Iterator interface {
	Next() (any, bool)
}

type Split interface {
	Next() (any, bool)
}

func Collect(arr any) Collection {
	t := reflect.TypeOf(arr)
	rv := reflect.ValueOf(arr)

	switch t.Kind() {
	case reflect.Slice:
		// Pour une slice, déterminer le type des éléments de la slice.
		elemType := t.Elem().Kind()
		switch elemType {
		case reflect.Int:
			var num = make([]Number[int], rv.Len())
			main := MainCollection{
				collection: arr,
			}
			for k, v := range arr.([]int) {
				num[k] = NewNumber(v)
			}
			return NumberCollection[int]{
				MainCollection: main,
				value:          num,
			}
		case reflect.Int8:
			var num = make([]Number[int8], rv.Len())
			main := MainCollection{
				collection: arr,
			}
			for k, v := range arr.([]int8) {
				num[k] = NewNumber(v)
			}
			return NumberCollection[int8]{
				MainCollection: main,
				value:          num,
			}
		case reflect.Int16:
			var num = make([]Number[int16], rv.Len())
			main := MainCollection{
				collection: arr,
			}
			for k, v := range arr.([]int16) {
				num[k] = NewNumber(v)
			}
			return NumberCollection[int16]{
				MainCollection: main,
				value:          num,
			}
		case reflect.Int32:
			var num = make([]Number[int32], rv.Len())
			main := MainCollection{
				collection: arr,
			}
			for k, v := range arr.([]int32) {
				num[k] = NewNumber(v)
			}
			return NumberCollection[int32]{
				MainCollection: main,
				value:          num,
			}
		case reflect.Int64:
			var num = make([]Number[int64], rv.Len())
			main := MainCollection{
				collection: arr,
			}
			for k, v := range arr.([]int64) {
				num[k] = NewNumber(v)
			}
			return NumberCollection[int64]{
				MainCollection: main,
				value:          num,
			}
		case reflect.Uint:
			var num = make([]Number[uint], rv.Len())
			main := MainCollection{
				collection: arr,
			}
			for k, v := range arr.([]uint) {
				num[k] = NewNumber(v)
			}
			return NumberCollection[uint]{
				MainCollection: main,
				value:          num,
			}
		case reflect.Uint8:
			var num = make([]Number[uint8], rv.Len())
			main := MainCollection{
				collection: arr,
			}
			for k, v := range arr.([]uint8) {
				num[k] = NewNumber(v)
			}
			return NumberCollection[uint8]{
				MainCollection: main,
				value:          num,
			}
		case reflect.Uint16:
			var num = make([]Number[uint16], rv.Len())
			main := MainCollection{
				collection: arr,
			}
			for k, v := range arr.([]uint16) {
				num[k] = NewNumber(v)
			}
			return NumberCollection[uint16]{
				MainCollection: main,
				value:          num,
			}
		case reflect.Uint32:
			var num = make([]Number[uint32], rv.Len())
			main := MainCollection{
				collection: arr,
			}
			for k, v := range arr.([]uint32) {
				num[k] = NewNumber(v)
			}
			return NumberCollection[uint32]{
				MainCollection: main,
				value:          num,
			}
		case reflect.Uint64:
			var num = make([]Number[uint64], rv.Len())
			main := MainCollection{
				collection: arr,
			}
			for k, v := range arr.([]uint64) {
				num[k] = NewNumber(v)
			}
			return NumberCollection[uint64]{
				MainCollection: main,
				value:          num,
			}
		case reflect.Float32:
			var num = make([]Number[float32], rv.Len())
			main := MainCollection{
				collection: arr,
			}
			for k, v := range arr.([]float32) {
				num[k] = NewNumber(v)
			}
			return NumberCollection[float32]{
				MainCollection: main,
				value:          num,
			}
		case reflect.Float64:
			var num = make([]Number[float64], rv.Len())
			main := MainCollection{
				collection: arr,
			}
			for k, v := range arr.([]float64) {
				num[k] = NewNumber(v)
			}
			return NumberCollection[float64]{
				MainCollection: main,
				value:          num,
			}
		default:
			return nil
		}
	case reflect.String:
		return nil
	default:
		return nil
	}
}
