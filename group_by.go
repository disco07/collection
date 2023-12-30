package collection

type GroupBy[T any] struct {
	slice      []T
	index      int
	comparator func(T, T) bool
}

func (g *GroupBy[T]) Next() ([]T, bool) {
	if g.index >= len(g.slice) {
		return nil, false
	}

	start := g.index
	for g.index < len(g.slice)-1 && g.comparator(g.slice[g.index], g.slice[g.index+1]) {
		g.index++
	}

	g.index++
	return g.slice[start:g.index], true
}
