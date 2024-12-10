package slicer_test

import (
	"testing"

	"github.com/Tom5521/slicelib"
)

func Test(t *testing.T) {
	makers := []func([]int) slicelib.Slicer[int]{
		func(i []int) slicelib.Slicer[int] {
			return slicelib.NewLinkedList(i...)
		},
		func(i []int) slicelib.Slicer[int] {
			return slicelib.NewSlice(i...)
		},
		func(i []int) slicelib.Slicer[int] {
			return slicelib.NewOrderedSlice(i...)
		},
		func(i []int) slicelib.Slicer[int] {
			return slicelib.NewComparableSlice(i...)
		},
	}

	type test struct {
		name           string
		input          []int
		input2, input3 any
		expected       any
		pass           func(s slicelib.Slicer[int], tt test) (pass bool)
	}

	tests := []test{
		{
			name:     "Create",
			input:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},

			pass: func(s slicelib.Slicer[int], tt test) bool {
				return s.Equal(tt.expected.([]int))
			},
		},
		{
			name:     "Len",
			input:    []int{1, 2, 3},
			expected: 3,

			pass: func(s slicelib.Slicer[int], tt test) (pass bool) {
				return s.Len() == tt.expected
			},
		},
		{
			name:     "At",
			input:    []int{1, 2, 3},
			input2:   2,
			expected: 3,

			pass: func(s slicelib.Slicer[int], tt test) (pass bool) {
				return s.At(tt.input2.(int)) == tt.expected
			},
		},
		{
			name:     "Pop",
			input:    []int{1, 2, 3},
			expected: []int{1, 2},

			pass: func(s slicelib.Slicer[int], tt test) (pass bool) {
				s.Pop(2)
				return s.Equal(tt.expected.([]int))
			},
		},
		{
			name:     "Contains",
			input:    []int{1, 2, 3},
			input2:   2,
			expected: true,

			pass: func(s slicelib.Slicer[int], tt test) (pass bool) {
				return s.Contains(tt.input2.(int))
			},
		},
		{
			name:     "Index",
			input:    []int{1, 2, 3},
			input2:   3,
			expected: 2,

			pass: func(s slicelib.Slicer[int], tt test) (pass bool) {
				return s.Index(tt.input2.(int)) == tt.expected
			},
		},
		{
			name:     "Clear",
			input:    []int{1, 2, 3},
			expected: []int{},

			pass: func(s slicelib.Slicer[int], tt test) (pass bool) {
				s.Clear()
				return s.Equal(tt.expected.([]int))
			},
		},
		{
			name:     "Delete",
			input:    []int{1, 2, 3, 4, 5, 6, 7, 8},
			expected: []int{1, 6, 7, 8},
			input2:   1,
			input3:   5,

			pass: func(s slicelib.Slicer[int], tt test) (pass bool) {
				s.Delete(tt.input2.(int), tt.input3.(int))
				return s.Equal(tt.expected.([]int))
			},
		},
		{
			name:     "Append",
			input:    []int{1, 2, 3},
			expected: []int{1, 2, 3, 4},
			input2:   4,

			pass: func(s slicelib.Slicer[int], tt test) (pass bool) {
				s.Append(tt.input2.(int))
				return s.Equal(tt.expected.([]int))
			},
		},
		{
			name:     "Reverse",
			input:    []int{1, 2, 3},
			expected: []int{3, 2, 1},

			pass: func(s slicelib.Slicer[int], tt test) (pass bool) {
				s.Reverse()
				return s.Equal(tt.expected.([]int))
			},
		},
		{
			name:     "Set",
			input:    []int{1, 2, 3},
			input2:   2,
			input3:   2,
			expected: []int{1, 2, 2},

			pass: func(s slicelib.Slicer[int], tt test) (pass bool) {
				s.Set(tt.input2.(int), tt.input3.(int))
				return s.Equal(tt.expected.([]int))
			},
		},
		{
			name:     "SliceLeft",
			input:    []int{1, 2, 3, 4},
			input2:   2,
			expected: []int{3, 4},

			pass: func(s slicelib.Slicer[int], tt test) (pass bool) {
				s.SliceLeft(2)
				return s.Equal(tt.expected.([]int))
			},
		},
		{
			name:     "SliceRight",
			input:    []int{1, 2, 3, 4},
			input2:   2,
			expected: []int{1, 2},

			pass: func(s slicelib.Slicer[int], tt test) (pass bool) {
				s.SliceRight(tt.input2.(int))
				return s.Equal(tt.expected.([]int))
			},
		},
		{
			name:     "SliceRange",
			input:    []int{1, 2, 3, 4},
			input2:   1,
			input3:   3,
			expected: []int{2, 3},

			pass: func(s slicelib.Slicer[int], tt test) (pass bool) {
				s.SliceRange(tt.input2.(int), tt.input3.(int))
				return s.Equal(tt.expected.([]int))
			},
		},
		{
			name:     "RemoveDuplicates",
			input:    []int{1, 1, 2, 2, 3, 3, 4, 4},
			expected: []int{1, 2, 3, 4},

			pass: func(s slicelib.Slicer[int], tt test) (pass bool) {
				s.RemoveDuplicates()
				return s.Equal(tt.expected.([]int))
			},
		},
		{
			name:     "Insert",
			input:    []int{1, 2, 3},
			input2:   3,
			input3:   []int{4, 5},
			expected: []int{1, 2, 3, 4, 5},

			pass: func(s slicelib.Slicer[int], tt test) (pass bool) {
				s.Insert(tt.input2.(int), tt.input3.([]int)...)
				return s.Equal(tt.expected.([]int))
			},
		},
		{
			name:  "Filter",
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
		{
			name:     "Equal",
			input:    []int{1, 2, 3},
			input2:   []int{1, 2, 3},
			expected: true,

			pass: func(s slicelib.Slicer[int], tt test) (pass bool) {
				return s.Equal(tt.input2.([]int))
			},
		},
		{
			name:     "Clone",
			input:    []int{1, 2, 3},
			input2:   []int{1, 2, 3},
			expected: true,

			pass: func(s slicelib.Slicer[int], tt test) (pass bool) {
				return s.Equal(tt.input2.([]int))
			},
		},
		{
			name:     "String",
			input:    []int{1, 2, 3},
			expected: "[ 1, 2, 3 ]",

			pass: func(s slicelib.Slicer[int], tt test) (pass bool) {
				return s.String() == tt.expected
			},
		},
	}

	for i, maker := range makers {
		t.Log("---^---")
		t.Logf("Testing maker nº %d", i+1)
		t.Log("---v---")
		for i, test := range tests {
			maked := maker(test.input)
			t.Logf("Test nº %d [%s]...", i+1, test.name)
			if test.pass == nil {
				t.Logf("Test nº %d [%s] is nil!", i+1, test.name)
				t.Log("Skipping...")
				continue
			}
			if !test.pass(maked, test) {
				t.Logf("%+v", test)
				t.Log("Result:", maked)
				t.Log("--- FAIL ---")
				t.Fail()
			} else {
				t.Log("PASS")
			}
		}
	}
}
