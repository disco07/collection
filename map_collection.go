package collection

// MapCollection is a struct that holds a map with values of type T.
type MapCollection[T any] struct {
	values map[interface{}]T
}
