package slicelib

import (
	"cmp"
	"slices"
)

type OrderedSlice[T cmp.Ordered] struct {
	Slice[T]
}

func NewOrderedSlice[T cmp.Ordered](slice ...T) OrderedSlice[T] {
	return OrderedSlice[T]{NewSlice(slice...)}
}

func (s *OrderedSlice[T]) Sort() {
	slices.Sort(s.slice)
}

func (s OrderedSlice[T]) BinarySearch(v T) (int, bool) {
	return slices.BinarySearch(s.slice, v)
}
