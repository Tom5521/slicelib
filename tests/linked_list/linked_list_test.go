package slicelib

import (
	"testing"

	"github.com/Tom5521/slicelib"
)

func TestCreate(t *testing.T) {
	a := slicelib.NewLinkedList(3, 2, 1)
	b := []int{3, 2, 1}

	if !a.Equal(b) {
		t.Fail()
	}
}

func TestLen(t *testing.T) {
	a := slicelib.NewLinkedList(1, 2, 3)

	if a.Len() != 3 {
		t.Fail()
	}
}

func TestAt(t *testing.T) {
	a := slicelib.NewLinkedList(1, 2, 3)

	if a.At(2) != 3 {
		t.Fail()
	}
}

func TestPop(t *testing.T) {
	a := slicelib.NewLinkedList(1, 2, 3)
	b := []int{1, 2}

	a.Pop(2)

	if !a.Equal(b) {
		t.Fail()
	}
}

func TestContains(t *testing.T) {
	a := slicelib.NewLinkedList(1, 2, 3)

	if !a.Contains(3) {
		t.Fail()
	}
}

func TestIndex(t *testing.T) {
	a := slicelib.NewLinkedList(1, 2, 3)

	if a.Index(3) != 2 {
		t.Fail()
	}
}

func TestClear(t *testing.T) {
	a := slicelib.NewLinkedList(1, 2, 3)
	a.Clear()

	if a.Len() != 0 {
		t.Fail()
	}
}

func TestDelete(t *testing.T) {
	a := slicelib.NewLinkedList(1, 2, 3, 4, 5, 6, 7, 8)
	b := []int{1, 6, 7, 8}

	a.Delete(1, 5)

	if !a.Equal(b) {
		t.Fail()
	}
}

func TestAppend(t *testing.T) {
	a := slicelib.NewLinkedList(1, 2, 3)
	b := []int{1, 2, 3, 4, 5, 6}

	a.Append(4, 5, 6)

	if !a.Equal(b) {
		t.Fail()
	}

	c := slicelib.NewLinkedList[int]()
	c.Append(1)
	c.Append(2)
	c.Append(3)
}

func TestReverse(t *testing.T) {
	a := slicelib.NewLinkedList(1, 2, 3)
	b := []int{3, 2, 1}

	a.Reverse()

	if !a.Equal(b) {
		t.Log(a)
		t.Fail()
	}
}

func TestSet(t *testing.T) {
	a := slicelib.NewLinkedList(1, 2, 3)
	b := []int{1, 2, 2}

	a.Set(2, 2)

	if !a.Equal(b) {
		t.Fail()
	}
}

func TestSliceLeft(t *testing.T) {
	a := slicelib.NewLinkedList(1, 2, 3, 4)
	b := []int{3, 4}

	a.SliceLeft(2)

	if !a.Equal(b) {
		t.Fail()
	}

	c := slicelib.NewLinkedList(1, 2, 3, 4)
	d := []int{}

	c.SliceLeft(4)
	if !c.Equal(d) {
		t.Fail()
	}
}

func TestSliceRight(t *testing.T) {
	a := slicelib.NewLinkedList(1, 2, 3, 4)
	b := []int{1, 2}

	a.SliceRight(2)

	if !a.Equal(b) {
		t.Log(a)
		t.Fail()
	}
}

func TestSliceRange(t *testing.T) {
	a := slicelib.NewLinkedList(1, 2, 3, 4)
	b := []int{2, 3}

	a.SliceRange(1, 3)

	if !a.Equal(b) {
		t.Fail()
	}
}

func TestRemoveDuplicates(t *testing.T) {
	a := slicelib.NewLinkedList(1, 1, 2, 2, 3, 3, 4, 4)
	b := []int{1, 2, 3, 4}

	a.RemoveDuplicates()

	if !a.Equal(b) {
		t.Fail()
	}
}

func TestInsert(t *testing.T) {
	a := slicelib.NewLinkedList(1, 2, 3)
	b := []int{1, 2, 3, 4, 5}

	a.Insert(
		2, // Index
		// Values...
		4,
		5,
	)

	if !a.Equal(b) {
		t.Fail()
	}
}

func TestFilter(t *testing.T) {
	a := slicelib.NewLinkedList[any](1, "Meow", 3.1, "22", '1')
	b := []any{"Meow", "22"}

	a.Filter(func(t any) (pass bool) {
		_, ok := t.(string)
		return ok
	})

	if !a.Equal(b) {
		t.Fail()
	}
}
