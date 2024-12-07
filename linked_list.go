package slicelib

import (
	"reflect"
	"slices"
)

// Package slicelib provides a generic implementation of a doubly-linked list.

// node represents an individual element in the LinkedList.
// It stores the data and maintains links to previous and next nodes.
type node[T any] struct {
	data           T
	previous, next *node[T]
}

// LinkedList is a generic doubly-linked list implementation.
// It provides dynamic data storage with efficient insertion and deletion operations.
type LinkedList[T any] struct {
	head *node[T] // First node of the list
	len  int      // Total number of elements in the list
}

// yield is an internal iterator method that traverses the list.
// Allows performing operations on each node with the ability to stop iteration.
//
// f is a function that receives the current index and node.
// Returns false from f to stop iteration.
func (ll *LinkedList[T]) yield(f func(i int, n *node[T]) bool) {
	c := ll.head
	var index int
	for c != nil {
		if !f(index, c) {
			break
		}

		c = c.next
		index++
	}
}

// last finds and returns the last node in the list.
// Returns nil if the list is empty.
func (ll *LinkedList[T]) last() *node[T] {
	c := ll.head
	for c != nil {
		c = c.next
	}
	return c
}

// at retrieves the node at a specific index.
// Returns nil if the index is out of range.
func (ll *LinkedList[T]) at(i int) (n *node[T]) {
	ll.yield(func(ii int, nn *node[T]) bool {
		if ii == i {
			n = nn
			return false
		}
		return true
	})

	return
}

// refreshLen recalculates the length of the list.
// Useful after operations that might modify the list structure.
func (ll *LinkedList[T]) refreshLen() {
	var l int
	ll.yield(func(i int, _ *node[T]) bool {
		l++
		return true
	})

	ll.len = l
}

// isOutOfRange checks if a given index is outside the list's bounds.
func (ll *LinkedList[T]) isOutOfRange(i int) bool {
	return i >= ll.len || i < 0
}

// NewLinkedList creates a new LinkedList with optional initial elements.
//
// Example:
//
//	list := NewLinkedList(1, 2, 3)
//	stringList := NewLinkedList("a", "b", "c")
func NewLinkedList[T any](slice ...T) *LinkedList[T] {
	ll := new(LinkedList[T])
	ll.Append(slice...)

	return ll
}

// Range iterates through the list, allowing operations on each element.
// Provides a functional iteration approach similar to slice range.
//
// The function receives (index, value) and can stop iteration by returning false.
func (ll *LinkedList[T]) Range(f func(int, T) bool) {
	c := ll.head
	var index int
	for c != nil {
		if !f(index, c.data) {
			break
		}
		c = c.next
		index++
	}
}

// At retrieves the element at the specified index.
// Panics if the index is out of range.
func (ll *LinkedList[T]) At(i int) (v T) {
	if ll.isOutOfRange(i) {
		outOfRangePanic(i, ll.len)
	}
	ll.Range(func(ii int, t T) bool {
		if ii == i {
			v = t
			return false
		}
		return true
	})

	return
}

// S converts the LinkedList to a standard Go slice.
// Useful for interoperability with slice-based functions.
func (ll *LinkedList[T]) S() (slice []T) {
	ll.Range(func(_ int, t T) bool {
		slice = append(slice, t)
		return true
	})
	return
}

// Append adds one or more elements to the end of the list.
func (ll *LinkedList[T]) Append(s ...T) {
	current := ll.last()
	if current == nil && len(s) > 0 {
		ll.head = &node[T]{
			data: s[0],
		}
		ll.len++
		current = ll.head
		s = s[1:]
	}

	for _, value := range s {
		newNode := &node[T]{
			previous: current,
			data:     value,
		}
		current.next = newNode
		current = newNode
		ll.len++
	}
}

// Len returns the number of elements in the list.
func (ll *LinkedList[T]) Len() int {
	return ll.len
}

// Pop removes the element at the specified index.
// Panics if the index is out of range.
func (ll *LinkedList[T]) Pop(i int) {
	if ll.isOutOfRange(i) {
		outOfRangePanic(i, ll.len)
	}

	n := ll.at(i)
	if n == nil {
		return
	}

	if n.previous != nil {
		n.previous.next = n.next
	} else {
		ll.head = n.next
	}
	if n.next != nil {
		n.next.previous = n.previous
	}

	ll.len--
}

