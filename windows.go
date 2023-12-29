package collection

type windows[T any] struct {
	slice  []T
	index  int
	window int
}

func (w *windows[T]) Next() ([]T, bool) {
	if w.index+w.window > len(w.slice) {
		return nil, false
	}
	result := w.slice[w.index : w.index+w.window]
	w.index++
	return result, true
}
