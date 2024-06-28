package slicelib_test

import (
	"testing"

	sliceutils "github.com/Tom5521/slicelib"
)

func TestComparableEqual(t *testing.T) {
	a := sliceutils.OrderedSlice[int]{1, 2, 3}
	b := []int{1, 2, 3}

	if !a.Equal(b) {
		t.Fail()
	}
}

func TestComparableIndex(t *testing.T) {
	a := sliceutils.ComparableSlice[int]{1, 2, 3}

	if a.Index(1) != 0 {
		t.Fail()
	}
}

func TestComparableContains(t *testing.T) {
	a := sliceutils.ComparableSlice[int]{1, 2, 3}
	if !a.Contains(1) {
		t.Fail()
	}
}
