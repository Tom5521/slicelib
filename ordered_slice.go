package slicelib

import (
	"cmp"
	"slices"
)

type OrderedSlice[T cmp.Ordered] struct {
	ComparableSlice[T]
}

// Create a new OrderedSlice object.
// Which only implements the cmp.Ordered and comparable interfaces.
func NewOrderedSlice[T cmp.Ordered](slice ...T) OrderedSlice[T] {
	return OrderedSlice[T]{NewComparableSlice(slice...)}
}

// A shortcut to slices.Sort.
func (s *OrderedSlice[T]) Sort() {
	slices.Sort(s.slice)
}

// Creates a copy of the current object, which is not the same as the current object.
// implements the slices.Clone function on the internal slice to create the new structure.
func (s OrderedSlice[T]) Clone() OrderedSlice[T] {
	return NewOrderedSlice(slices.Clone(s.slice)...)
}

// A shortcut to slices.IsSorted.
func (s OrderedSlice[T]) IsSorted() bool {
	return slices.IsSorted(s.slice)
}

// A shortcut to slices.BinarySearch.
func (s OrderedSlice[T]) BinarySearch(v T) (int, bool) {
	return slices.BinarySearch(s.slice, v)
}
