package slicelib

import (
	"slices"
)

type ComparableSlice[T comparable] struct {
	Slice[T]
}

// Creates a new object of ComparableSlice wich only implements
// the comparable interface.
func NewComparableSlice[T comparable](slice ...T) ComparableSlice[T] {
	return ComparableSlice[T]{NewSlice(slice...)}
}

// Creates a copy of the current object, which is not the same as the current object.
//
// implements the slices.Clone function on the internal slice to create the new structure.
func (s ComparableSlice[T]) Copy() ComparableSlice[T] {
	return NewComparableSlice(slices.Clone(s.slice)...)
}

// A shortcut to slices.Index.
func (s ComparableSlice[T]) Index(v T) int {
	return slices.Index(s.slice, v)
}

// A shortcut to slices.Contains.
func (s ComparableSlice[T]) Contains(v T) bool {
	return slices.Contains(s.slice, v)
}

// A shortcut to slices.Equal.
func (s ComparableSlice[T]) Equal(v []T) bool {
	return slices.Equal(s.slice, v)
}
