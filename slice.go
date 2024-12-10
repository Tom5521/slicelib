package slicelib

import (
	"reflect"
	"slices"
)

// Slice is a generic wrapper around a built-in slice that provides additional utility methods.
// It allows for more flexible slice operations with generic type support.
type Slice[T any] struct {
	slice []T
}

// NewSlice creates a new Slice instance with the provided elements.
// It allows creating a generic slice with any number of initial elements.
//
// Example:
//
//	intSlice := NewSlice(1, 2, 3)
//	stringSlice := NewSlice("a", "b", "c")
func NewSlice[T any](slice ...T) *Slice[T] {
	return &Slice[T]{slices.Clone(slice)}
}

// Elem returns the element at the specified index.
//
// Deprecated: Use At() method instead for consistency.
// This method will be removed in future versions.
func (s Slice[T]) Elem(index int) T {
	return s.At(index)
}

// At returns the element at the specified index.
// Panics if the index is out of bounds.
func (s Slice[T]) At(index int) T {
	return s.slice[index]
}

// S returns the underlying built-in slice.
// Useful for interoperability with standard Go slice functions.
func (s Slice[T]) S() []T {
	return s.slice
}

// SliceP returns a pointer to the underlying slice.
// Allows direct manipulation of the slice's underlying storage.
func (s *Slice[T]) SliceP() *[]T {
	return &s.slice
}

// Append adds one or more elements to the end of the slice.
// Equivalent to using the built-in append function.
func (s *Slice[T]) Append(items ...T) {
	s.slice = append(s.slice, items...)
}

// Cap returns the capacity of the underlying slice.
// Mirrors the cap() built-in function behavior.
func (s *Slice[T]) Cap() int {
	return cap(s.slice)
}

func (s *Slice[T]) Set(i int, v T) {
	s.slice[i] = v
}

// Clear removes all elements from the slice,
// leaving it with zero length but not nil.
func (s *Slice[T]) Clear() {
	s.slice = []T{}
}

// Clone creates a deep copy of the current Slice.
// Uses slices.Clone to create a new slice with the same elements.
//
// Returns a new Slice instance with copied elements.
func (s Slice[T]) Clone() *Slice[T] {
	return NewSlice(slices.Clone(s.slice)...)
}

// CloneS returns a clone of the underlying slice.
// Directly uses slices.Clone without creating a new Slice wrapper.
func (s Slice[T]) CloneS() []T {
	return slices.Clone(s.slice)
}

// Index finds the first occurrence of the specified value in the slice.
// Uses different comparison strategies based on the type's comparability:
// - For comparable types, uses standard equality
// - For non-comparable types, uses deep reflection comparison
//
// Returns the index of the first matching element, or -1 if not found.
func (s Slice[T]) Index(v T) int {
	if reflect.TypeFor[T]().Comparable() {
		return slices.IndexFunc(s.slice, comparableEqual(v))
	}
	return slices.IndexFunc(s.slice, deepEqual(v))
}

func (s *Slice[T]) LastIndex(val T) int {
	i := -1
	if reflect.TypeFor[T]().Comparable() {
		s.ReverseRange(func(ii int, t T) bool {
			found := any(t) == any(val)
			if found {
				i = ii
			}
			return !found
		})
	} else {
		s.Range(func(k int, v T) bool {
			found := reflect.DeepEqual(val, v)
			if found {
				i = k
			}
			return !found
		})
	}
	return i
}

// Insert adds one or more elements at the specified index.
// Shifts existing elements to make room for the new elements.
//
// Panics if the index is out of bounds.
func (s *Slice[T]) Insert(index int, items ...T) {
	s.slice = slices.Insert(s.slice, index, items...)
}

// Delete removes elements from the slice between indices i and j.
// Follows the behavior of slices.Delete.
func (s *Slice[T]) Delete(i, j int) {
	s.slice = slices.Delete(s.slice, i, j)
}

// Pop removes and returns the element at the specified index.
// Reduces the slice length by one.
func (s *Slice[T]) Pop(index int) {
	s.slice = slices.Delete(s.slice, index, index+1)
}

// Remove finds and removes the first occurrence of the specified value.
// Uses the Index method to locate the element before removing it.
func (s *Slice[T]) Remove(v T) {
	s.Pop(s.Index(v))
}

func (s *Slice[T]) RemoveLast(v T) {
	s.Pop(s.LastIndex(v))
}

// Reverse changes the order of elements in the slice to their reverse.
func (s *Slice[T]) Reverse() {
	slices.Reverse(s.slice)
}

// IsEmpty checks if the slice contains no elements.
// Returns true if the slice length is zero.
func (s Slice[T]) IsEmpty() bool {
	return len(s.slice) == 0
}

