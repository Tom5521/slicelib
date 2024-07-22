package slicelib

import (
	"reflect"
	"slices"
)

type Slice[T any] struct {
	slice []T
}

// Creates a new object of type Slice.
func NewSlice[T any](slice ...T) Slice[T] {
	return Slice[T]{
		slice: slice,
	}
}

// Return the element of the given index.
//
// Deprecated: Uses Slice.At instead.
func (s Slice[T]) Elem(index int) T {
	return s.At(index)
}

// Return the element of the given index.
func (s Slice[T]) At(index int) T {
	return s.slice[index]
}

// Returns a normal slice of go.
func (s Slice[T]) S() []T {
	return s.slice
}

// Returns a pointer of a normal go slice.
func (s *Slice[T]) SliceP() *[]T {
	return &s.slice
}

// Equivalent to slice = append(slice,myElems...).
func (s *Slice[T]) Append(items ...T) {
	s.slice = append(s.slice, items...)
}

// Removes all slice elements leaving a slice with length 0 and 0 elements.
func (s *Slice[T]) Clear() {
	s.slice = []T{}
}

// Creates a copy of the current object, which is not the same as the current object.
//
// implements the slices.Clone function on the internal slice to create the new structure.
func (s Slice[T]) Copy() Slice[T] {
	return NewSlice(slices.Clone(s.slice)...)
}

// Performs a slices.IndexFunc comparing it with reflect.DeepEqual on the internal slice,
// returning the result.
//
// Returns the first occurrence of the supplied element.
func (s Slice[T]) Index(v T) int {
	return slices.IndexFunc(s.slice, func(e T) bool {
		return reflect.DeepEqual(e, v)
	})
}

// Call the slices.Insert function with the internal slice.
func (s *Slice[T]) Insert(index int, items ...T) {
	s.slice = slices.Insert(s.slice, index, items...)
}

// Call the slices.Delete function with the internal slice.
func (s *Slice[T]) Delete(i, j int) {
	s.slice = slices.Delete(s.slice, i, j)
}

// Removes only the provided index.
func (s *Slice[T]) Pop(index int) {
	s.slice = slices.Delete(s.slice, index, index+1)
}

// Removes the first occurrence of the provided element.
func (s *Slice[T]) Remove(v T) {
	s.Pop(s.Index(v))
}

// Call the slices.Reverse function with the internal slice.
func (s *Slice[T]) Reverse() {
	slices.Reverse(s.slice)
}

// Return if the slice length is 0.
func (s Slice[T]) IsEmpty() bool {
	return len(s.slice) == 0
}

// Returns the length of the slice.
func (s Slice[T]) Len() int {
	return len(s.slice)
}

// Returns true if the slice contains the element,
// comparing with slices.ContainsFunc and reflect.DeepEqual.
func (s Slice[T]) Contains(v T) bool {
	return slices.ContainsFunc(s.slice, func(e T) bool {
		return reflect.DeepEqual(e, v)
	})
}

// Removes duplicates from the slice making all elements unique.
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

// Returns a slices.EqualFunc compared with reflect.DeepEqual.
func (s Slice[T]) Equal(v []T) bool {
	return slices.EqualFunc(s.slice, v, func(e1 T, e2 T) bool {
		return reflect.DeepEqual(e1, e2)
	})
}

// A shortcut to slices.EqualFunc.
func (s Slice[T]) EqualFunc(v []T, f func(e1, e2 T) bool) bool {
	return slices.EqualFunc(s.slice, v, f)
}

// A direct access to slices.SortFunc.
func (s *Slice[T]) SortFunc(f func(a, b T) int) {
	slices.SortFunc(s.slice, f)
}

// Delete all elements that do not match the logic of the function.
func (s *Slice[T]) Filter(f func(T) bool) {
	var newSlice []T
	for _, i := range s.slice {
		if f(i) {
			newSlice = append(newSlice, i)
		}
	}

	s.slice = newSlice
}

func (s *Slice[T]) Range(yield func(k int, v T) bool) {
	for i, j := range s.slice {
		yield(i, j)
	}
}
