package slicelib

import (
	"reflect"
	"slices"
)

type node[T any] struct {
	data           T
	previous, next *node[T]
}

type LinkedList[T any] struct {
	head *node[T]
	len  int
}

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

func (ll *LinkedList[T]) last() *node[T] {
	c := ll.head
	for c != nil {
		c = c.next
	}
	return c
}

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

func (ll *LinkedList[T]) refreshLen() {
	var l int
	ll.yield(func(i int, _ *node[T]) bool {
		l++
		return true
	})

	ll.len = l
}

func (ll *LinkedList[T]) isOutOfRange(i int) bool {
	return i >= ll.len || i < 0
}

// TODO: Finish this.
func (ll *LinkedList[T]) createNodes(s ...T) *node[T] {
	return nil
}

func NewLinkedList[T any](slice ...T) *LinkedList[T] {
	ll := new(LinkedList[T])
	ll.Append(slice...)

	return ll
}

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

func (ll *LinkedList[T]) S() (slice []T) {
	ll.Range(func(_ int, t T) bool {
		slice = append(slice, t)
		return true
	})
	return
}

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

func (ll *LinkedList[T]) Len() int {
	return ll.len
}

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

func (ll *LinkedList[T]) InRange(i int) bool {
	return !ll.isOutOfRange(i)
}

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

func (ll *LinkedList[T]) Remove(val T) {
	ll.Pop(ll.Index(val))
}

func (ll *LinkedList[T]) String() string {
	return makeString(ll.Range, ll.Len())
}

func (ll *LinkedList[T]) Clear() {
	ll.head = nil
	ll.len = 0
}

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

func (ll *LinkedList[T]) Equal(s []T) bool {
	if reflect.TypeFor[T]().Comparable() {
		return ll.EqualFunc(s, comparableEqual2[T])
	}

	return ll.EqualFunc(s, deepEqual2[T])
}

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

func (ll *LinkedList[T]) EqualSlicerFunc(s Slicer[T], f func(T, T) bool) (eq bool) {
	if ll.Len() != s.Len() {
		return false
	}
	if p, ok := s.(interface{ SliceP() *[]T }); ok {
		s := p.SliceP()
		if len(*s) != cap(*s) {
			return false
		}
	}

	ll.Range(func(i int, t T) bool {
		eq = f(t, s.At(i))
		return eq
	})

	return
}

func (ll *LinkedList[T]) EqualSlicer(s Slicer[T]) bool {
	if reflect.TypeFor[T]().Comparable() {
		return ll.EqualSlicerFunc(s, comparableEqual2[T])
	}
	return ll.EqualSlicerFunc(s, deepEqual2[T])
}

// TODO: Finish this.
func (ll *LinkedList[T]) Insert(i int, values ...T) {
	if ll.isOutOfRange(i) {
		outOfRangePanic(i, ll.len)
	}
	n := ll.at(i)
	if n == nil {
		return
	}
}

func (ll *LinkedList[T]) IsEmpty() bool {
	return ll.head == nil
}
