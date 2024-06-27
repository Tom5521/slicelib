package sliceutils

import (
	"cmp"
	"slices"
)

type OrderedSlice[T cmp.Ordered] struct {
	ComparableSlice[T]
}

func (s *OrderedSlice[T]) Sort() {
	slices.Sort(s.Slice)
}

func (s OrderedSlice[T]) BinarySearch(v T) (int, bool) {
	return slices.BinarySearch(s.Slice, v)
}
