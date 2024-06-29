package slicelib

import (
	"slices"
)

type ComparableSlice[T comparable] struct {
	Slice[T]
}

func NewComparableSlice[T comparable](slice ...T) ComparableSlice[T] {
	return ComparableSlice[T]{NewSlice(slice...)}
}

func (s ComparableSlice[T]) Copy() ComparableSlice[T] {
	return NewComparableSlice(slices.Clone(s.slice)...)
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
