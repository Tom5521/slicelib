package slicelib

import (
	"reflect"
	"slices"
)

type Slice[T any] []T

func (s Slice[T]) Slice() []T {
	return []T(s)
}

func (s *Slice[T]) Append(items ...T) {
	*s = append(*s, items...)
}

func (s *Slice[T]) Clear() {
	*s = Slice[T]{}
}

func (s Slice[T]) Copy() Slice[T] {
	return slices.Clone(s)
}

func (s Slice[T]) Index(v T) int {
	return slices.IndexFunc(s, func(e T) bool {
		return reflect.DeepEqual(e, v)
	})
}

func (s *Slice[T]) Insert(index int, items ...T) {
	*s = slices.Insert(*s, index, items...)
}

func (s *Slice[T]) Delete(i, j int) {
	*s = slices.Delete(*s, i, j)
}

func (s *Slice[T]) Pop(index int) {
	*s = slices.Delete(*s, index, index+1)
}

func (s *Slice[T]) Remove(v T) {
	s.Pop(s.Index(v))
}

func (s *Slice[T]) Reverse() {
	slices.Reverse(*s)
}

func (s Slice[T]) IsEmpty() bool {
	return len(s) == 0
}

func (s Slice[T]) Len() int {
	return len(s)
}

func (s Slice[T]) Contains(v T) bool {
	return slices.ContainsFunc(s, func(e T) bool {
		return reflect.DeepEqual(e, v)
	})
}

func (s *Slice[T]) RemoveDuplicates() {
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

func (s Slice[T]) Equal(v Slice[T]) bool {
	return slices.EqualFunc(s, v, func(e1 T, e2 T) bool {
		return reflect.DeepEqual(e1, e2)
	})
}

func (s Slice[T]) EqualFunc(v Slice[T], f func(e1, e2 T) bool) bool {
	return slices.EqualFunc(s, v, f)
}

func (s *Slice[T]) SortFunc(f func(a, b T) int) {
	slices.SortFunc(*s, f)
}

func (s *Slice[T]) Filter(f func(T) bool) {
	var newSlice Slice[T]
	for _, i := range *s {
		if f(i) {
			newSlice.Append(i)
		}
	}

	*s = newSlice
}
