package slicelib_test

import (
	"testing"

	sliceutils "github.com/Tom5521/slicelib"
)

func TestEqual(t *testing.T) {
	a := sliceutils.NewSlice(1, 2, 3)
	b := []int{1, 2, 3}
	// Test Comparable.
	if !a.Equal(b) {
		t.Log("Fail on comparable")
		t.Fail()
	}

	// Test Not comparable.
	type str struct {
		A1 int
	}

	c := sliceutils.NewSlice(str{1}, str{2}, str{3})
	d := []str{{1}, {2}, {3}}

	if !c.Equal(d) {
		t.Log("Fail on non comparable")
		t.Fail()
	}
}

func TestReverse(t *testing.T) {
	a := sliceutils.NewSlice(1, 2, 3)
	b := []int{3, 2, 1}

	a.Reverse()

	if !a.Equal(b) {
		t.Fail()
	}
}

func TestContains(t *testing.T) {
	a := sliceutils.NewSlice(1, 2, 3)
	if !a.Contains(1) {
		t.Fail()
	}
}

func TestRemoveDuplicates(t *testing.T) {
	a := sliceutils.NewSlice(1, 1, 2, 2, 3, 3)
	b := []int{1, 2, 3}

	a.RemoveDuplicates()

	if !a.Equal(b) {
		t.Fail()
	}
}

func TestIndex(t *testing.T) {
	a := sliceutils.NewSlice(1, 2, 3)

	if a.Index(1) != 0 {
		t.Fail()
	}
}

func TestPop(t *testing.T) {
	a := sliceutils.NewSlice(1, 2, 3)
	b := []int{2, 3}

	a.Pop(0)

	if !a.Equal(b) {
		t.Fail()
	}
}

func TestRemove(t *testing.T) {
	a := sliceutils.NewSlice(1, 2, 3)
	b := []int{1, 3}

	a.Remove(2)

	if !a.Equal(b) {
		t.Fail()
	}
}

func TestDelete(t *testing.T) {
	a := sliceutils.NewSlice(1, 2, 3)
	b := []int{1, 3}

	a.Delete(1, 2)

	if !a.Equal(b) {
		t.Fail()
	}
}

func TestInsert(t *testing.T) {
	a := sliceutils.NewSlice(1, 2, 3)
	b := []int{1, 2, 3, 3}

	a.Insert(len(a.S())-1, 3)

	if !a.Equal(b) {
		t.Fail()
	}
}

func TestFilter(t *testing.T) {
	a := sliceutils.NewSlice(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	b := []int{2, 4, 6, 8, 10}

	a.Filter(func(i int) bool {
		return i%2 == 0
	})

	if !a.Equal(b) {
		t.Fail()
	}
}

func TestCopy(t *testing.T) {
	a := sliceutils.NewSlice(1, 2, 3)

	if !a.Equal(a.Clone().S()) {
		t.Fail()
	}
}

func TestString(t *testing.T) {
	a := sliceutils.NewSlice(1, 2, 3)
	expected := "[ 1, 2, 3 ]"

	if a.String() != expected {
		t.Fail()
	}
}

func TestGrow(t *testing.T) {
	a := sliceutils.NewSlice(1, 2, 3)
	a.Grow(2)

	if a.Cap() != 6 {
		t.Fail()
	}
}

func TestClip(t *testing.T) {
	a := sliceutils.NewSlice(1, 2, 3)
	a.Grow(3)

	a.Clip()
	if a.Cap() != 3 {
		t.Fail()
	}
}
