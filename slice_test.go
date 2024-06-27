package sliceutils_test

import (
	"testing"

	sliceutils "github.com/Tom5521/slice-utils"
)

func TestEqual(t *testing.T) {
	a := sliceutils.Slice[int]{1, 2, 3}
	b := []int{1, 2, 3}

	if !a.Equal(b) {
		t.Fail()
	}
}

func TestReverse(t *testing.T) {
	a := sliceutils.Slice[int]{1, 2, 3}
	b := []int{3, 2, 1}

	a.Reverse()

	if !a.Equal(b) {
		t.Fail()
	}
}

func TestContains(t *testing.T) {
	a := sliceutils.Slice[int]{1, 2, 3}
	if !a.Contains(1) {
		t.Fail()
	}
}

func TestRemoveDuplicates(t *testing.T) {
	a := sliceutils.Slice[int]{1, 1, 2, 3}
	b := []int{1, 2, 3}

	a.RemoveDuplicates()

	if !a.Equal(b) {
		t.Fail()
	}
}

func TestIndex(t *testing.T) {
	a := sliceutils.Slice[int]{1, 2, 3}

	if a.Index(1) != 0 {
		t.Fail()
	}
}

func TestPop(t *testing.T) {
	a := sliceutils.Slice[int]{1, 2, 3}
	b := []int{2, 3}

	a.Pop(0)

	if !a.Equal(b) {
		t.Fail()
	}
}

func TestRemove(t *testing.T) {
	a := sliceutils.Slice[int]{1, 2, 3}
	b := []int{1, 3}

	a.Remove(2)

	if !a.Equal(b) {
		t.Fail()
	}
}

func TestDelete(t *testing.T) {
	a := sliceutils.Slice[int]{1, 2, 3}
	b := []int{1, 3}

	a.Delete(1, 2)

	if !a.Equal(b) {
		t.Fail()
	}
}

func TestInsert(t *testing.T) {
	a := sliceutils.Slice[int]{1, 2, 3}
	b := []int{1, 2, 3, 3}

	a.Insert(len(a)-1, 3)

	if !a.Equal(b) {
		t.Fail()
	}
}
