package slicelib_test

import (
	"testing"

	"github.com/Tom5521/slicelib"
)

func TestOrderedSort(t *testing.T) {
	a := slicelib.OrderedSlice[int]{3, 1, 2}
	b := []int{1, 2, 3}

	a.Sort()

	if !a.Equal(b) {
		t.Fail()
	}
}

func TestOrderedBinarySearch(t *testing.T) {
	a := slicelib.OrderedSlice[int]{1, 2, 3, 4, 5, 6, 7, 8, 9}

	_, found := a.BinarySearch(5)

	if !found {
		t.Fail()
	}
}
