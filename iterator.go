package collection

type iterator[T any] struct {
	slice []T
	index int
}

func (it *iterator[T]) Next() (*T, bool) {
	if it.index >= len(it.slice) {
		return nil, false
	}
	element := &it.slice[it.index]
	it.index++
	return element, true
}
