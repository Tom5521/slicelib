package slicelib

import (
	"testing"

	"github.com/Tom5521/slicelib"
)

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
	a := slicelib.NewLinkedList(1, 2, 3)
	b := []int{1, 3}

	a.Delete(1, 2)

	if !a.Equal(b) {
		t.Fail()
	}
}
