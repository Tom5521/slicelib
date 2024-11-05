package slicelib

import (
	"fmt"
	"reflect"
	"slices"
)

func deepEqual[T any](t1 T) func(T) bool {
	return func(t T) bool {
		return reflect.DeepEqual(t1, t)
	}
}

func deepEqual2[T any](t1, t2 T) bool {
	return reflect.DeepEqual(t1, t2)
}

type Slice[T any] struct {
	slice []T
}

// Creates a new object of type Slice.
func NewSlice[T any](slice ...T) Slice[T] {
	return Slice[T]{slice}
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

// Return the built-in slice.
func (s Slice[T]) S() []T {
	return s.slice
}

// Return the pointer of the golang built-in slice.
func (s *Slice[T]) SliceP() *[]T {
	return &s.slice
}

// Equivalent to slice = append(slice,myElems...).
func (s *Slice[T]) Append(items ...T) {
	s.slice = append(s.slice, items...)
}

// Removes all slice elements leaving a slice with length 0 and 0 elements, but not nil.
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
	return slices.IndexFunc(s.slice, deepEqual(v))
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
	if reflect.TypeFor[T]().Comparable() {
		for _, j := range s.slice {
			if any(j) == any(v) {
				return true
			}
		}
	}
	return slices.ContainsFunc(s.slice, deepEqual(v))
}

// Removes duplicates from the slice making all elements unique.
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

// Returns a slices.EqualFunc compared with reflect.DeepEqual.
func (s Slice[T]) Equal(v []T) bool {
	if reflect.TypeFor[T]().Comparable() {
		return slices.EqualFunc(s.slice, v, func(e1 T, e2 T) bool {
			return any(e1) == any(e2)
		})
	}

	return slices.EqualFunc(s.slice, v, deepEqual2[T])
}

func (s Slice[T]) EqualSlice(v Slice[T]) bool {
	return s.Equal(v.S())
}

// A shortcut to slices.EqualFunc.
func (s Slice[T]) EqualFunc(v []T, f func(e1, e2 T) bool) bool {
	return slices.EqualFunc(s.slice, v, f)
}

func (s Slice[T]) EqualSliceFunc(v Slice[T], f func(e1, e2 T) bool) bool {
	return slices.EqualFunc(s.slice, v.S(), f)
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

// An implementation of iter.Seq2 for go 1.23 or later.
func (s *Slice[T]) Range(yield func(k int, v T) bool) {
	for i, j := range s.slice {
		yield(i, j)
	}
}

func (s Slice[T]) String() (txt string) {
	txt = "["
	s.Range(func(k int, v T) bool {
		txt += fmt.Sprintf(" %v", v)
		if k != s.Len()-1 {
			txt += ","
		}

		return true
	})

	txt += " ]"

	return
}
