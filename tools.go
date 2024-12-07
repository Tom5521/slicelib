package slicelib

import (
	"fmt"
	"reflect"
)

// deepEqual creates a deep comparison function for a specific value
// Uses reflect.DeepEqual to compare complex structures and types
// T is a generic type that can be of any type
// Returns a function that compares a value with the original value.
func deepEqual[T any](t1 T) func(T) bool {
	return func(t T) bool {
		return reflect.DeepEqual(t1, t)
	}
}

// comparableEqual creates a simple comparison function for a specific value
// Uses the == comparison operator for comparable types
// T is a generic type that can be of any type
// Returns a function that compares a value with the original value.
func comparableEqual[T any](t1 T) func(T) bool {
	return func(t T) bool {
		return any(t1) == any(t)
	}
}

// deepEqual2 compares two values using a deep comparison
// Uses reflect.DeepEqual to compare complex structures and types
// T is a generic type that can be of any type
// Returns true if the values are equal, false otherwise.
func deepEqual2[T any](t1, t2 T) bool {
	return reflect.DeepEqual(t1, t2)
}

// comparableEqual2 compares two values using the == operator
// T is a generic type that can be of any type
// Returns true if the values are equal, false otherwise.
func comparableEqual2[T any](t1, t2 T) bool {
	return any(t1) == any(t2)
}

// makeString converts an iterator into a string representation
// yield is a function that iterates over the elements
// size indicates the total number of elements
// Returns a string with formatted elements between brackets.
func makeString[T any](yield func(func(k int, v T) bool), size int) (txt string) {
	txt = "["
	yield(func(k int, v T) bool {
		txt += fmt.Sprintf(" %v", v)
		if k != size-1 {
			txt += ","
		}
		return true
	})
	txt += " ]"
	return
}

// outOfRangePanic generates a panic when an index is out of slice bounds
// i is the index that caused the error
// l is the length of the slice.
func outOfRangePanic(i, l int) {
	panic(fmt.Sprintf("runtime error: index out of range [%v] with length %v", i, l))
}

// equalSlicersFunc compares two Slicer using a custom comparison function
// s1 and s2 are the slices to compare
// f is the custom comparison function
// Returns true if the slices are considered equal according to the comparison function.
func equalSlicersFunc[T any](s1, s2 Slicer[T], f func(T, T) bool) (eq bool) {
	// Check that the slices have the same length
	if s1.Len() != s2.Len() {
		return false
	}

	// Interface for types that can return a slice pointer
	type pointerer interface{ SliceP() *[]T }

	// Check if both slices can provide a pointer
	_, s1CanPointer := s1.(pointerer)
	_, s2CanPointer := s2.(pointerer)

	// Optimization for pointer comparison
	if s1CanPointer && s2CanPointer {
		ptr1 := s1.(pointerer).SliceP()
		ptr2 := s2.(pointerer).SliceP()

		// If pointers are identical, slices are equal
		if ptr1 == ptr2 {
			return true
		}

		// Check the capacity of the slices
		if cap(*ptr1) != cap(*ptr2) {
			return false
		}
	}

	// Compare elements using the provided function
	s1.Range(func(i int, t T) bool {
		eq = f(t, s2.At(i))
		return eq
	})

	return
}
