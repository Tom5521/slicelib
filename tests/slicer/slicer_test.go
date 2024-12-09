package slicer_test

import (
	"testing"

	"github.com/Tom5521/slicelib"
)

type TestMode int

func Test(t *testing.T) {
	makers := []func(...int) slicelib.Slicer[int]{
		func(i ...int) slicelib.Slicer[int] {
			return slicelib.NewLinkedList(i...)
		},
		func(i ...int) slicelib.Slicer[int] {
			return slicelib.NewSlice(i...)
		},
		func(i ...int) slicelib.Slicer[int] {
			return slicelib.NewOrderedSlice(i...)
		},
		func(i ...int) slicelib.Slicer[int] {
			return slicelib.NewComparableSlice(i...)
		},
	}

	type test struct {
		input          []int
		input2, input3 any
		expected       any
		pass           func(s slicelib.Slicer[int], tt test) (pass bool)
	}

	tests := []test{
		// Create
		{
			input:    []int{3, 2, 1},
			expected: []int{3, 2, 1},

			pass: func(s slicelib.Slicer[int], tt test) bool {
				return s.Equal(tt.expected.([]int))
			},
		},
		// Len
		{
			input:    []int{1, 2, 3},
			expected: 3,

			pass: func(s slicelib.Slicer[int], tt test) (pass bool) {
				return s.Len() == tt.expected
			},
		},
		// At
		{
			input:    []int{1, 2, 3},
			input2:   2,
			expected: 3,

			pass: func(s slicelib.Slicer[int], tt test) (pass bool) {
				return s.At(tt.input2.(int)) == tt.expected
			},
		},
		// Pop
		{
			input:    []int{1, 2, 3},
			expected: []int{1, 2},

			pass: func(s slicelib.Slicer[int], tt test) (pass bool) {
				s.Pop(2)
				return s.Equal(tt.expected.([]int))
			},
		},
		// Contains
		{
			input:    []int{1, 2, 3},
			input2:   2,
			expected: true,

			pass: func(s slicelib.Slicer[int], tt test) (pass bool) {
				return s.Contains(tt.input2.(int))
			},
		},
		// Index
		{
			input:    []int{1, 2, 3},
			input2:   3,
			expected: 2,

			pass: func(s slicelib.Slicer[int], tt test) (pass bool) {
				return s.Index(tt.input2.(int)) == tt.expected
			},
		},
		// Clear
		{
			input:    []int{1, 2, 3},
			expected: []int{},

			pass: func(s slicelib.Slicer[int], tt test) (pass bool) {
				s.Clear()
				return s.Equal(tt.expected.([]int))
			},
		},
		// Delete
		{
			input:    []int{1, 2, 3, 4, 5, 6, 7, 8},
			expected: []int{1, 6, 7, 8},
			input2:   1,
			input3:   5,

			pass: func(s slicelib.Slicer[int], tt test) (pass bool) {
				s.Delete(tt.input2.(int), tt.input3.(int))
				return s.Equal(tt.expected.([]int))
			},
		},
		// Append
		{
			input:    []int{1, 2, 3},
			expected: []int{1, 2, 3, 4},
			input2:   4,

			pass: func(s slicelib.Slicer[int], tt test) (pass bool) {
				s.Append(tt.input2.(int))
				return s.Equal(tt.expected.([]int))
			},
		},
		// Reverse
		{
			input:    []int{1, 2, 3},
			expected: []int{3, 2, 1},

			pass: func(s slicelib.Slicer[int], tt test) (pass bool) {
				s.Reverse()
				return s.Equal(tt.expected.([]int))
			},
		},
		// Set
		{
			input:    []int{1, 2, 3},
			input2:   2,
			input3:   2,
			expected: []int{1, 2, 2},

			pass: func(s slicelib.Slicer[int], tt test) (pass bool) {
				s.Set(tt.input2.(int), tt.input3.(int))
				return s.Equal(tt.expected.([]int))
			},
		},
		// SliceLeft
		{
			input:    []int{1, 2, 3, 4},
			input2:   2,
			expected: []int{3, 4},

			pass: func(s slicelib.Slicer[int], tt test) (pass bool) {
				s.SliceLeft(2)
				return s.Equal(tt.expected.([]int))
			},
		},
		// SliceRight
		{
			input:    []int{1, 2, 3, 4},
			input2:   2,
			expected: []int{1, 2},

			pass: func(s slicelib.Slicer[int], tt test) (pass bool) {
				s.SliceRight(tt.input2.(int))
				return s.Equal(tt.expected.([]int))
			},
		},
		// SliceRange
		{
			input:    []int{1, 2, 3, 4},
			input2:   1,
			input3:   3,
			expected: []int{2, 3},

			pass: func(s slicelib.Slicer[int], tt test) (pass bool) {
				s.SliceRange(tt.input2.(int), tt.input3.(int))
				return s.Equal(tt.expected.([]int))
			},
		},
		// RemoveDuplicates
		{
			input:    []int{1, 1, 2, 2, 3, 3, 4, 4},
			expected: []int{1, 2, 3, 4},

			pass: func(s slicelib.Slicer[int], tt test) (pass bool) {
				s.RemoveDuplicates()
				return s.Equal(tt.expected.([]int))
			},
		},
		// Insert
		{
			input:    []int{1, 2, 3},
			input2:   2,
			input3:   []int{4, 5},
			expected: []int{1, 2, 3, 4, 5},

			pass: func(s slicelib.Slicer[int], tt test) (pass bool) {
				s.Insert(tt.input2.(int), tt.input3.([]int)...)
				return s.Equal(tt.expected.([]int))
			},
		},
		// Filter
		{
			input: []int{10, 20, 30, 40, 50, 60, 70, 80, 200, 100, 500},
			input2: func(v int) bool {
				return v > 10 && v < 100
			},
			expected: []int{20, 30, 40, 50, 60, 70, 80},

			pass: func(s slicelib.Slicer[int], tt test) (pass bool) {
				s.Filter(tt.input2.(func(int) bool))
				return s.Equal(tt.expected.([]int))
			},
		},
		// Equal (comparable)
		{
			input:    []int{1, 2, 3},
			input2:   []int{1, 2, 3},
			expected: true,

			pass: func(s slicelib.Slicer[int], tt test) (pass bool) {
				return s.Equal(tt.input2.([]int))
			},
		},
		// Clone
		{
			input:    []int{1, 2, 3},
			input2:   []int{1, 2, 3},
			expected: true,

			pass: func(s slicelib.Slicer[int], tt test) (pass bool) {
				return s.Equal(tt.input2.([]int))
			},
		},
		// String
		{
			input:    []int{1, 2, 3},
			expected: "[ 1, 2, 3 ]",

			pass: func(s slicelib.Slicer[int], tt test) (pass bool) {
				return s.String() == tt.expected
			},
		},
	}

	for i, maker := range makers {
		t.Logf("Testing maker nº %d", i+1)
		for i, test := range tests {
			maked := maker(test.input...)
			t.Logf("Testing nº %d...", i+1)
			if test.pass == nil {
				t.Logf("Test nº %d is nil!", i+1)
				t.Log("Skipping...")
				continue
			}
			if !test.pass(maked, test) {
				t.Logf("Test nº %d fail!", i+1)
				t.Log("Input:", test.input, "Expected:", test.expected, "Received:", maked)
				t.Fail()
			} else {
				t.Log("PASS")
			}
		}
	}
}
