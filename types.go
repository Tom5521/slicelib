package slicelib

type Slicer[T any] interface {
	At(int) T
	S() []T
	Append(...T)
	Len() int
	Remove(T)
	RemoveLast(T)
	Pop(int)
	Delete(int, int)
	String() string
	Range(func(int, T) bool)
	ReverseRange(func(int, T) bool)
	Index(T) int
	LastIndex(T) int
	Contains(T) bool
	Clear()
	Insert(int, ...T)
	Reverse()
	IsEmpty() bool
	RemoveDuplicates()
	Equal([]T) bool
	EqualSlicer(Slicer[T]) bool
	EqualFunc([]T, func(T, T) bool) bool
	EqualSlicerFunc(Slicer[T], func(T, T) bool) bool
	SortFunc(func(T, T) int)
	SliceRight(int)      // [:x]
	SliceLeft(int)       // [x:]
	SliceRange(int, int) // [x:y]
	Set(int, T)
	InRange(int) bool
	Filter(func(T) bool)
}
