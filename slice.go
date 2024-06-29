package slicelib

import (
	"reflect"
	"slices"
)

type Slice[T any] struct {
	slice []T
}

func NewSlice[T any](slice ...T) Slice[T] {
	return Slice[T]{
		slice: slice,
	}
}

func (s Slice[T]) Slice() []T {
	return s.slice
}

func (s *Slice[T]) SliceP() *[]T {
	return &s.slice
}

func (s *Slice[T]) Append(items ...T) {
	s.slice = append(s.slice, items...)
}

func (s *Slice[T]) Clear() {
	s.slice = []T{}
}

func (s Slice[T]) Copy() Slice[T] {
	return Slice[T]{
		slice: slices.Clone(s.slice),
	}
}

func (s Slice[T]) Index(v T) int {
	return slices.IndexFunc(s.slice, func(e T) bool {
		return reflect.DeepEqual(e, v)
	})
}

func (s *Slice[T]) Insert(index int, items ...T) {
	s.slice = slices.Insert(s.slice, index, items...)
}

func (s *Slice[T]) Delete(i, j int) {
	s.slice = slices.Delete(s.slice, i, j)
}

func (s *Slice[T]) Pop(index int) {
	s.slice = slices.Delete(s.slice, index, index+1)
}

func (s *Slice[T]) Remove(v T) {
	s.Pop(s.Index(v))
}

func (s *Slice[T]) Reverse() {
	slices.Reverse(s.slice)
}

func (s Slice[T]) IsEmpty() bool {
	return len(s.slice) == 0
}

func (s Slice[T]) Len() int {
	return len(s.slice)
}

func (s Slice[T]) Contains(v T) bool {
	return slices.ContainsFunc(s.slice, func(e T) bool {
		return reflect.DeepEqual(e, v)
	})
}

func (s *Slice[T]) RemoveDuplicates() {
	seen := make(map[any]bool)
	j := 0
	for i := 0; i < len(s.slice); i++ {
		if !seen[(s.slice)[i]] {
			seen[(s.slice)[i]] = true
			(s.slice)[j] = (s.slice)[i]
			j++
		}
	}
	s.slice = s.slice[:j]
}

func (s Slice[T]) Equal(v []T) bool {
	return slices.EqualFunc(s.slice, v, func(e1 T, e2 T) bool {
		return reflect.DeepEqual(e1, e2)
	})
}

func (s Slice[T]) EqualFunc(v []T, f func(e1, e2 T) bool) bool {
	return slices.EqualFunc(s.slice, v, f)
}

func (s *Slice[T]) SortFunc(f func(a, b T) int) {
	slices.SortFunc(s.slice, f)
}

func (s *Slice[T]) Filter(f func(T) bool) {
	var newSlice []T
	for _, i := range s.slice {
		if f(i) {
			newSlice = append(newSlice, i)
		}
	}

	s.slice = newSlice
}
