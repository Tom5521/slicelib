//go:build !test
// +build !test

package slicelib

var (
	_ Slicer[any] = (*Slice[any])(nil)
	_ Slicer[any] = (*LinkedList[any])(nil)
	_ Slicer[int] = (*ComparableSlice[int])(nil)
	_ Slicer[int] = (*OrderedSlice[int])(nil)
)
