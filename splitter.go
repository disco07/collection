package collection

type splitter[T any] struct {
	slices [][]T
	index  int
}

func (s *splitter[T]) Next() ([]T, bool) {
	if s.index >= len(s.slices) {
		return nil, false
	}
	result := s.slices[s.index]
	s.index++
	return result, true
}
