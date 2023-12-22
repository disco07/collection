package collections

type MainCollection struct {
	collection any
}

func (m MainCollection) Value() any {
	//TODO implement me
	panic("implement me")
}

func (m MainCollection) All() []any {
	//TODO implement me
	panic("implement me")
}

func (m MainCollection) IsEmpty() bool {
	//TODO implement me
	panic("implement me")
}

func (m MainCollection) IsNotEmpty() bool {
	//TODO implement me
	panic("implement me")
}

func (m MainCollection) Len() int {
	//TODO implement me
	panic("implement me")
}

func (m MainCollection) Clear() Collection {
	//TODO implement me
	panic("implement me")
}

func (m MainCollection) Dedup() Collection {
	//TODO implement me
	panic("implement me")
}

func (m MainCollection) ExtractIf(fn func()) Collection {
	//TODO implement me
	panic("implement me")
}

func (m MainCollection) Flatten() Collection {
	//TODO implement me
	panic("implement me")
}

func (m MainCollection) First() any {
	//TODO implement me
	panic("implement me")
}

func (m MainCollection) SplitFirst() (any, Collection) {
	//TODO implement me
	panic("implement me")
}

func (m MainCollection) SplitLast() (any, Collection) {
	//TODO implement me
	panic("implement me")
}

func (m MainCollection) Last() any {
	//TODO implement me
	panic("implement me")
}

func (m MainCollection) Get(i int) any {
	//TODO implement me
	panic("implement me")
}

func (m MainCollection) Swap(i int, i2 int) Collection {
	//TODO implement me
	panic("implement me")
}

func (m MainCollection) Reverse() Collection {
	//TODO implement me
	panic("implement me")
}

func (m MainCollection) Iter() Iterator {
	//TODO implement me
	panic("implement me")
}

func (m MainCollection) Windows() Iterator {
	//TODO implement me
	panic("implement me")
}

func (m MainCollection) GroupBy(f func()) Iterator {
	//TODO implement me
	panic("implement me")
}

func (m MainCollection) SplitAt(i int) (Collection, Collection) {
	//TODO implement me
	panic("implement me")
}

func (m MainCollection) Split(f func()) Split {
	//TODO implement me
	panic("implement me")
}

func (m MainCollection) SplitInclusive(f func()) Split {
	//TODO implement me
	panic("implement me")
}

func (m MainCollection) Avg(key ...string) any {
	//TODO implement me
	panic("implement me")
}

func (m MainCollection) Sum(key ...string) any {
	//TODO implement me
	panic("implement me")
}

func (m MainCollection) Min(key ...string) any {
	//TODO implement me
	panic("implement me")
}

func (m MainCollection) Max(key ...string) any {
	//TODO implement me
	panic("implement me")
}
