package slicelib

import (
	"fmt"
	"reflect"
)

func deepEqual[T any](t1 T) func(T) bool {
	return func(t T) bool {
		return reflect.DeepEqual(t1, t)
	}
}

func comparableEqual[T any](t1 T) func(T) bool {
	return func(t T) bool {
		return any(t1) == any(t)
	}
}

func deepEqual2[T any](t1, t2 T) bool {
	return reflect.DeepEqual(t1, t2)
}

func comparableEqual2[T any](t1, t2 T) bool {
	return any(t1) == any(t2)
}

func makeString[T any](yield func(func(k int, v T) bool), size int) (txt string) {
	txt = "["

	yield(func(k int, v T) bool {
		txt += fmt.Sprintf(" %v", v)
		if k != size-1 {
			txt += ","
		}

		return true
	})

	txt += " ]"

	return
}

func outOfRangePanic(i, l int) {
	panic(fmt.Sprintf("runtime error: index out of range [%v] with length %v", i, l))
}