// InRange checks if the given index is within the list's bounds.
func (ll *LinkedList[T]) InRange(i int) bool {
	return !ll.isOutOfRange(i)
}

// Contains checks if the list includes a specific value.
// Uses appropriate comparison strategy based on type comparability.
func (ll *LinkedList[T]) Contains(v T) (contains bool) {
	if reflect.TypeFor[T]().Comparable() {
		ll.Range(func(_ int, t T) bool {
			contains = any(t) == any(v)
			return !contains
		})

		return
	}
	ll.Range(func(_ int, t T) bool {
		contains = reflect.DeepEqual(t, v)
		return !contains
	})

	return
}

// Index finds the first occurrence of a value in the list.
// Returns -1 if the value is not found.
func (ll *LinkedList[T]) Index(val T) (index int) {
	index = -1
	if reflect.TypeFor[T]().Comparable() {
		ll.Range(func(i int, t T) bool {
			if any(t) == any(val) {
				index = i
				return false
			}
			return true
		})
		return
	}

	ll.Range(func(i int, t T) bool {
		if reflect.DeepEqual(val, t) {
			index = i
			return false
		}
		return true
	})
	return
}

// Remove finds and removes the first occurrence of a value.
func (ll *LinkedList[T]) Remove(val T) {
	ll.Pop(ll.Index(val))
}

// String provides a string representation of the list.
func (ll *LinkedList[T]) String() string {
	return makeString(ll.Range, ll.Len())
}

// Clear removes all elements from the list.
func (ll *LinkedList[T]) Clear() {
	ll.head = nil
	ll.len = 0
}

// Delete removes elements between indices i and j.
// Panics if either index is out of range.
func (ll *LinkedList[T]) Delete(i, j int) {
	if ll.isOutOfRange(i) {
		outOfRangePanic(i, ll.len)
	}
	if ll.isOutOfRange(j) {
		outOfRangePanic(i, ll.len)
	}

	s := slices.Delete(ll.S(), i, j)
	ll.Clear()
	ll.Append(s...)
	ll.refreshLen()
}

// Equal compares the list with a slice for equality.
// Uses deep or standard comparison based on type comparability.
func (ll *LinkedList[T]) Equal(s []T) bool {
	if reflect.TypeFor[T]().Comparable() {
		return ll.EqualFunc(s, comparableEqual2[T])
	}

	return ll.EqualFunc(s, deepEqual2[T])
}

// EqualFunc allows custom comparison of the list with a slice.
func (ll *LinkedList[T]) EqualFunc(s []T, f func(T, T) bool) (eq bool) {
	if len(s) != ll.Len() || cap(s) != len(s) {
		return false
	}

	ll.Range(func(i int, t T) bool {
		eq = f(s[i], t)
		return eq
	})

	return
}

// EqualSlicerFunc compares the list with another Slicer using a custom function.
func (ll *LinkedList[T]) EqualSlicerFunc(s Slicer[T], f func(T, T) bool) bool {
	return equalSlicersFunc[T](ll, s, f)
}

// EqualSlicer compares the list with another Slicer.
// Uses appropriate comparison strategy based on type comparability.
func (ll *LinkedList[T]) EqualSlicer(s Slicer[T]) bool {
	if reflect.TypeFor[T]().Comparable() {
		return ll.EqualSlicerFunc(s, comparableEqual2[T])
	}
	return ll.EqualSlicerFunc(s, deepEqual2[T])
}

// IsEmpty checks if the list contains no elements.
func (ll *LinkedList[T]) IsEmpty() bool {
	return ll.head == nil
}

// Insert is a placeholder for future implementation of inserting elements at a specific index.
// Currently not fully implemented.
func (ll *LinkedList[T]) Insert(i int, values ...T) {
	if ll.isOutOfRange(i) {
		outOfRangePanic(i, ll.len)
	}
	n := ll.at(i)
	if n == nil {
		return
	}
	// TODO: Implement insertion logic
}

// TODO: Implement methods for:
// - RemoveDuplicates
// - Reverse
// - SortFunc

func (ll *LinkedList[T]) RemoveDuplicates()           {}
func (ll *LinkedList[T]) Reverse()                    {}
func (ll *LinkedList[T]) SortFunc(f func(a, b T) int) {}