// Len returns the number of elements in the slice.
func (s Slice[T]) Len() int {
	return len(s.slice)
}

// Contains checks if the slice includes the specified value.
// Uses different comparison strategies based on the type's comparability.
func (s Slice[T]) Contains(v T) bool {
	if reflect.TypeFor[T]().Comparable() {
		return slices.ContainsFunc(s.slice, comparableEqual(v))
	}
	return slices.ContainsFunc(s.slice, deepEqual(v))
}

// RemoveDuplicates eliminates duplicate elements, keeping only unique values.
// Preserves the order of first occurrences.
func (s *Slice[T]) RemoveDuplicates() {
	seen := make(map[any]bool)
	var j int
	for i, v := range s.slice {
		if !seen[v] {
			seen[v] = true
			s.slice[j] = s.slice[i]
			j++
		}
	}

	s.slice = s.slice[:j]
}

// Equal compares the slice with another slice for equality.
// Uses different comparison strategies based on the type's comparability.
func (s Slice[T]) Equal(v []T) bool {
	if reflect.TypeFor[T]().Comparable() {
		return slices.EqualFunc(s.slice, v, comparableEqual2[T])
	}

	return slices.EqualFunc(s.slice, v, deepEqual2[T])
}

// EqualSlice compares the current slice with another Slice instance.
func (s Slice[T]) EqualSlice(v Slice[T]) bool {
	return s.Equal(v.S())
}

// EqualFunc allows custom equality comparison using a provided function.
func (s Slice[T]) EqualFunc(v []T, f func(e1, e2 T) bool) bool {
	return slices.EqualFunc(s.slice, v, f)
}

// EqualSliceFunc compares two Slice instances using a custom comparison function.
func (s Slice[T]) EqualSliceFunc(v Slice[T], f func(e1, e2 T) bool) bool {
	return slices.EqualFunc(s.slice, v.S(), f)
}

// SortFunc allows custom sorting using a comparison function.
func (s *Slice[T]) SortFunc(f func(a, b T) int) {
	slices.SortFunc(s.slice, f)
}

// Filter removes elements that do not match the provided predicate function.
// Keeps only elements for which the function returns true.
func (s *Slice[T]) Filter(f func(T) (pass bool)) {
	var newSlice []T
	for _, i := range s.slice {
		if f(i) {
			newSlice = append(newSlice, i)
		}
	}

	s.slice = newSlice
}

// Range provides an iterator-like functionality for the slice.
// Compatible with Go 1.23+ iter.Seq2 interface.
// The yield function receives the index and value, and can stop iteration by returning false.
func (s *Slice[T]) Range(yield func(k int, v T) bool) {
	for i, j := range s.slice {
		if !yield(i, j) {
			break
		}
	}
}

func (s *Slice[T]) ReverseRange(f func(int, T) bool) {
	for i := len(s.slice) - 1; i >= 0; i-- {
		if !f(i, s.slice[i]) {
			break
		}
	}
}

// String returns a string representation of the slice.
// Provides a readable format for printing or logging.
func (s Slice[T]) String() string {
	return makeString(s.Range, s.Len())
}

// Grow increases the slice's capacity to accommodate more elements.
// Useful for performance optimization when adding multiple elements.
func (s *Slice[T]) Grow(newSize int) {
	s.slice = slices.Grow(s.slice, newSize)
}

// Clip reduces the slice's capacity to its length.
// Helps in memory optimization by removing excess capacity.
func (s *Slice[T]) Clip() {
	s.slice = slices.Clip(s.slice)
}

// EqualSlicerFunc compares the current slice with another Slicer using a custom comparison function.
func (s *Slice[T]) EqualSlicerFunc(v Slicer[T], f func(T, T) bool) bool {
	return equalSlicersFunc(s, v, f)
}

// EqualSlicer compares the current slice with another Slicer.
// Uses appropriate comparison strategy based on the type's comparability.
func (s *Slice[T]) EqualSlicer(v Slicer[T]) bool {
	if reflect.TypeFor[T]().Comparable() {
		return s.EqualSlicerFunc(v, comparableEqual2[T])
	}
	return s.EqualSlicerFunc(v, deepEqual2[T])
}

// Is equal to slice[:x]
func (s *Slice[T]) SliceRight(i int) {
	s.slice = s.slice[:i]
}

// Is equal to slice[x:]
func (s *Slice[T]) SliceLeft(i int) {
	s.slice = s.slice[i:]
}

// Is equal to slice[x:y]
func (s *Slice[T]) SliceRange(i, j int) {
	s.slice = s.slice[i:j]
}

func (s *Slice[T]) InRange(i int) bool {
	return i >= 0 && i < len(s.slice)
}
