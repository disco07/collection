package collections

type NumberCollection[T Num] struct {
	value []Number[T]
	MainCollection
}

func (n NumberCollection[T]) Value() any {
	//TODO implement me
	panic("implement me")
}

func (n NumberCollection[T]) All() []any {
	//TODO implement me
	panic("implement me")
}

func (n NumberCollection[T]) IsEmpty() bool {
	return len(n.value) == 0
}

func (n NumberCollection[T]) IsNotEmpty() bool {
	return len(n.value) > 0
}

func (n NumberCollection[T]) Len() int {
	return len(n.value)
}

func (n NumberCollection[T]) Clear() Collection {
	//TODO implement me
	panic("implement me")
}

func (n NumberCollection[T]) Dedup() Collection {
	//TODO implement me
	panic("implement me")
}

func (n NumberCollection[T]) ExtractIf(fn func()) Collection {
	//TODO implement me
	panic("implement me")
}

func (n NumberCollection[T]) Flatten() Collection {
	//TODO implement me
	panic("implement me")
}

func (n NumberCollection[T]) First() any {
	//TODO implement me
	panic("implement me")
}

func (n NumberCollection[T]) SplitFirst() (any, Collection) {
	//TODO implement me
	panic("implement me")
}

func (n NumberCollection[T]) SplitLast() (any, Collection) {
	//TODO implement me
	panic("implement me")
}

func (n NumberCollection[T]) Last() any {
	//TODO implement me
	panic("implement me")
}

func (n NumberCollection[T]) Get(i int) any {
	if i >= len(n.value) || len(n.value) == 0 {
		return nil
	}
	return n.value[i].value
}

func (n NumberCollection[T]) Swap(i int, i2 int) Collection {
	//TODO implement me
	panic("implement me")
}

func (n NumberCollection[T]) Reverse() Collection {
	//TODO implement me
	panic("implement me")
}

func (n NumberCollection[T]) Iter() Iterator {
	//TODO implement me
	panic("implement me")
}

func (n NumberCollection[T]) Windows() Iterator {
	//TODO implement me
	panic("implement me")
}

func (n NumberCollection[T]) GroupBy(f func()) Iterator {
	//TODO implement me
	panic("implement me")
}

func (n NumberCollection[T]) SplitAt(i int) (Collection, Collection) {
	//TODO implement me
	panic("implement me")
}

func (n NumberCollection[T]) Split(f func()) Split {
	//TODO implement me
	panic("implement me")
}

func (n NumberCollection[T]) SplitInclusive(f func()) Split {
	//TODO implement me
	panic("implement me")
}

func (n NumberCollection[T]) Avg(key ...string) any {
	var sum Number[T]

	for _, num := range n.value {
		sum.value = sum.Add(num.value)
	}

	return sum.Div(T(NewNumber(len(n.value)).value))
}

func (n NumberCollection[T]) Sum(key ...string) any {
	var sum Number[T]

	for _, num := range n.value {
		sum.value = sum.Add(num.value)
	}

	return sum.value
}

func (n NumberCollection[T]) Min(key ...string) any {
	var smallest Number[T]

	for i, num := range n.value {
		if i == 0 {
			smallest = NewNumber(num.value)
			continue
		}
		if smallest.GreaterThan(num.value) {
			smallest = NewNumber(num.value)
		}
	}

	return smallest.value
}

func (n NumberCollection[T]) Max(key ...string) any {
	var greatest Number[T]

	for i, num := range n.value {
		if i == 0 {
			greatest = NewNumber(num.value)
			continue
		}
		if greatest.LessThan(num.value) {
			greatest = NewNumber(num.value)
		}
	}

	return greatest.value
}
