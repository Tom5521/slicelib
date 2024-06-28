package slicelib

import (
	"slices"
)

type ComparableSlice[T comparable] []T

func (s ComparableSlice[T]) Slice() []T {
	return []T(s)
}

func (s *ComparableSlice[T]) Append(items ...T) {
	*s = append(*s, items...)
}

func (s *ComparableSlice[T]) Clear() {
	*s = ComparableSlice[T]{}
}

func (s ComparableSlice[T]) Copy() ComparableSlice[T] {
	return slices.Clone(s)
}

func (s *ComparableSlice[T]) Insert(index int, items ...T) {
	*s = slices.Insert(*s, index, items...)
}

func (s *ComparableSlice[T]) Delete(i, j int) {
	*s = slices.Delete(*s, i, j)
}

func (s *ComparableSlice[T]) Pop(index int) {
	*s = slices.Delete(*s, index, index+1)
}

func (s *ComparableSlice[T]) Remove(v T) {
	s.Pop(s.Index(v))
}

func (s *ComparableSlice[T]) Reverse() {
	slices.Reverse(*s)
}

func (s ComparableSlice[T]) IsEmpty() bool {
	return len(s) == 0
}

func (s ComparableSlice[T]) Len() int {
	return len(s)
}

func (s *ComparableSlice[T]) RemoveDuplicates() {
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

func (s ComparableSlice[T]) EqualFunc(v Slice[T], f func(e1, e2 T) bool) bool {
	return slices.EqualFunc(s, v, f)
}

func (s *ComparableSlice[T]) Sort(f func(a, b T) int) {
	slices.SortFunc(*s, f)
}

func (s *ComparableSlice[T]) Filter(f func(T) bool) {
	var newSlice ComparableSlice[T]
	for _, i := range *s {
		if f(i) {
			newSlice.Append(i)
		}
	}

	*s = newSlice
}

// CUSTOM METHODS

func (s ComparableSlice[T]) Index(v T) int {
	return slices.Index(s, v)
}

func (s ComparableSlice[T]) Contains(v T) bool {
	return slices.Contains(s, v)
}

func (s ComparableSlice[T]) Equal(v ComparableSlice[T]) bool {
	return slices.Equal(s, v)
}
