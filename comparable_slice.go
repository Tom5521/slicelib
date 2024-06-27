package sliceutils

import (
	"slices"
)

type ComparableSlice[T comparable] struct {
	Slice[T]
}

func (s ComparableSlice[T]) Index(v T) int {
	return slices.Index(s.Slice, v)
}

func (s ComparableSlice[T]) Contains(v T) bool {
	return slices.Contains(s.Slice, v)
}

func (s ComparableSlice[T]) Equal(v ComparableSlice[T]) bool {
	return slices.Equal(s.Slice, v.Slice)
}
