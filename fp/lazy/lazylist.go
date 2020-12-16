package main

import (
	"fmt"
)

type LazyInt = func() int
type LazyState = func() State

type State struct {
	head *LazyInt
	tail *LazyList
}

var EmptyState State = State{nil, nil}
var EmptyLazyState LazyState = func() State { return EmptyState }
var NilLazyList LazyList = LazyList{&EmptyLazyState}

type LazyList struct {
	state *LazyState
}

func (l LazyList) Cons(i LazyInt) LazyList {
	xs := LazyList{l.state}
	state := func() State { return State{&i, &xs} }

	return LazyList{&state}
}

func (l LazyList) Map(f func(e int) int) LazyList {
	newState := func() State {
		state := (*l.state)()
		h := func() int { return f((*state.head)()) }
		t := state.tail.Map(f)

		return State{&h, &t}
	}
	return LazyList{&newState}
}

func (l LazyList) IsEmpty() bool { return l == NilLazyList }

func (l LazyList) Filter(p func(e int) bool) LazyList {
	restRef := l

	newState := func() State {

		var elem int
		var found = false
		var rest = restRef

		for !found && !rest.IsEmpty() {
			restState := (*rest.state)()
			elem = (*restState.head)()
			found = p(elem)
			rest = *restState.tail
			restRef = rest
		}

		if found {
			h := func() int { return elem }
			t := rest.Filter(p)
			return State{&h, &t}
		} else {
			return EmptyState
		}
	}

	return LazyList{&newState}
}

func (l LazyList) Take(n int) LazyList {
	if n <= 0 {
		return NilLazyList
	}

	newState := func() State {
		state := (*l.state)()
		t := state.tail.Take(n - 1)
		return State{state.head, &t}
	}
	return LazyList{&newState}
}

// Creates infinite list of numbers
func InfiniteIntList(from int) LazyList {
	el := func() int { return from }
	state := func() State {
		rest := InfiniteIntList(from + 1)
		return State{&el, &rest}
	}
	return LazyList{&state}
}

func main() {
	fmt.Println("Hello Lazy")

	xs := InfiniteIntList(1).
		Map(func(e int) int { return e + 1 }).
		Map(func(e int) int { return e - 1 }).
		Filter(func(e int) bool { return e%2 == 0 }).
		Filter(func(e int) bool { return e > 10 }).Take(100)

	for i := 0; !xs.IsEmpty(); i++ {
		state := (*xs.state)()
		fmt.Println((*state.head)())
		xs = *state.tail
	}
}
