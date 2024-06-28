package slicelib

import (
	"cmp"
	"slices"
)

type OrderedSlice[T cmp.Ordered] []T

func (s OrderedSlice[T]) Slice() []T {
	return []T(s)
}

func (s *OrderedSlice[T]) Append(items ...T) {
	*s = append(*s, items...)
}

func (s *OrderedSlice[T]) Clear() {
	*s = OrderedSlice[T]{}
}

func (s OrderedSlice[T]) Copy() OrderedSlice[T] {
	return slices.Clone(s)
}

func (s *OrderedSlice[T]) Insert(index int, items ...T) {
	*s = slices.Insert(*s, index, items...)
}

func (s *OrderedSlice[T]) Delete(i, j int) {
	*s = slices.Delete(*s, i, j)
}

func (s *OrderedSlice[T]) Pop(index int) {
	*s = slices.Delete(*s, index, index+1)
}

func (s *OrderedSlice[T]) Remove(v T) {
	s.Pop(s.Index(v))
}

func (s *OrderedSlice[T]) Reverse() {
	slices.Reverse(*s)
}

func (s OrderedSlice[T]) IsEmpty() bool {
	return len(s) == 0
}

func (s OrderedSlice[T]) Len() int {
	return len(s)
}

func (s *OrderedSlice[T]) RemoveDuplicates() {
	seen := make(map[any]bool)
	j := 0
	for i := 0; i < len(*s); i++ {
		if !seen[(*s)[i]] {
			seen[(*s)[i]] = true
			(*s)[j] = (*s)[i]
			j++
		}
	}
	*s = (*s)[:j]
}

func (s OrderedSlice[T]) EqualFunc(v Slice[T], f func(e1, e2 T) bool) bool {
	return slices.EqualFunc(s, v, f)
}

func (s *OrderedSlice[T]) SortFunc(f func(a, b T) int) {
	slices.SortFunc(*s, f)
}

func (s *OrderedSlice[T]) Filter(f func(T) bool) {
	var newSlice OrderedSlice[T]
	for _, i := range *s {
		if f(i) {
			newSlice.Append(i)
		}
	}

	*s = newSlice
}

// CUSTOM METHODS

func (s OrderedSlice[T]) Index(v T) int {
	return slices.Index(s, v)
}

func (s OrderedSlice[T]) Contains(v T) bool {
	return slices.Contains(s, v)
}

func (s OrderedSlice[T]) Equal(v OrderedSlice[T]) bool {
	return slices.Equal(s, v)
}

func (s *OrderedSlice[T]) Sort() {
	slices.Sort(*s)
}

func (s OrderedSlice[T]) BinarySearch(v T) (int, bool) {
	return slices.BinarySearch(s, v)
}
