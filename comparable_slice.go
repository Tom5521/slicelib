package slicelib

import (
	"slices"
)

type ComparableSlice[T comparable] struct {
	Slice[T]
}

func (s ComparableSlice[T]) Index(v T) int {
	return slices.Index(s.slice, v)
}

func (s ComparableSlice[T]) Contains(v T) bool {
	return slices.Contains(s.slice, v)
}

func (s ComparableSlice[T]) Equal(v []T) bool {
	return slices.Equal(s.slice, v)
}
