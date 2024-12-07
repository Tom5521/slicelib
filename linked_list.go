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
	tail *node[T] // Last node of the list
	len  int      // Total number of elements in the list
}

func (ll *LinkedList[T]) makeNodeChain(items ...T) (h, t *node[T], l int) {
	l = len(items)
	var cur *node[T]
	if l > 0 {
		cur = &node[T]{
			data: items[0],
		}
		items = items[1:]
		h = cur
	}
	if l == 1 {
		t = h
	}

	for i, item := range items {
		newNode := &node[T]{
			previous: cur,
			data:     item,
		}
		cur.next = newNode
		cur = newNode

		if i == len(items)-1 {
			t = newNode
		}
	}

	return
}

// iter is an internal iterator method that traverses the list.
// Allows performing operations on each node with the ability to stop iteration.
//
// f is a function that receives the current index and node.
// Returns false from f to stop iteration.
func (ll *LinkedList[T]) iter(f func(i int, n *node[T]) bool) {
	for i, c := 0, ll.head; c != nil; i, c = i+1, c.next {
		if !f(i, c) {
			break
		}
	}
}

func (ll *LinkedList[T]) reverseIter(f func(i int, n *node[T]) bool) {
	for i, c := ll.len-1, ll.tail; c != nil; i, c = i-1, c.previous {
		if !f(i, c) {
			break
		}
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

func (ll *LinkedList[T]) first() *node[T] {
	c := ll.tail
	for c != nil {
		c = c.previous
	}
	return c
}

// at retrieves the node at a specific index.
// Returns nil if the index is out of range.
func (ll *LinkedList[T]) at(i int) (n *node[T]) {
	if !ll.InRange(i) {
		outOfRangePanic(i, ll.len)
	}

	yield := func(ii int, nn *node[T]) bool {
		found := ii == i
		if found {
			n = nn
		}
		return !found
	}

	if i > ll.len/2 {
		ll.reverseIter(yield)
	} else {
		ll.iter(yield)
	}

	return
}

// refreshLen recalculates the length of the list.
// Useful after operations that might modify the list structure.
func (ll *LinkedList[T]) refreshLen() {
	var l int
	ll.iter(func(_ int, _ *node[T]) bool {
		l++
		return true
	})

	ll.len = l
}

func (ll *LinkedList[T]) index(rangeFunc func(func(int, *node[T]) bool), val T) (index int) {
	index = -1
	if reflect.TypeFor[T]().Comparable() {
		rangeFunc(func(i int, n *node[T]) bool {
			if any(n.data) == any(val) {
				index = i
				return false
			}
			return true
		})
		return
	}

	rangeFunc(func(i int, n *node[T]) bool {
		if reflect.DeepEqual(val, n.data) {
			index = i
			return false
		}
		return true
	})
	return
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
	ll.iter(func(i int, n *node[T]) bool {
		return f(i, n.data)
	})
}

func (ll *LinkedList[T]) ReverseRange(f func(int, T) bool) {
	ll.reverseIter(func(i int, n *node[T]) bool {
		return f(i, n.data)
	})
}

// At retrieves the element at the specified index.
// Panics if the index is out of range.
func (ll *LinkedList[T]) At(i int) T {
	return ll.at(i).data
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
	head, tail, length := ll.makeNodeChain(s...)
	if ll.head == nil {
		ll.head = head
		ll.tail = tail
		ll.len = length
	} else if length > 0 {
		head.previous = ll.tail
		ll.tail.next = head
		ll.tail = tail
		ll.len += length
	}
}

// Len returns the number of elements in the list.
func (ll *LinkedList[T]) Len() int {
	return ll.len
}

// Pop removes the element at the specified index.
// Panics if the index is out of range.
func (ll *LinkedList[T]) Pop(i int) {
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
	} else {
		ll.tail = n.previous
	}

	ll.len--
}

// InRange checks if the given index is within the list's bounds.
func (ll *LinkedList[T]) InRange(i int) bool {
	return i >= 0 && i < ll.len
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

// Index finds the last occurrence of a value in the list.
// Returns -1 if the value is not found.
func (ll *LinkedList[T]) LastIndex(val T) int {
	return ll.index(ll.reverseIter, val)
}

// Index finds the first occurrence of a value in the list.
// Returns -1 if the value is not found.
func (ll *LinkedList[T]) Index(val T) int {
	return ll.index(ll.iter, val)
}

// Remove finds and removes the first occurrence of a value.
func (ll *LinkedList[T]) Remove(val T) {
	ll.Pop(ll.Index(val))
}

// RemoveLast finds and removes the last occurrence of a value.
func (ll *LinkedList[T]) RemoveLast(val T) {
	ll.Pop(ll.LastIndex(val))
}

// String provides a string representation of the list.
func (ll *LinkedList[T]) String() string {
	return makeString(ll.Range, ll.Len())
}

// Clear removes all elements from the list.
func (ll *LinkedList[T]) Clear() {
	ll.tail = nil
	ll.head = nil
	ll.len = 0
}

// Delete removes elements between indices i and j.
// Panics if either index is out of range.
func (ll *LinkedList[T]) Delete(i, j int) {
	if !ll.InRange(i) {
		outOfRangePanic(i, ll.len)
	}
	if !ll.InRange(j) {
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
	return equalSlicersFunc(ll, s, f)
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
	if !ll.InRange(i) {
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
// - CutUntil
// - CutFrom
// - CutRange

func (ll *LinkedList[T]) RemoveDuplicates() {
	seen := make(map[any]bool)

	var j int
	ll.Range(func(i int, t T) bool {
		if !seen[t] {
			seen[t] = true
			ll.Set(j, ll.At(i))
			j++
		}

		return true
	})

	ll.CutUntil(j)
}

func (ll *LinkedList[T]) Reverse() {
	var (
		prev *node[T]
		next *node[T]
	)

	curr := ll.head

	for curr != nil {
		next = curr.next
		curr.next = prev
		prev = curr
		curr = next
	}

	ll.head = prev
	ll.tail = next
}

func (ll *LinkedList[T]) Set(i int, v T) {
	ll.at(i).data = v
}

func (ll *LinkedList[T]) SortFunc(cmp func(a, b T) int) {
	slice := ll.S()
	slices.SortFunc(slice, cmp)
	ll.Clear()
	ll.Append(slice...)
}

func (ll *LinkedList[T]) CutUntil(i int) {
	if i <= 0 {
		ll.Clear()
		return
	}

	if i > ll.len {
		return
	}
	n := ll.at(i)
	n.previous.next = nil
	ll.tail = n.previous

	ll.len = i
}
func (ll *LinkedList[T]) CutFrom(i int)     {}
func (ll *LinkedList[T]) CutRange(i, j int) {}
