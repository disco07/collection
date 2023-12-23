package collections

type MainCollection[T any] struct {
	collection any
}

func (m MainCollection[T]) Value() any {
	//TODO implement me
	panic("implement me")
}

func (m MainCollection[T]) All() []any {
	//TODO implement me
	panic("implement me")
}

func (m MainCollection[T]) IsEmpty() bool {
	//TODO implement me
	panic("implement me")
}

func (m MainCollection[T]) IsNotEmpty() bool {
	//TODO implement me
	panic("implement me")
}

func (m MainCollection[T]) Len() int {
	//TODO implement me
	panic("implement me")
}

func (m MainCollection[T]) Clear() Collection[T] {
	//TODO implement me
	panic("implement me")
}

func (m MainCollection[T]) Dedup() Collection[T] {
	//TODO implement me
	panic("implement me")
}

func (m MainCollection[T]) ExtractIf(fn func()) Collection[T] {
	//TODO implement me
	panic("implement me")
}

func (m MainCollection[T]) Flatten() Collection[T] {
	//TODO implement me
	panic("implement me")
}

func (m MainCollection[T]) First() any {
	//TODO implement me
	panic("implement me")
}

func (m MainCollection[T]) SplitFirst() (any, Collection[T]) {
	//TODO implement me
	panic("implement me")
}

func (m MainCollection[T]) SplitLast() (any, Collection[T]) {
	//TODO implement me
	panic("implement me")
}

func (m MainCollection[T]) Last() any {
	//TODO implement me
	panic("implement me")
}

func (m MainCollection[T]) Get(i int) any {
	//TODO implement me
	panic("implement me")
}

func (m MainCollection[T]) Swap(i int, i2 int) Collection[T] {
	//TODO implement me
	panic("implement me")
}

func (m MainCollection[T]) Reverse() Collection[T] {
	//TODO implement me
	panic("implement me")
}

func (m MainCollection[T]) Iter() Iterator {
	//TODO implement me
	panic("implement me")
}

func (m MainCollection[T]) Windows() Iterator {
	//TODO implement me
	panic("implement me")
}

func (m MainCollection[T]) GroupBy(f func()) Iterator {
	//TODO implement me
	panic("implement me")
}

func (m MainCollection[T]) SplitAt(i int) (Collection[T], Collection[T]) {
	//TODO implement me
	panic("implement me")
}

func (m MainCollection[T]) Split(f func()) Split {
	//TODO implement me
	panic("implement me")
}

func (m MainCollection[T]) SplitInclusive(f func()) Split {
	//TODO implement me
	panic("implement me")
}

func (m MainCollection[T]) Avg(key ...string) any {
	//TODO implement me
	panic("implement me")
}

func (m MainCollection[T]) Sum(key ...string) any {
	//TODO implement me
	panic("implement me")
}

func (m MainCollection[T]) Min(key ...string) any {
	//TODO implement me
	panic("implement me")
}

func (m MainCollection[T]) Max(key ...string) any {
	//TODO implement me
	panic("implement me")
}
